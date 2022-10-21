package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() //goroutine finished, set as -1
	fmt.Println("hello goroutine waitgroup", i)

}

////两个协程之间相互同步，waitGroup
//func main() {
//	for i := 0; i < 10; i++ {
//		wg.Add(1) //launch a goroutine, count 1
//		go hello(i)
//	}
//	wg.Wait() //wait for all goroutine finished
//}

//对比用没有waitGroup,不用wait的话，主程序跑完之后goroutine就没了

func showMsg(i int) {
	defer wg.Done()
	fmt.Printf("i: %v\n", i)
}

//reference
//https://blog.csdn.net/weixin_45486746/article/details/122161953
func main() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
		wg.Add(1)
	}
	wg.Wait()
}
