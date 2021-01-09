package metric

import "time"

type Name string

const (
	SENT_MESSAGES Name = "sent_messages"
	SEEN_MESSAGES Name = "seen_messages"
	CREATED_USERS Name = "created_users"
)

var pending = make(chan metricDTO, 1000)

func RecordMetric(metric Name, f func() error) {
	start := time.Now()
	err := f()

	pending <- metricDTO{
		Name:         string(metric),
		DurationInMs: time.Now().Sub(start).Milliseconds(),
		Successful:   err == nil,
	}
}

func init() {
	go worker()
}
