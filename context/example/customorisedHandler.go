package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

const requestKey int = 0

func WithRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		reqId := request.Header.Get("X-Request-Id")
		//use withValue context to customise value type, avoid conflicts
		ctx := context.WithValue(request.Context(), requestKey, reqId)
		//create a new request
		req := request.WithContext(ctx)
		next.ServeHTTP(writer, req)
	})
}

func GetRequestId(ctx context.Context) string {
	id := ctx.Value(requestKey).(string)
	return id
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func CustomisedHandler(writer http.ResponseWriter, request *http.Request) error {
	q := request.URL.Query().Get("err")
	if q != "" {
		return errors.New(q)
	}

	reqId := request.Header.Get("X-Request-Id")
	//use withValue context to customise value type, avoid conflicts
	ctx := context.WithValue(request.Context(), requestKey, reqId)
	//create a new request
	req := request.WithContext(ctx)
	reqId = GetRequestId(req.Context())
	fmt.Printf("request id is: %s\n\n", reqId)
	writer.WriteHeader(http.StatusOK)
	//todo
	writer.Write([]byte(fmt.Sprintf("welcome to Cashier service.\n")))
	return nil
}

/*
*
curl --location --request GET 'http://localhost:3333/cashier' \
--header 'X-Request-ID: asdasdasdasd111213css' \
--header 'Cookie: sid=772de8f4149a789987114a9742435237'

对于 Web 服务端开发，往往希望将一个请求处理的整个过程串起来，这就非常依赖于 Thread Local（对于 Go 可理解为单个协程所独有） 的变量，
而在 Go 语言中并没有这个概念，因此需要在函数调用的时候传递 context。
*/
func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Method("GET", "/cashier", Handler(CustomisedHandler))
	http.ListenAndServe(":3333", r)
	//handler := WithRequestId(http.HandlerFunc(CustomisedHandler))
	//http.ListenAndServe("localhost:8000", handler)
}
