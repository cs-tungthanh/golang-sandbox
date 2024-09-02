package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	instruction := ctx.Value("instruction").(string)
	fmt.Println(instruction)

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("worker: finished cleaning the house")
	case <-ctx.Done():
		fmt.Println("worker: context cancelled")
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "instruction", "clean the house")
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	go worker(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("main function ended")
}
