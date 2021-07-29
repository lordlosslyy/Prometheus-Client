package main

import (
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var customCounter = prometheus.NewCounter(prometheus.CounterOpts{
    Name: "custom_counter",
    Help: "number of custom count",
})

func init() {
    prometheus.MustRegister(customCounter)
}

func prometheusHandler() gin.HandlerFunc {
    h := promhttp.Handler()

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}

func main() {
    gin.SetMode(gin.ReleaseMode)  // otherwise, debug mode
    r := gin.New()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, "Hello world!")
        customCounter.Inc()
    })

    r.GET("/metrics", prometheusHandler())

    r.Run()
}

