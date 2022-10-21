package main

import (
	"fmt"
	"runtime"
)

func main() {
	go show("javaaaa") //goroutine
	for i := 0; i < 2; i++ {
		/**
		goexit()用法
		*/

		runtime.Gosched() //我有权利执行任务了，然后可以让给其他子携程先执行吧
		fmt.Println("golang")
	}
	fmt.Println("end")
}

func show2() {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit() //就直接退出了，让出cpu了
		}
		runtime.Gosched() //我有权利执行任务了，然后可以让给其他子携程先执行吧
		fmt.Println("golang")
	}
}

func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}
