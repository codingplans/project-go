package main

import (
	"log"
	"net/http"
	"time"

	"codeup.aliyun.com/qimao/leo/gin-timeout"
	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, `{"error": "timeout error"} `+c.Query("i"))
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(1000*time.Millisecond),

		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

func main() {
	r := gin.New()
	r.Use(timeoutMiddleware())
	r.GET("/slow", func(c *gin.Context) {

		time.Sleep(999*time.Millisecond + 500*time.Microsecond) // wait
		c.String(http.StatusOK, `handler`+c.Query("i"))
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
