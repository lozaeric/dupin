package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/domain"
	"github.com/lozaeric/dupin/domain/validation"
)

func User(c *gin.Context) {
	ID := c.Param("id")
	if !validation.IsValidID(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}
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
	if err := validation.Validate(user); err != nil {
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

func UpdateUser(c *gin.Context) {
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
	user, err := userStore.User(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, "user not found")
		return
	}
	err = user.Update(values)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid values")
		return
	}
	if err := userStore.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, "db")
	} else {
		c.JSON(http.StatusOK, user)
	}
}
