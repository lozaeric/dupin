package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/domain"
)

func User(c *gin.Context) {
	ID := c.Param("id")
	// todo: validate id
	if user, err := userStore.User(ID); err != nil {
		c.JSON(http.StatusNotFound, "id not found")
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
	user := new(domain.User)
	if err := c.BindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	if err := domain.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := userStore.Create(user); err != nil {
		c.JSON(http.StatusNotFound, "id not found")
	} else {
		c.JSON(http.StatusOK, user)
	}
}
