package main

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	service "my-test-service/service"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	service.TestUserService()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
