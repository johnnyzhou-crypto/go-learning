package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带缓冲的通道

	//内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
		//use close(c) to close a channel
	}()

	time.Sleep(2 * time.Second) //延时2s
	for i := 0; i < 3; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}
	fmt.Println("main进程结束")

	//simpleChannel()
}

func simpleChannel() {

	// 使用 `make(chan val-type)` 创建一个新的通道。
	// 通道类型就是他们需要传递值的类型。
	messages := make(chan string)

	// 使用 `channel <-` 语法 _发送_ 一个新的值到通道中。
	// 这里我们在一个新的协程中发送 `"ping"` 到上面创建的 `messages` 通道中。
	go func() { messages <- "ping" }()

	// 使用 `<-channel` 语法从通道中 _接收_ 一个值。
	// 这里我们会收到在上面发送的 `"ping"` 消息并将其打印出来。
	msg := <-messages
	fmt.Println(msg)
}
