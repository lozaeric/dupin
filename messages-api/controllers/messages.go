package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/messages-api/clients"
	"github.com/lozaeric/dupin/messages-api/domain"
	"github.com/lozaeric/dupin/toolkit/auth"
	"github.com/lozaeric/dupin/toolkit/validation"
)

func Message(c *gin.Context) {
	ID := c.Param("id")
	if !validation.IsValidID(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}
	token, err := auth.ParseToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "invalid token")
		return
	}
	if message, err := messageStore.Message(ID); err != nil {
		c.JSON(http.StatusNotFound, "id not found")
	} else if message.SenderID != token.UserID && message.ReceiverID != token.UserID {
		c.JSON(http.StatusForbidden, "you must be the sender or receiver")
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
	token, err := auth.ParseToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "invalid token")
		return
	}
	message.SenderID = token.UserID
	if _, err := clients.User(message.ReceiverID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "receiver not found",
		})
		return
	}
	if _, err := clients.User(message.SenderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "sender not found",
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
	if field != "" || value != "" {
		err := domain.CheckMessageValues(map[string]interface{}{field: value})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid values",
			})
			return
		}
	}
	token, err := auth.ParseToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, "invalid token")
		return
	}
	if messages, err := messageStore.Search(token.UserID, field, value); err != nil {
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
	tk, _ := c.Get("token")
	if token := tk.(*auth.Token); message.SenderID != token.UserID && message.ReceiverID != token.UserID {
		c.JSON(http.StatusForbidden, "you must be the sender or receiver")
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
