package controllers

import (
	"net/http"

	"github.com/MattRighetti/fdm-repository-backend/models"

	"github.com/gin-gonic/gin"
)

//GetUsers gets all users from database
func GetUsers(c *gin.Context) {
	var floodModel []models.FloodModel
	err := models.GetAllUsers(&floodModel)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, floodModel)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.FloodModel
	err := models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
