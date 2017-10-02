package promrus

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

// PrometheusHook exposes Prometheus counters for each of logrus' log levels.
type PrometheusHook struct {
	debugCounter prometheus.Counter
	infoCounter  prometheus.Counter
	warnCounter  prometheus.Counter
	errorCounter prometheus.Counter
}

// NewPrometheusHook creates a new instance of PrometheusHook which exposes Prometheus counters named after the provided prefix.
func NewPrometheusHook(prefix string) (*PrometheusHook, error) {
	debugCounter, err := createAndRegisterCounter(prefix, logrus.DebugLevel)
	if err != nil {
		return nil, err
	}
	infoCounter, err := createAndRegisterCounter(prefix, logrus.InfoLevel)
	if err != nil {
		return nil, err
	}
	warnCounter, err := createAndRegisterCounter(prefix, logrus.WarnLevel)
	if err != nil {
		return nil, err
	}
	errorCounter, err := createAndRegisterCounter(prefix, logrus.ErrorLevel)
	if err != nil {
		return nil, err
	}
	return &PrometheusHook{
		debugCounter: debugCounter,
		infoCounter:  infoCounter,
		warnCounter:  warnCounter,
		errorCounter: errorCounter,
	}, nil
}

func createAndRegisterCounter(prefix string, level logrus.Level) (prometheus.Counter, error) {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%v_%v", prefix, level),
		Help: fmt.Sprintf("Number of log statements at %v level.", level),
	})
	err := prometheus.Register(counter)
	if err == nil {
		return counter, nil
	}
	return nil, err
}

// Fire increments the appropriate Prometheus counter depending on the entry's log-level.
func (hook *PrometheusHook) Fire(entry *logrus.Entry) error {
	switch entry.Level {
	case logrus.ErrorLevel:
		hook.errorCounter.Inc()
	case logrus.WarnLevel:
		hook.warnCounter.Inc()
	case logrus.InfoLevel:
		hook.infoCounter.Inc()
	case logrus.DebugLevel:
		hook.debugCounter.Inc()
	}
	return nil
}

var supportedLevels = []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel}

// Levels returns all supported log-levels, i.e.: Debug, Info, Warn and Error, as there is no point incrementing a counter just before exiting/panicking.
func (hook *PrometheusHook) Levels() []logrus.Level {
	return supportedLevels
}
