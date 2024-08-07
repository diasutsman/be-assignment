package controllers

import (
	"be-test-concrete-ai/account-manager/models"
	"be-test-concrete-ai/account-manager/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIsExists := services.UserIsExistsByUsername(user.Username)

	if userIsExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	_, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := services.LoginUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
