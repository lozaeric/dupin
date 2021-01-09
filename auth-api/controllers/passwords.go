package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/auth-api/passwords"
)

func CreatePassword(c *gin.Context) {
	data, _ := c.GetRawData()
	if err := passwords.Create(data); err != nil {
		c.JSON(err.StatusCode, gin.H{
			"message": err.Message,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "succesfully created",
	})
}
