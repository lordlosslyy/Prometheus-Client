package main

import (
	"gin-prometheus-3/increase"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode) // otherwise, debug mode
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
		//customCounter.Inc()
	})

	r.GET("/metrics", prometheusHandler())

	increase.Register()
	for i := 0; i < 10000000; i++ {
		go increase.Increment()
	}
	increase.Increment()
	r.Run()
}
