package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	later, _ := time.ParseDuration("5s")

	ctx, cancel := context.WithDeadline(context.Background(), now.Add(later))
	defer cancel()
	go Monitor(ctx)

	time.Sleep(10 * time.Second)

}

func Monitor(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(10 * time.Second):
		fmt.Println("stop monitor")
	}
}
