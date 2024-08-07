package controllers

import (
	"be-test-concrete-ai/account-manager/models/requests"
	"be-test-concrete-ai/account-manager/services"
	"be-test-concrete-ai/account-manager/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAccounts(c *gin.Context) {
	userID := c.MustGet("userID").(int)
	accounts, err := services.GetAccounts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func GetAccountTransactions(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id is not int")
	}

	transactions, err := services.GetAccountTransactions(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func CreateAccount(c *gin.Context) {
	userID := c.MustGet("userID").(int)
	var account requests.CreateAccountRequest
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	errValidation := validate.Struct(account)
	validationErrors := errValidation
	if validationErrors != nil {
		errorArray := utils.GetErrorArray(validationErrors.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{"error": errorArray})
		return
	}
	_, err := services.CreateAccount(&account, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully!"})
}
