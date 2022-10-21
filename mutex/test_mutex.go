package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var wt sync.WaitGroup

var lock sync.Mutex

func add() {
	defer wt.Done()
	lock.Lock()
	i += 1
	fmt.Printf("i++ %v\n", i)
	time.Sleep(time.Millisecond * 2)
	lock.Unlock()
}

func sub() {
	defer wt.Done()
	lock.Lock()
	i -= 1
	fmt.Printf("i-- %v\n", i)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		//都是同步的，所以都会是100
		//加了go关键字之后，就不是了
		wt.Add(1)
		go add()
		wt.Add(1)
		go sub()
	}

	wt.Wait()
	//加了sleep之后，错误的几率更大了，所以需要用到mutex
	//共享资源枷锁，枷锁，互斥锁mutex
	fmt.Printf("i result %v\n", i)
}
