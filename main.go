package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace:   "app",
	Subsystem:   "system",
	Name:        "request",
	Help:        "request counter",
	ConstLabels: map[string]string{},
})

func init() {
	prometheus.DefaultRegisterer.Unregister(collectors.NewGoCollector())
	prometheus.MustRegister(requestCounter)
}

// go语言实现 http服务端
// http://127.0.0.1:8989/hello
func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		requestCounter.Inc()
		_, _ = writer.Write([]byte("hello world"))
	})
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe("127.0.0.1:8989", nil)
	if err != nil {
		panic(err)
	}
}
