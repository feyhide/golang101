package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // If context is canceled
			fmt.Println("Operation cancelled!")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond) // Simulating work
		}
	}
}

func runEx2() {
	ctx, cancel := context.WithCancel(context.Background())

	go doSomething(ctx)

	time.Sleep(2 * time.Second) // Let it run for a while
	cancel()                    // Cancel the operation
	time.Sleep(1 * time.Second) // Give goroutine time to exit
}
