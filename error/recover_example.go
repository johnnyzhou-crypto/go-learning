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

Golang错误和异常（panic）是可以互相转换的：

错误转异常：比如程序逻辑上尝试请求某个URL，最多尝试三次，尝试三次的过程中请求失败是错误，尝试完第三次还不成功的话，失败就被提升为异常了。

异常转错误：比如panic触发的异常被recover恢复后，将返回值中error类型的变量进行赋值，以便上层函数继续走错误处理流程。
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

/*
*
go源代码很多地方写panic, 但是工程实践业务代码不要主动写panic，
理论上panic只存在于server启动阶段，比如config文件解析失败，端口监听失败等等，所有业务逻辑禁止主动panic，所有异步的goroutine都要用recover去兜底处理。
*/
func panic_example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("calm down and eat a banana")
		}
	}()
	panic(errors.New("explosion"))
}
