package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

/*
*
我们基于context.Background创建一个携带trace_id的ctx，然后通过context树一起传递，从中派生的任何context都会获取此值，我们最后打印日志的时候就可以从ctx中取值输出到日志中。
目前一些RPC框架都是支持了Context，所以trace_id的向下传递就更方便了。
*/
const (
	KEY = "trace_id"
)

func NewRequestID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func NewContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
	return ctx
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s | info | trace_id = %s | %s\n", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
}

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "Cashier go go go")
}

func main() {
	ProcessEnter(NewContextWithTraceID())
}
