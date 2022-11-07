package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"net/http"
	"time"
)

func main() {

	r := chi.NewRouter()
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("adasda"))
	})
	http.ListenAndServe(":2333", r)
}
