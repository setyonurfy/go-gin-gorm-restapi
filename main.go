package main

import (
	"go-gin-gorm/config"
	"go-gin-gorm/model"
	"go-gin-gorm/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set up routes
	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{})
	router.SetRouter(r)

	// Start the server
	r.Run(":8080")

}
