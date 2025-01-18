package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	testMetric = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "test_metric_total",
			Help: "A test metric for demonstration purposes.",
		},
	)
	metricsMutex sync.Mutex
)

func init() {
	prometheus.MustRegister(testMetric)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		metricsMutex.Lock()
		testMetric.Inc()
		metricsMutex.Unlock()
		fmt.Fprintln(w, "Test metric incremented with labels!")
	})

	fmt.Println("Starting server on port 9090...")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
