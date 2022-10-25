package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

/**
对第三方调用要传入 context，用于控制远程调用。
不要将上下文存储在结构类型中，尽可能的作为函数第一位形参传入。
函数调用链必须传播上下文，实现完整链路上的控制。
context 的继承和派生，保证父、子级 context 的联动。
不传递 nil context，不确定的 context 应当使用 TODO。
context 仅传递必要的值，不要让可选参数揉在一起。

为什么要使用上下文(context)而不是计时器(timer)加通道(channel)的方式来控制协程
https://blog.csdn.net/pengpengzhou/article/details/107123560

*/
// START_OMIT
func processing(ctx context.Context) {
	limitOrderTransactionId, ok := ctx.Value("limit_order_transaction_id").(string)
	if ok {
		fmt.Printf("processing the transaction id: %s\n\n", limitOrderTransactionId)
	} else {
		fmt.Println("processing is over.")
	}
}
func main() {
	ctx := context.Background()
	processing(ctx)
	ctx = context.WithValue(ctx, "limit_order_transaction_id", uuid.New().String())
	processing(ctx)
}

// END_OMIT
