package metrics

import "encoding/json"

func Increment(metric NAME) {
	metrics[metric].Add(1)
}

func toJSON() ([]byte, error) {
	return json.Marshal(metrics)
}
