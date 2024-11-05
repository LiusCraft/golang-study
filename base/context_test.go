package main

import (
	"context"
	"fmt"
	"testing"
)

func maiaan(ctx context.Context) context.Context {
	value := ""
	// 使用指针来持有值
	ctx = context.WithValue(ctx, "key", &value)

	// 打印父级上下文的值
	fmt.Println("Before callback:", ctx.Value("key").(*string))

	// 回调函数
	callback := func(c context.Context) {
		if v, ok := c.Value("key").(*string); ok {
			*v = "modified value" // 修改指针指向的内容
		}
	}

	// 在这里调用回调函数
	callback(ctx)

	fmt.Println("After callback:", ctx.Value("key").(*string))

	// 返回更新后的上下文
	return ctx
}

func TestSetPointValue(t *testing.T) {
	background := context.Background()
	// 获取更新后的上下文
	updatedCtx := maiaan(background)

	// 打印更新后的值
	fmt.Println("After maiaan:", (*updatedCtx.Value("key").(*string)))

}
