package main

import (
	"context"
	"fmt"
	"time"
)

func test(ctx context.Context, duration time.Duration) {
	fmt.Println("time:", duration)
	select {
	case <-ctx.Done():
		fmt.Println("handler", ctx.Err())
	case <-time.After(duration):
		fmt.Println("request 500", ctx.Value("wawdadw"))
	default:
	}
	fmt.Println("aaaaaaaaaaa")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	ctx.Value("wawdadw")
	defer cancel()

	go test(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
