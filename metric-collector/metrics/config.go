package metrics

import (
	"errors"
	"expvar"

	"github.com/zserge/metric"
)

type Kind int

const (
	COUNTER Kind = iota
	DURATION
)

const (
	SENT_MESSAGES = "sent_messages"
	SEEN_MESSAGES = "seen_messages"
	CREATED_USERS = "created_users"
	INCREMENT     = 1
	///////////////////////////////
	errSuffix      = "_err"
	durationSuffix = "_time"
	counterSuffix  = "_count"
)

var metrics = map[string]metric.Metric{
	SENT_MESSAGES:             metric.NewCounter("10m1m"),
	SEEN_MESSAGES:             metric.NewCounter("10m1m"),
	CREATED_USERS:             metric.NewCounter("10m1m"),
	SENT_MESSAGES + errSuffix: metric.NewCounter("10m1m"),
	SEEN_MESSAGES + errSuffix: metric.NewCounter("10m1m"),
	CREATED_USERS + errSuffix: metric.NewCounter("10m1m"),
}

type metricDTO struct {
	Name         string  `json:"name"`
	DurationInMs float64 `json:"duration_in_ms"`
	Successful   bool    `json:"successful"`
}

func metricName(name string, kind Kind, isErr bool) string {
	actualName := name
	if kind == COUNTER {
		actualName += counterSuffix
	} else if kind == DURATION {
		actualName += durationSuffix
	}
	if isErr {
		actualName += errSuffix
	}
	return actualName
}

func recordMetric(dto *metricDTO) error {
	counter := metricName(dto.Name, COUNTER, !dto.Successful)
	duration := metricName(dto.Name, DURATION, !dto.Successful)
	if metrics[counter] == nil || metrics[duration] == nil {
		return errors.New("invalid metric name")
	}
	metrics[counter].Add(INCREMENT)
	metrics[duration].Add(dto.DurationInMs)
	return nil
}

func init() {
	names := []string{SENT_MESSAGES, SEEN_MESSAGES, CREATED_USERS}
	for _, n := range names {
		for _, k := range []Kind{COUNTER, DURATION} {
			for _, e := range []bool{false, true} {
				metrics[metricName(n, k, e)] = metric.NewCounter("10m1m")
			}
		}
	}
	for name, metric := range metrics {
		expvar.Publish(name, metric)
	}
}