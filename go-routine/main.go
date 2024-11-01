package main

import (
	"context"
	"fmt"
	"math"
	"sync"
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
	var wg sync.WaitGroup
	for i := 0; i < math.MaxInt32; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}
