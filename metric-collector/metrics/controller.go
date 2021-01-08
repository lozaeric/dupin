package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zserge/metric"
)

func WebHandler(c *gin.Context) {
	metric.Handler(metric.Exposed).ServeHTTP(c.Writer, c.Request)
}

func JSONHandler(c *gin.Context) {
	name := c.Param("name")
	var value interface{}

	if name == "" {
		value = metrics
	} else {
		value = metrics[name]
	}

	if value == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "metric name is invalid.",
		})
	} else {
		c.JSON(http.StatusOK, value)
	}
}

func MetricIncrementHandler(c *gin.Context) {
	name := c.Param("name")
	metric := metrics[name]

	if metric == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "metric name is invalid.",
		})
	} else {
		metric.Add(INCREMENT)
		c.JSON(http.StatusOK, metric)
	}
}
