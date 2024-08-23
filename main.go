package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of Http Requests",
	},
		[]string{"path", "method"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds",
		},
		[]string{"path", "method"},
	)
)


func init() {
	// Register metrics
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func main() {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start).Seconds()

		httpRequestsTotal.WithLabelValues(ctx.FullPath(), ctx.Request.Method).Inc()
		httpRequestDuration.WithLabelValues(ctx.FullPath(), ctx.Request.Method).Observe(duration)
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Prometheus with Gin!")
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	log.Fatal(r.Run(":9000"))

}
