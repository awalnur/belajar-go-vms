package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define your routes
	router.GET("/", HomeHandler)

	// Start the server
	router.Run(":8080")
}

func HomeHandler(c *gin.Context) {
	test := "IKI TEST BAE"
	c.String(200, "Hello, World! "+test)
}
