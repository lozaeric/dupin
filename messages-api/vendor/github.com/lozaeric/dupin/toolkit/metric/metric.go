package metric

import (
	"os"
	"time"
)

type Name string

const (
	SENT_MESSAGES Name = "sent_messages"
	SEEN_MESSAGES Name = "seen_messages"
	CREATED_USERS Name = "created_users"
)

var pending = make(chan metricDTO, 1000)

func RecordMetric(metric Name, start time.Time, statusCode int) {
	pending <- metricDTO{
		Name:         string(metric),
		DurationInMs: time.Now().Sub(start).Milliseconds(),
		Successful:   statusCode >= 200 && statusCode < 300,
	}
}

func init() {
	if os.Getenv("ENV") == "production" {
		go worker()
	}
}
