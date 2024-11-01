package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWRMap(t *testing.T) {
	theMap := map[string]string{}

	//  W one
	for i := range 20 {
		t.Log(i)
		go func() {
			theMap["a"] = "a"
		}()
	}

	for i := range 20 {
		t.Log(i)
		go func() {
			t.Log(theMap["a"])
		}()
	}

	time.Sleep(2 * time.Second)

}

func Test1(t *testing.T) {
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

func TestGoroutineCount(t *testing.T) {
	for true {
		go func() {
			time.Sleep(1 * time.Second)
		}()
	}
}
