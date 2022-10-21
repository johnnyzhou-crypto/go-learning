package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
这个程序建立了2个管道一个传输int，一个传输string，同时启动了3个协程，前2个协程非常简单，
就是每隔1s向管道输出数据，第三个协程是不停的从管道取数据，和之前的例子不一样的地方是，pump1 和 pump2是2个不同的管道，
通过select可以实现在不同管道之间切换，哪个管道有数据就从哪个管道里面取数据，如果都没数据就等着，
还有一个定时器功能可以每隔一段时间向管道输出内容！而且我们可以很容易启动多个消费者。


*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(time.Duration(time.Second * 30))
	
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
		time.Sleep(time.Duration(time.Second))
	}
}

func pump2(ch chan string) {
	for i := 0; ; i++ {
		ch <- strconv.Itoa(i + 5)
		time.Sleep(time.Duration(time.Second))
	}
}

func suck(ch1 chan int, ch2 chan string) {
	chRate := time.Tick(time.Duration(time.Second * 5)) // 定时器
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %s\n", v)
		case <-chRate:
			fmt.Printf("Log log...\n")
		}
	}
}
