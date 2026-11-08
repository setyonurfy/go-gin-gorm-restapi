package controller

import (
	"go-gin-gorm/config"
	"go-gin-gorm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create
func CreateUser(c *gin.Context) {
	var users model.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// get All
func GetAll(c *gin.Context) {
	var users []model.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})

}

// get by id
func GetById(c *gin.Context) {
	var users model.User
	if err := config.DB.First(&users, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// update
func UpdateUser(c *gin.Context) {
	var user model.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = 0 // cegah client mengubah ID

	if err := config.DB.Model(&user).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// delete
func DeleteUser(c *gin.Context) {
	var users model.User
	if err := config.DB.First(&users, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	config.DB.Delete(&users)
	c.JSON(http.StatusOK, gin.H{"Message": "deleted successfully!"})
}
