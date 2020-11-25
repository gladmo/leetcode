package leet

import (
	"context"
	"time"
)

func Timeout(duration time.Duration, fn func(ctx context.Context, cancel context.CancelFunc)) (ok bool) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(duration))
	defer cancel()

	go fn(ctx, cancel)

	select {
	case <-ctx.Done():
		t, _ := ctx.Deadline()
		return time.Now().After(t)
	}
}
