# How to use 

```
func SetInterval(cb func(), second time.Duration) {
	for range time.Tick(time.Second * second) {
		cb()
	}
}

redisPubsub, err := redis.NewRedisPubsub(&redis.RedisPubsubOpts{})
err = redisPubsub.Subscribe("cmd")
err = redisPubsub.Subscribe("cmd2")

time.AfterFunc(10*time.Second, func() {
    fmt.Println("Unsubscribe cmd2")
    redisPubsub.Unsubscribe("cmd2")
})

go SetInterval(func() {
    fmt.Println("Send")
    err = redisPubsub.Publish("cmd", "hello cmd")
    err = redisPubsub.Publish("cmd2", "hello cmd2")
}, 3)

for msg := range adapter.Message() {
    fmt.Printf("Received message for : %s\n", msg)
}
```