package infra

import (
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	client       *redis.Client
	message      chan BaseMessage
	stopChannels map[string]chan BaseMessage
	mutex        sync.Mutex
}

type RedisClientOptions struct {
	Host string
	Port string
}

func NewRedisClient(opts RedisClientOptions) AdapterPort {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", opts.Host, opts.Port),
	})

	return &RedisAdapter{
		client:       client,
		message:      make(chan BaseMessage),
		stopChannels: make(map[string]chan BaseMessage),
		mutex:        sync.Mutex{},
	}
}

func (i *RedisAdapter) Publish(topic string, payload interface{}) bool {
	err := i.client.Publish(context.TODO(), topic, payload).Err()
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

	pubsub := i.client.Subscribe(context.TODO(), topic)
	defer pubsub.Close()

	for {
		select {
		case <-stopChan:
			return
		default:
			msg, err := pubsub.ReceiveMessage(context.TODO())
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
