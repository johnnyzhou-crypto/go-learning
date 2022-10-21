package main

import (
	"fmt"
)

/**
notes
默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得goroutine同步变的更加的简单，而不需要显式的lock。

当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

当 参数capacity= 0 时，channel 是无缓冲阻塞读写的；当capacity > 0 时，channel 有缓冲、是非阻塞的，直到写满 capacity个元素才阻塞写入。

channel非常像生活中的管道，一边可以存放东西，另一边可以取出东西。channel通过操作符 <- 来接收和发送数据，发送和接收数据语法


    make(chan Type)  //等价于make(chan Type, 0)
    make(chan Type, capacity)


和map类似，channel也一个对应make创建的底层数据结构的引用。

当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

*/
func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("child go end")

		fmt.Println("child go is running……")

		c <- 1234567890 //1234567890 was sent to channel
	}()

	num := <-c //read from channel

	fmt.Println("num = ", num)
	fmt.Println("main go end.")
}
