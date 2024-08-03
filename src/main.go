package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cstungthanh/sandbox/src/infra"
	"github.com/cstungthanh/sandbox/src/utils"
)

func main() {
	// go utils.PrintUsage()

	fmt.Println(" =============== Start App ==================")
	config, err := utils.LoadConfig(".")
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

	ctx, cancel := context.WithCancel(context.Background())
	go utils.SetInterval(ctx, func() {
		fmt.Println("Send")
		adapter.Publish("cmd", "hello cmd")
		adapter.Publish("cmd2", "hello cmd2")
	}, 3)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Canceling context")
		cancel()
	})

	for msg := range adapter.Message() {
		fmt.Printf("Received message for : %s\n", msg)

	}
}
