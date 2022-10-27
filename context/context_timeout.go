package main

import (
	"context"
	"fmt"
	"time"
)

/*
*
我们可以通过一个代码片段了解 context.Context 是如何对信号进行同步的。
在这段代码中，我们创建了一个过期时间为 1s 的上下文，并向上下文传入 handle 函数，该方法会使用 500ms 的时间处理传入的请求：
因为过期时间大于处理时间，所以我们有足够的时间处理该请求
handle 函数没有进入超时的 select 分支，但是 main 函数的 select 却会等待 context.Context 超时并打印出 main context deadline exceeded。
如果我们将处理请求时间增加至 1500ms，整个程序都会因为上下文的过期而被中止，：
*/
// START_OMIT
func handler(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handling", ctx.Err())
	case <-time.After(duration):
		fmt.Println("processing request with", duration)
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	/**
	讲重点
	我讲了这个东西要节约大家的开发时间，开箱即用
	*/
	go handler(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

// END_OMIT
