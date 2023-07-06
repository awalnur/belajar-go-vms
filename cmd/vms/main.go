package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"vms/api/routes"
	_ "vms/cmd/vms/docs"
	"vms/internal/controller"
	"vms/internal/domain"
)

// @title Visitor Management System App
// @description Documentation api for VMS pindad
// @version 1.0.0
// @host localhost:8080
// @BasePath /api/vms/v1
// @securityDefinitions.apikey api_key
// @in header
// @name Authorization
func main() {
	// Create a new Gin router
	router := gin.Default()
	userService := domain.UserDomain{}

	userController := controller.UserController{Service: &userService}
	// Define your routes
	router.GET("/", HomeHandler)
	routes.UserRouter(router.Group("/api/vms"), &userController)
	// Start the server
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func HomeHandler(c *gin.Context) {
	test := "IKI TEST BAE"
	c.String(200, "Hello, World! "+test)
}
