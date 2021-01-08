package app

import "github.com/lozaeric/dupin/metric-collector/metrics"

func setRoutes() {
	router.GET("/metrics", metrics.WebHandler)
	router.GET("/api/metrics", metrics.JSONHandler)
	router.GET("/api/metrics/:name", metrics.JSONHandler)
	router.POST("/api/metrics/:name", metrics.MetricIncrementHandler)
}
