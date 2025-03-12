package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	RedirectCount prometheus.Counter
}

func New() *Metrics {
	m := Metrics{
		RedirectCount: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "redirect_count",
				Help: "Total redirects count",
			},
		),
	}

	prometheus.MustRegister(m.RedirectCount)

	return &m
}

func (m *Metrics) Handler() http.Handler {
	return promhttp.Handler()
}

func (m *Metrics) CollectRedirect() {
	m.RedirectCount.Inc()
}
