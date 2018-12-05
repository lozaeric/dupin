package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/domain/validation"
)

func Message(c *gin.Context) {
	ID := c.Param("id")
	if !validation.IsValidID(ID) {
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
	if err := validation.Validate(message); err != nil {
		fmt.Println(err.Error())
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

func SearchMessages(c *gin.Context) {
	field, ID := c.Query("field"), c.Query("id")
	// validate field and value
	if field == "" || ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "empty field",
		})
		return
	}
	if messages, err := messageStore.Search(field, ID); err != nil || len(messages) == 0 {
		c.JSON(http.StatusNotFound, "id not found")
	} else {
		c.JSON(http.StatusOK, messages)
	}
}
