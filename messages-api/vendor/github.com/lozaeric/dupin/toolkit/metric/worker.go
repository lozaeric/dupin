package metric

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-resty/resty"
)

type metricDTO struct {
	Name         string `json:"name"`
	DurationInMs int64  `json:"duration_in_ms"`
	Successful   bool   `json:"successful"`
}

var metricsCli = resty.New().
	SetTimeout(50 * time.Millisecond).
	SetRetryCount(1).
	AddRetryCondition(func(r *resty.Response) (bool, error) {
		return r == nil || r.Error() != nil || r.StatusCode() >= 500, nil
	}).
	SetHostURL("http://metric-collector:8080")

func doRecordMetric(dto metricDTO) error {
	r, err := metricsCli.R().SetBody(dto).Put("/api/metric/" + dto.Name)
	if err != nil || r.StatusCode() != http.StatusOK {
		return errors.New("metric-collector error")
	}
	return nil
}

func worker() {
	for m := range pending {
		doRecordMetric(m)
	}
}
