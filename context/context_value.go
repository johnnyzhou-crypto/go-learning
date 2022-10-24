package main

import (
	"context"
	"fmt"
)

type key string

/*
*
context可以用来在goroutine之间传递上下文信息，相同的context可以传递给运行在不同goroutine中的函数，上下文对于多个goroutine同时使用是安全的，
context包定义了上下文类型，可以使用background、TODO创建一个上下文，在函数调用链之间传播context，
也可以使用WithDeadline、WithTimeout、WithCancel 或 WithValue 创建的修改副本替换它，听起来有点绕，
其实总结起就是一句话：context的作用就是在不同的goroutine之间同步请求特定的数据、取消信号以及处理请求的截止日期。

目前我们常用的一些库都是支持context的，例如gin、database/sql等库都是支持context的，这样更方便我们做并发控制了，只要在服务器入口创建一个context上下文，不断透传下去即可。
*/
func main() {
	ctx := context.WithValue(context.Background(), key("cashier"), "Iam cashier")
	GetValue(ctx, key("cashier"))
}

func GetValue(ctx context.Context, k key) {
	if vv, ok := ctx.Value(k).(string); ok {
		fmt.Println(vv)
	}
}
