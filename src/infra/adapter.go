package infra

type BaseMessage struct {
	Topic   string
	Payload interface{}
}

type AdapterPort interface {
	Publish(topic string, payload interface{}) bool
	Subscribe(topic string)
	Unsubscribe(topic string)
	Message() <-chan BaseMessage
}
