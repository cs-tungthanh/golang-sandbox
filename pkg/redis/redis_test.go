package redis_test

import (
	"context"
	"sync"
	"testing"
	"time"

	. "github.com/cstungthanh/sandbox/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRedisClient struct {
	*redis.Client
	mock.Mock
}

func NewMockRedisClient() *MockRedisClient {
	return &MockRedisClient{
		Client: redis.NewClient(&redis.Options{}),
	}
}

func (m *MockRedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	args := m.Called(ctx)
	return args.Get(0).(*redis.StatusCmd)
}

func (m *MockRedisClient) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	args := m.Called(ctx, channel, message)
	return args.Get(0).(*redis.IntCmd)
}

func (m *MockRedisClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	args := m.Called(ctx, channels)
	return args.Get(0).(*redis.PubSub)
}

func (m *MockRedisClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestPublish(t *testing.T) {
	mockClient := NewMockRedisClient()
	mockClient.On("Publish", mock.Anything, "test-topic", "test-message").Return(redis.NewIntResult(1, nil))

	pubsub := NewWithClient(mockClient)

	err := pubsub.Publish("test-topic", "test-message")
	assert.NoError(t, err)

	mockClient.AssertExpectations(t)
}

func TestSubscribe(t *testing.T) {
	mockClient := new(MockRedisClient)
	mockPubSub := redis.NewPubSub()
	mockClient.On("Subscribe", mock.Anything, []string{"test-topic"}).Return(mockPubSub)

	pubsub := &RedisPubsub{
		client:             mockClient,
		message:            make(chan BaseMessage),
		subscribedChannels: make(map[string]chan struct{}),
		mutex:              sync.Mutex{},
	}

	err := pubsub.Subscribe("test-topic")
	assert.NoError(t, err)

	// Allow some time for the goroutine to start
	time.Sleep(100 * time.Millisecond)

	assert.Len(t, pubsub.subscribedChannels, 1)
	assert.Contains(t, pubsub.subscribedChannels, "test-topic")

	mockClient.AssertExpectations(t)
}

func TestUnsubscribe(t *testing.T) {
	pubsub := &RedisPubsub{
		subscribedChannels: make(map[string]chan struct{}),
		mutex:              sync.Mutex{},
	}

	stopChan := make(chan struct{})
	pubsub.subscribedChannels["test-topic"] = stopChan

	err := pubsub.Unsubscribe("test-topic")
	assert.NoError(t, err)

	assert.Len(t, pubsub.subscribedChannels, 0)
	assert.NotContains(t, pubsub.subscribedChannels, "test-topic")

	// Check if the channel is closed
	_, ok := <-stopChan
	assert.False(t, ok)
}

func TestMessage(t *testing.T) {
	pubsub := &RedisPubsub{
		message: make(chan BaseMessage),
	}

	go func() {
		pubsub.message <- BaseMessage{Topic: "test-topic", Payload: "test-message"}
	}()

	msg := <-pubsub.Message()
	assert.Equal(t, "test-topic", msg.Topic)
	assert.Equal(t, "test-message", msg.Payload)
}
