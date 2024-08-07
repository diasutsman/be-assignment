package routes

import (
    "be-test-concrete-ai/account-manager/controllers"
    "be-test-concrete-ai/account-manager/middleware"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/register", controllers.RegisterUser)
    r.POST("/login", controllers.LoginUser)

    auth := r.Group("/")
    auth.Use(middleware.Auth())
    auth.GET("/accounts", controllers.GetAccounts)
	auth.POST("/accounts", controllers.CreateAccount)
    auth.GET("/accounts/:id/transactions", controllers.GetAccountTransactions)
}
