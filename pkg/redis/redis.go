package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisPubsub struct {
	client             *redis.Client
	message            chan BaseMessage
	subscribedChannels map[string]chan BaseMessage
	mutex              sync.RWMutex
}

type RedisPubsubOpts struct {
	Host string
	Port int
}

func NewWithClient(client *redis.Client) PubsubPort {
	return &RedisPubsub{
		client:             client,
		message:            make(chan BaseMessage),
		subscribedChannels: make(map[string]chan BaseMessage),
		mutex:              sync.RWMutex{},
	}
}

func New(opts *RedisPubsubOpts) (PubsubPort, error) {
	if opts.Host == "" {
		opts.Host = "localhost"
	}
	if opts.Port == 0 {
		opts.Port = 6379
	}

	client := redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%d", opts.Host, opts.Port),
		MaxRetries:      5,
		MinRetryBackoff: 1 * time.Second,
		MaxRetryBackoff: 2 * time.Second,
		DialTimeout:     10 * time.Second,
	})

	fmt.Printf("[RedisPubsub] Connecting to Redis: %s:%d... \n", opts.Host, opts.Port)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}
	fmt.Println("[RedisPubsub] Connected to Redis:")

	return &RedisPubsub{
		client:             client,
		message:            make(chan BaseMessage),
		subscribedChannels: make(map[string]chan BaseMessage),
		mutex:              sync.RWMutex{},
	}, nil
}

func (i *RedisPubsub) Publish(topic string, payload interface{}) error {
	err := i.client.Publish(context.TODO(), topic, payload).Err()
	if err != nil {
		return fmt.Errorf("publish message error: %v", err)
	}
	return nil
}

func (i *RedisPubsub) Subscribe(topic string) error {
	fmt.Println("[RedisPubsub] Subscribed to", topic)
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if _, ok := i.subscribedChannels[topic]; ok {
		return nil
	}

	stopChan := make(chan BaseMessage)
	i.subscribedChannels[topic] = stopChan

	// Start listener message from topic
	go func() {
		pubsub := i.client.Subscribe(context.TODO(), topic)
		defer pubsub.Close()

		for {
			select {
			case <-stopChan:
				return
			case msg := <-pubsub.Channel():
				i.message <- BaseMessage{
					Topic:   topic,
					Payload: msg.Payload,
				}
			}
		}
	}()

	return nil
}

func (i *RedisPubsub) Unsubscribe(topic string) error {
	fmt.Println("[RedisPubsub] Unsubscribed from", topic)
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if c, ok := i.subscribedChannels[topic]; ok {
		close(c)
		delete(i.subscribedChannels, topic)
	}
	return nil
}

func (i *RedisPubsub) Message() <-chan BaseMessage {
	return i.message
}

// Add a Close method for graceful shutdown
func (i *RedisPubsub) Close() error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	for topic, stopChan := range i.subscribedChannels {
		close(stopChan)
		delete(i.subscribedChannels, topic)
	}

	return i.client.Close()
}
