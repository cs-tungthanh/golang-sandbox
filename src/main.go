package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cstungthanh/sandbox/src/infra"
	"github.com/cstungthanh/sandbox/src/shared"
)

func SetInterval(cb func(), second time.Duration) {
	for range time.Tick(time.Second * second) {
		cb()
	}
}

func main() {
	fmt.Println("Start App!!!!!!")
	config, err := shared.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	adapter := infra.NewRedisClient(infra.RedisClientOptions{
		Host: config.RedisHost,
		Port: config.RedisPort,
	})
	go adapter.Subscribe("cmd")
	go adapter.Subscribe("cmd2")

	time.AfterFunc(10*time.Second, func() {
		fmt.Println("Unsubscribe cmd2")
		adapter.Unsubscribe("cmd2")
	})

	go SetInterval(func() {
		fmt.Println("Send")
		adapter.Publish("cmd", "hello cmd")
		adapter.Publish("cmd2", "hello cmd2")
	}, 3)

	for msg := range adapter.Message() {
		fmt.Printf("Received message for : %s\n", msg)

	}
}
