package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// 定义一个Histogram类型的指标
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "myapp",
			Subsystem: "http_server",
			Name:      "request_duration_seconds",
			Help:      "Histogram of response latency (seconds) of http requests.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "handler", "code"},
	)
	// 定义一个Counter类型的指标
	requestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "myapp",
			Subsystem: "http_server",
			Name:      "requests_total",
			Help:      "Counter of http requests.",
		},
		[]string{"method", "handler", "code"},
	)
)

func init() {
	// 注册指标
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestTotal)
}

func main() {
	// 暴露metrics接口
	http.Handle("/metrics", promhttp.Handler())

	// 定义业务路由
	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		// 记录请求时间
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			requestDuration.WithLabelValues(r.Method, "/example", "200").Observe(v)
		}))
		defer timer.ObserveDuration()

		// 记录请求总数
		requestTotal.WithLabelValues(r.Method, "/example", "200").Inc()

		w.Write([]byte("Hello, Prometheus!"))
	})
	// .
	fmt.Println("server is run ")
	http.ListenAndServe(":8080", nil)
}
