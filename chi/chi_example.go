package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

/**
go-chi特征
轻巧 -chi路由器的时钟约为1000 LOC
快速 -是，请参阅基准
100％与net / http兼容 -在生态系统中使用也与以下版本兼容的任何http或中间件pkgnet/http
专为模块化/组合式API设计 -中间件，嵌入式中间件，路由组和子路由器安装
上下文控制 -基于新context程序包，提供价值链，取消和超时
稳健 -在Pressly，CloudFlare，Heroku，99Designs和许多其他公司中进行生产（请参阅讨论）
文档生成 - docgen自动将路由文档从您的源生成为JSON或Markdown
没有外部依赖 -普通ol st Golib + net / http
*/
// START_OMIT
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}

// END_OMIT
