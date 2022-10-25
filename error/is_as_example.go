package main

import (
	"errors"
	"fmt"
)

type myErr struct {
	name string
}

/*
*
https://go.dev/blog/go1.13-errors
在包装链上寻找和target类型匹配的error，如果找到就将找到的error赋值给target（多个匹配项取第一个）并返回true，否则返回false。
必须实现error接口或者是接口类型，并且不能为空指针，否则会panic。
*/
func (e *myErr) Error() string {
	return e.name
}

func main() {
	errA := errors.New("this is error A; ")
	errB := fmt.Errorf("this is error B; %w", errA)
	errC := fmt.Errorf("this is error C; %w", errB)
	fmt.Println(errC)
	//判断包装链上是否存在给定的error
	hasErrA := errors.Is(errC, errA)
	fmt.Println(hasErrA)
	/**
	The errors.Is function compares an error to a value.
	// Similar to:
	//   if err == ErrNotFound { … }
	if errors.Is(err, ErrNotFound) {
	    // something wasn't found
	}
	check the error chain if contains target error, like a value comparison
	*/
	hasErrD := errors.Is(errC, errors.New("this is error D; "))
	/**
	The As function tests whether an error is a specific type.

	// Similar to:
	//   if e, ok := err.(*QueryError); ok { … }
	var e *QueryError
	// Note: *QueryError is the type of the error.
	if errors.As(err, &e) {
	    // err is a *QueryError, and e is set to the error's value
	}
	*/
	fmt.Println(hasErrD)
	testAs()
}

/*
*
更多error.As() 的封装， error code
https://github.com/cloudwego/kitex-examples/blob/main/bizdemo/easy_note/pkg/errno/errno.go
*/
func testAs() {
	errA := &myErr{name: "this is error A; "}
	errB := fmt.Errorf("this is error B; %w", errA)
	errC := fmt.Errorf("this is error C; %w", errB)
	fmt.Println(errC)
	var errD *myErr
	asResult := errors.As(errC, &errD)
	fmt.Println(asResult)
	fmt.Println(errD)
}
