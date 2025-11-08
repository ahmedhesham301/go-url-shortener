package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var UrlsCount = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "url_shortener_shorten_urls_total",
		Help: "Total number of shorten urls",
	},
)

var HttpRequestLatency = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "url_shortener_http_request_latency_seconds",
		Help: "Request latency in seconds.",
	},
	[]string{"path"},
)
