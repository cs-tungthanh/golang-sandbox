package redis

type BaseMessage struct {
	Topic   string
	Payload interface{}
}

type PubsubPort interface {
	Close() error
	Publish(topic string, payload interface{}) error
	Subscribe(topic string) error
	Unsubscribe(topic string) error
	Message() <-chan BaseMessage
}
