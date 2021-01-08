package metrics

import (
	"expvar"

	"github.com/zserge/metric"
)

const (
	SENT_MESSAGES = "sent_messages"
	CREATED_USERS = "created_users"
	INCREMENT     = 1
)

var metrics = map[string]metric.Metric{
	SENT_MESSAGES: metric.NewCounter("10m1m"),
	CREATED_USERS: metric.NewCounter("10m1m"),
}

func init() {
	metrics[CREATED_USERS].Add(INCREMENT)
	metrics[SENT_MESSAGES].Add(INCREMENT)
	metrics[SENT_MESSAGES].Add(INCREMENT)
	for name, metric := range metrics {
		expvar.Publish(name, metric)
	}
}
