package main

import (
	"context"
	"fmt"
	"time"
)

/*
*
Context机制最核心的功能是在goroutine之间传递cancel信号，但是它的实现是不完全的。
Cancel可以细分为主动与被动两种，通过传递context参数，让调用goroutine可以主动cancel被调用goroutine。
但是如何得知被调用goroutine什么时候执行完毕，这部分Context机制是没有实现的。
而现实中的确又有一些这样的场景，比如一个组装数据的goroutine必须等待其他goroutine完成才可开始执行，这是context明显不够用了，必须借助sync.WaitGroup。
*/
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go HelloHandle(ctx, 2000*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("Hello Handle ", ctx.Err())
	}

}

func HelloHandle(ctx context.Context, duration time.Duration) {

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}

}
