package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ginHealth = promauto.NewCounterVec(
    prometheus.CounterOpts{
	    Name: "gin_request_count",
	    Help: "Number of gin count",
    },
    []string{"url"},
)

func main () {
    cachedHealthPort := "12345"
	gin.SetMode(gin.ReleaseMode) 
	router := gin.New()  // New() ? // default 
	healthGrp := router.Group("/")
	healthGrp.GET("/health", ReadinessProbeHandler)
    healthGrp.GET("/metrics", gin.WrapH(promhttp.Handler()))
    router.Run(fmt.Sprintf(":%s", cachedHealthPort))
}

/*
func LivenessProbeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Cached is living\n")
}
*/
func ReadinessProbeHandler(c *gin.Context) {
    ginHealth.WithLabelValues("test").Inc()
    c.String(http.StatusOK, "Cached is ready\n")
}
