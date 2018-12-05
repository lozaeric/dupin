package controllers

import (
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := messageStore.Create(message); err != nil {
		c.JSON(http.StatusInternalServerError, "db error")
	} else {
		c.JSON(http.StatusOK, message)
	}
}

func SearchMessages(c *gin.Context) {
	field, value := c.Query("field"), c.Query("value")
	err := domain.CheckMessageValues(map[string]interface{}{field: value})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values",
		})
		return
	}
	if messages, err := messageStore.Search(field, value); err != nil {
		c.JSON(http.StatusInternalServerError, "db error")
	} else if len(messages) == 0 {
		c.JSON(http.StatusNotFound, "messages not found")
	} else {
		c.JSON(http.StatusOK, messages)
	}
}

func UpdateMessage(c *gin.Context) {
	ID := c.Param("id")
	if !validation.IsValidID(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}
	values := make(map[string]interface{})
	if err := c.BindJSON(&values); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body",
		})
		return
	}
	message, err := messageStore.Message(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, "message not found")
		return
	}
	err = message.Update(values)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid values")
		return
	}
	if err := messageStore.Update(message); err != nil {
		c.JSON(http.StatusInternalServerError, "db error")
	} else {
		c.JSON(http.StatusOK, message)
	}
}
