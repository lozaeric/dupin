package app

import "github.com/lozaeric/dupin/metric-collector/metrics"

func setRoutes() {
	router.GET("/metrics", metrics.WebHandler)
	router.GET("/metrics/json", metrics.JSONHandler)
}
