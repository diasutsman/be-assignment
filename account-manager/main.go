package main

import (
	"be-test-concrete-ai/account-manager/config"
	_ "be-test-concrete-ai/account-manager/docs"
	"be-test-concrete-ai/account-manager/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Account Manager API
// @version 1.0
// @description This is a sample server for account management.
// @host localhost:8081
// @BasePath /

func main() {
	config.Init()
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8081") // Run on port 8081
}
