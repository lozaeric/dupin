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

func RecordMetric(metric Name, start time.Time, statusCode func() int) {
	duration := int64(time.Now().Sub(start) / time.Millisecond)
	if duration == 0 {
		duration = 1
	}
	pending <- metricDTO{
		Name:         string(metric),
		DurationInMs: duration,
		Successful:   statusCode() >= 200 && statusCode() < 300,
	}
}

func init() {
	if os.Getenv("ENV") == "production" {
		go worker()
	}
}
