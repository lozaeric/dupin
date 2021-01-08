package metrics

import (
	"expvar"

	"github.com/zserge/metric"
)

type NAME string

const (
	SENT_MESSAGES NAME = "sent_messages"
	CREATED_USERS NAME = "created_users"
)

var metrics = map[NAME]metric.Metric{
	SENT_MESSAGES: metric.NewCounter("10m1m"),
	CREATED_USERS: metric.NewCounter("10m1m"),
}

func init() {
	Increment(CREATED_USERS)
	Increment(SENT_MESSAGES)
	Increment(SENT_MESSAGES)
	for name, metric := range metrics {
		expvar.Publish(string(name), metric)
	}
}
