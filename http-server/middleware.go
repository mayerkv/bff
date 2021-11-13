package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func metricsMiddleware(service, method string, histogramVec *prometheus.HistogramVec, counterVec *prometheus.CounterVec) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()

		ctx.Next()

		labels := map[string]string{
			"service": service,
			"method":  method,
			"code":    fmt.Sprintf("%d", ctx.Writer.Status()),
		}

		counterVec.With(labels).Add(1)
		histogramVec.With(labels).Observe(time.Since(now).Seconds())
	}
}
