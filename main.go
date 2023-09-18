package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//  1.定义并注册指标（类型，名字，帮助信息），promauto.NewCounter方法会注册自定义指标
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "The total number of processed events",
	})
	counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "call_requests_total",
			Help: "Total number of API requests",
		},
		[]string{"method", "endpoint"},
	)
)

func main() {
	// 注册指标
	prometheus.MustRegister(counter)
	// 创建一个HTTP处理函数，每次请求时增加计数器
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// 增加计数器的值，并传递相应的标签值
		counter.WithLabelValues("GET", "/users").Inc()
		// 返回响应
		w.Write([]byte("users hello world!\n"))
	})

	http.HandleFunc("/admins", func(w http.ResponseWriter, r *http.Request) {
		// 增加计数器的值，并传递相应的标签值
		counter.WithLabelValues("GET", "/admins").Inc()
		// 返回响应
		w.Write([]byte("admins hello world!\n"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 增加计数器的值，并传递相应的标签值
		opsProcessed.Inc()
		// 返回响应
		w.Write([]byte("hello world!\n"))
	})
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
