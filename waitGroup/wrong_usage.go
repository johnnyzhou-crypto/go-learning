package main

import (
	"fmt"
	"sync"
)

func sayHello(num int, wg *sync.WaitGroup) int {
	wg.Add(1)
	fmt.Println("every project comes from Hello world.")
	return num
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		go sayHello(i, &wg)
	}
	wg.Wait()
	fmt.Println("main end.")
}
