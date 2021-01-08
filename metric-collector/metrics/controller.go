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
	body, err := toJSON()
	if err != nil {
		c.String(http.StatusInternalServerError, "error while creating json")
	} else {
		c.Data(http.StatusOK, "application/json", body)
	}
}
