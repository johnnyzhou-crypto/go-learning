package main

import (
	"fmt"
	"sync"
)

func sayHelloInRightWay(num int, wg *sync.WaitGroup) int {
	defer wg.Done()
	fmt.Println("every project comes from Hello world.")
	return num
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go sayHelloInRightWay(i, &wg)
	}
	wg.Wait()
	fmt.Println("main end..")
}
