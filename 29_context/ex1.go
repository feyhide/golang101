package main

import (
	"context"
	"fmt"
	"time"
)

func slowFunc(ctx context.Context) {
	select {
	case <-time.After(time.Second * 3): // Simulate long work
		fmt.Println("function done")
	case <-ctx.Done(): // Cancel if timeout occurs
		fmt.Println("Timeout")
	}
}

func runEx1() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	// after two seconds it will timeout
	defer cancel()

	// will get timeout as timeout happen in 2 sec and func takes 3 sec
	go slowFunc(ctx)

	time.Sleep(time.Second * 3)
}
