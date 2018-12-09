package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lozaeric/dupin/auth-api/domain"
	"github.com/lozaeric/dupin/toolkit/validation"
)

func Token(c *gin.Context) {
	ID := c.Param("id")
	if !domain.IsValidID(ID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}
	if token, err := tokenStore.Token(ID); err != nil {
		c.JSON(http.StatusNotFound, "token not found")
	} else {
		c.JSON(http.StatusOK, token)
	}
}

func CreateToken(c *gin.Context) {
	token := new(domain.Token)
	if err := c.BindJSON(token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values",
		})
		return
	}
	if err := validation.Validate(token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	secureToken, err := tokenGenerator.Generate(token.ApplicationID, token.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "token generator error")
		return
	}
	if err := tokenStore.Create(secureToken); err != nil {
		c.JSON(http.StatusInternalServerError, "db error")
	} else {
		c.JSON(http.StatusOK, secureToken)
	}
}
