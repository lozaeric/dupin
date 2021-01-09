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

func MetricValueHandler(c *gin.Context) {
	dto := new(metricDTO)
	dto.Name = c.Param("name")

	if err := c.BindJSON(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "dto is invalid.",
		})
		return
	}
	if dto.DurationInUs <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "dto has invalid values.",
		})
		return
	}

	err := recordMetric(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "metric name is invalid.",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "metric value was saved.",
		})
	}
}
