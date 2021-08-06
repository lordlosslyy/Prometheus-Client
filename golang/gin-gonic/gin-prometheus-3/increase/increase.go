package increase

import "github.com/prometheus/client_golang/prometheus"

var customCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "custom_counter",
	Help: "number of custom count",
})

func Register() {
	prometheus.MustRegister(customCounter)
}

func Increment() {
	customCounter.Inc()
}
