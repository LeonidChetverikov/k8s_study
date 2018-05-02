package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"path/filepath"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	metricapi "github.com/kubernetes/dashboard/src/app/backend/integration/metric/api"
	

)


var (
	Pod = "nginx"
	Namespace = "Web"
	cpuRequest = prometheus.NewGauge(prometheus.GaugeOpts{
		Subsystem: Pod,
		Namespace: Namespace,
		Name:      "pod_resource_requests_cpu",
		Help:      "Cpu usage.",
	})
	memRequest = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Subsystem: Pod,
			Namespace: Namespace,
			Name: "pod_resource_requests_mem",
			Help: "Memory usage.",
		})
	)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuRequest)
	prometheus.MustRegister(memRequest)
}

func main() {
config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	cpu, err := strconv.ParseFloat(metricapi.CpuUsage, 64)
	mem, err := strconv.ParseFloat(metricapi.MemoryUsage, 64)
	cpuRequest.Set(cpu)
	memRequest.Set(mem)
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

f
