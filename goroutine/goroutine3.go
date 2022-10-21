package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count = 10

//互斥锁
func Counter(l *sync.Mutex) {
	l.Lock()
	count++
	add(count, count)
	fmt.Println("count=", count)
	l.Unlock()
}

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Counter(lock)
	}
	for {
		lock.Lock()
		c := count
		lock.Unlock()
		runtime.Gosched() //让出时间片
		if c >= 10 {
			break
		}
	}

}
