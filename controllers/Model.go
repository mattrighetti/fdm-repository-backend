package controllers

import (
	"net/http"

	Models "github.com/MattRighetti/fdm-repository-backend/models"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.FloodModel
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.FloodModel
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
