package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) int {
	return a / b
}

/*
*
如果你调用的方法的开发者不够仔细，并没有将所有的异常都考虑到并作为错误返回，那你的程序可能就会被其影响而崩溃。对这种情况，GO 语言提供了一个叫recover()的函数，用于处理这种问题，保障程序不崩溃。
一般常用于服务启动的入口函数，因为网络等外部因素，极有可能会导致程序异常，这些异常就需要这个函数来捕获。
recover函数与panic函数相对应，recover可以对panic的goroutine进行处理，需要注意的是recover只有在defer才能生效，并且只能处理自己goroutine中的panic。
*/
func main() {
	//panic_example()
	recover_example()
}

func recover_example() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("error occur: %s\n", err)
		}
	}()
	res := divide(1, 0)
	fmt.Println(res)
}

func panic_example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("calm down and eat a banana")
		}
	}()
	panic(errors.New("explosion"))
}
