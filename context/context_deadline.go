package main

import (
	"context"
	"fmt"
	"time"
)

/**
WithDeadline 返回父上下文的副本，并将截止日期调整为不晚于d。如果父母的截止日期早于d，WithDeadline（parent，d）在语义上等同于父母。
当截止日期到期，返回的取消功能被调用时，或者父上下文的完成通道关闭时，返回的上下文的完成通道将关闭，以先发生者为准。
 取消这个上下文会释放与它相关的资源，所以只要完成在这个Context 中运行的操作，代码就应该调用cancel。
*/
// 如果父节点 context 的 deadline 早于指定时间。直接构建一个可取消的 context。
// 原因是一旦父节点超时，自动调用 cancel 函数，子节点也会随之取消。
// 所以不用单独处理子节点的计时器时间到了之后，自动调用 cancel 函数
// Note: anyway, 即使ctx过期，最好还是调用
// 取消功能无论如何。 如果不这样做可能会保留
// 上下文及其父级的活动时间超过必要时间
// START_OMIT
func main() {
	d := time.Now().Add(5000000 * time.Millisecond)
	//father
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept") //monitoring
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// END_OMIT
