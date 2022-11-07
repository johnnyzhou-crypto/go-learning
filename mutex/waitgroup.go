package main

import (
	"fmt"
	"sync"
)

type container struct {
	m       sync.Mutex
	counter map[string]int
}

func (c *container) increase(name string) {
	c.m.Lock()
	defer c.m.Unlock()
	c.counter[name]++
}

func main() {
	c := container{counter: map[string]int{"a": 0, "b": 0}}
	var wait sync.WaitGroup
	doIncrease := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.increase(name)
		}
		wait.Done()
	}
	wait.Add(3)
	go doIncrease("a", 1000)
	go doIncrease("b", 1000)
	go doIncrease("c", 1000)

	wait.Wait()
	fmt.Println(c.counter)
}
