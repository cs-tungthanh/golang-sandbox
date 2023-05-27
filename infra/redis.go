package infra

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

type RedisAdapter struct {
	client       *redis.Client
	message      chan BaseMessage
	stopChannels map[string]chan BaseMessage
	mutex        sync.Mutex
}

func NewClient() AdapterPort {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &RedisAdapter{
		client:       client,
		message:      make(chan BaseMessage),
		stopChannels: make(map[string]chan BaseMessage),
		mutex:        sync.Mutex{},
	}
}

func (i *RedisAdapter) Publish(topic string, payload interface{}) bool {
	err := i.client.Publish(topic, payload).Err()
	return err != nil
}

func (i *RedisAdapter) Subscribe(topic string) {
	if _, ok := i.stopChannels[topic]; ok {
		return
	}

	i.mutex.Lock()
	stopChan := make(chan BaseMessage)
	i.stopChannels[topic] = stopChan
	i.mutex.Unlock()

	pubsub := i.client.Subscribe(topic)
	defer pubsub.Close()

	for {
		select {
		case <-stopChan:
			return
		default:
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				fmt.Println("Error receiving message:", err)
				return
			}

			i.message <- BaseMessage{
				Topic:   topic,
				Payload: msg.Payload,
			}

		}
	}
}

func (i *RedisAdapter) Unsubscribe(topic string) {
	i.mutex.Lock()
	if c, ok := i.stopChannels[topic]; ok {
		close(c)
		delete(i.stopChannels, topic)
	}
	i.mutex.Unlock()
}

func (i *RedisAdapter) Message() <-chan BaseMessage {
	return i.message
}
