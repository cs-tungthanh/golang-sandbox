package utils

import (
	"context"
	"time"
)

func SetInterval(ctx context.Context, cb func(), second time.Duration) {
	ticker := time.NewTicker(time.Second * second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			cb()
		}
	}
}
