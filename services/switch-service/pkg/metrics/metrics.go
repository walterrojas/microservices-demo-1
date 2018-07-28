package metrics

import (
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics wraps a prometheus Registry
type Metrics struct {
	service  string
	Registry *prometheus.Registry
}

// NewVoidMetrics creates a new metrics for testing purposes
func NewVoidMetrics() *Metrics {
	return &Metrics{
		Registry: prometheus.NewRegistry(),
	}
}

// NewMetrics creates a new metrics
func NewMetrics(service string) *Metrics {
	service = strings.Replace(service, "-", "_", -1)

	registry := prometheus.NewRegistry()
	registry.MustRegister(prometheus.NewGoCollector())
	registry.MustRegister(prometheus.NewProcessCollector(os.Getpid(), service))

	return &Metrics{
		service:  service,
		Registry: registry,
	}
}

// NewCounter creates and registers a new counter metric
func (m *Metrics) NewCounter(name, help string, labels []string) *prometheus.CounterVec {
	opts := prometheus.CounterOpts{
		Name: name,
		Help: help,
	}

	counter := prometheus.NewCounterVec(opts, labels)
	m.Registry.MustRegister(counter)

	return counter
}

// NewGauge creates and registers a new gauge metric
func (m *Metrics) NewGauge(name, help string, labels []string) *prometheus.GaugeVec {
	opts := prometheus.GaugeOpts{
		Name: name,
		Help: help,
	}

	gauge := prometheus.NewGaugeVec(opts, labels)
	m.Registry.MustRegister(gauge)

	return gauge
}

// NewHistogram creates and registers a new histogram metric
func (m *Metrics) NewHistogram(name, help string, buckets []float64, labels []string) *prometheus.HistogramVec {
	opts := prometheus.HistogramOpts{
		Name:    name,
		Help:    help,
		Buckets: buckets,
	}

	histogram := prometheus.NewHistogramVec(opts, labels)
	m.Registry.MustRegister(histogram)

	return histogram
}

// NewSummary creates and registers a new summary metric
func (m *Metrics) NewSummary(name, help string, quantiles map[float64]float64, labels []string) *prometheus.SummaryVec {
	opts := prometheus.SummaryOpts{
		Name:       name,
		Help:       help,
		Objectives: quantiles,
	}

	summary := prometheus.NewSummaryVec(opts, labels)
	m.Registry.MustRegister(summary)

	return summary
}

// Handler returns http handler for metrics endpoint
func (m *Metrics) Handler() http.Handler {
	return promhttp.HandlerFor(m.Registry, promhttp.HandlerOpts{})
}