package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.Default()

	r := Tss()

	fmt.Println(123)
	r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println(123)
}

func Tss() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
