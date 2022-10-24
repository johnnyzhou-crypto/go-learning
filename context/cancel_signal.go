package main

import (
	"context"
	"fmt"
	"time"
)

func useContext(ctx context.Context, id int) {
	for range time.Tick(time.Second) {
		select {
		/**
		只读的chan
		c.done 是“懒汉式”创建，只有调用了 Done() 方法的时候才会被创建。再次说明，函数返回的是一个只读的 channel，而且没有地方向这个 channel 里面写数据。
		所以，直接调用读这个 channel，协程会被 block 住。一般通过搭配 select 来使用。一旦关闭，就会立即读出零值。

		*/
		case <-ctx.Done():
			fmt.Println("comes to stop.", id)
			return
		default:
			fmt.Println("default return.")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go useContext(ctx, 10068)
	time.Sleep(5 * time.Second)
	/**
	总体来看，cancel() 方法的功能就是关闭 channel：c.done；递归地取消它的所有子节点；从父节点从删除自己。达到的效果是通过关闭 channel，将取消信号传递给了它的所有子节点。
	goroutine 接收到取消信号的方式就是 select 语句中的读 c.done 被选中。
	*/
	cancel()
	time.Sleep(1 * time.Second)
}
