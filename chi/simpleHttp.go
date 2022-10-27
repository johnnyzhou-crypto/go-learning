package main

import (
	"fmt"
	"net/http"
)

/**
声明定义若干个 handler(w http.ResponseWriter, r *http.Request), 每个 handler 也就是服务端提供的每个服务的逻辑载体。
调用 http.HandleFunc 来注册 handler 到对应的路由下。
调用 http.ListenAndServe 来启动服务，监听某个指定的端口。
*/

// step1. 通过listener, 接受了一个请求
// step2. 确认请求未超时之后，创建一个conn 对象
// step3. 单独创建一个gorouting, 负责处理这个请求
/**

根据conn结构体的serve方法，负责解析 request 的函数是func readRequest(b *bufio.Reader, deleteHostHeader bool) (req *Request, err error),
这个函数将请求头和请求体中的字段放到 Request 结构体中并返回。

*/
// START_OMIT
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8000", nil)
}

// END_OMIT
