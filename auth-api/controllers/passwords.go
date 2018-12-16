package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/auth-api/domain"
	"github.com/lozaeric/dupin/toolkit/validation"
)

func CreatePassword(c *gin.Context) {
	dto := make(map[string]string)
	if err := c.BindJSON(&dto); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values",
		})
		return
	}
	if !validation.IsValidID(dto["user_id"]) || dto["password"] == "" { // password need validations and user must exist
		fmt.Println(dto)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values",
		})
		return
	}

	info, ok := domain.NewSecureInfo(dto["user_id"], dto["password"])
	if !ok {
		c.JSON(http.StatusInternalServerError, "password generator error")
	} else if err := secureInfoStore.Create(info); err != nil {
		c.JSON(http.StatusInternalServerError, "db error")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "created sucesfully",
		})
	}
}

func ValidatePassword(c *gin.Context) {
	dto := make(map[string]string)
	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values",
		})
		return
	}
	if !validation.IsValidID(dto["user_id"]) || dto["password"] == "" { // password need validations and user must exist
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values",
		})
		return
	}

	if info, err := secureInfoStore.SecureInfo(dto["user_id"]); err != nil {
		c.JSON(http.StatusInternalServerError, "db error") // not found?
	} else if !info.IsCorrect(dto["password"]) {
		c.JSON(http.StatusBadRequest, "incorrect values")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "valid",
		})
	}
}
