package middleware

import (
	"time"
	"urlshortener/monitor"

	"github.com/gin-gonic/gin"
)

func RequestLatency(c *gin.Context) {
	start := time.Now()
	c.Next()
	elapsed := time.Since(start)

	monitor.HttpRequestLatency.WithLabelValues(c.FullPath()).Observe(elapsed.Seconds())
}
