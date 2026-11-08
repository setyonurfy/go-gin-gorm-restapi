package router

import (
	"go-gin-gorm/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) { // gin.Engine untuk router dan middleware
	// create
	r.POST("/users", controller.CreateUser)
	// get all
	r.GET("/users", controller.GetAll)
	// get by id
	r.GET("/users/:id", controller.GetById)
	// update
	r.PUT("/users/:id", controller.UpdateUser)
	// delete
	r.DELETE("/users/:id", controller.DeleteUser)
}
