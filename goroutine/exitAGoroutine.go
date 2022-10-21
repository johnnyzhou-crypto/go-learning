package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B") // 不会执行
		}()

		fmt.Println("A") // 不会执行
	}()

	for i := 0; ; {
		i++
	}
}