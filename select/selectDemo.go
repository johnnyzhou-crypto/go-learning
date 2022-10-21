package main

import (
	"fmt"
)

/**
这样使用select就可以避免永久等待的问题，因为程序会在timeout中获取到一个数据后继续执行，而无论对ch的读取是否还处于等待状态

*/
func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
