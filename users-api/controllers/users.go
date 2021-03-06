package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/toolkit/auth"
	"github.com/lozaeric/dupin/toolkit/metric"
	"github.com/lozaeric/dupin/users-api/users"
)

func User(c *gin.Context) {
	ID := c.Param("id")
	user, err := users.Get(ID)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{
			"messsage": err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	defer metric.RecordMetric(metric.CREATED_USERS, time.Now(), c.Writer.Status)

	data, _ := c.GetRawData()
	user, err := users.Create(data)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{
			"messsage": err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	ID := c.Param("id")
	token, er := auth.ParseToken(c)
	if er != nil {
		c.JSON(http.StatusForbidden, "invalid token")
		return
	} else if token.UserID != ID {
		c.JSON(http.StatusForbidden, "you must be owner of the user")
		return
	}
	data, _ := c.GetRawData()
	user, err := users.Update(ID, data)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{
			"messsage": err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
