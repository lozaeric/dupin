package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/utils"
)

func Message(c *gin.Context) {
	ID := c.Param("id")
	if !utils.IsValidID(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}
	if message, err := messageStore.Message(ID); err != nil {
		c.JSON(http.StatusNotFound, "id not found")
	} else {
		c.JSON(http.StatusOK, message)
	}
}

func CreateMessage(c *gin.Context) {
	message := new(domain.Message)
	if err := c.BindJSON(message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid message",
		})
		return
	}
	if err := utils.Validate(message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := messageStore.Create(message); err != nil {
		c.JSON(http.StatusNotFound, "id not found")
	} else {
		c.JSON(http.StatusOK, message)
	}
}
