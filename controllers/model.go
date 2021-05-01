package controllers

import (
	"github.com/mattrighetti/fdm-repository-backend/database"
	"net/http"

	"github.com/mattrighetti/fdm-repository-backend/models"

	"github.com/gin-gonic/gin"
)

//GetModels gets all users from database
func GetModels(c *gin.Context) {
	var floodModel []models.FloodModel
	err := database.GetAllModels(&floodModel)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, floodModel)
	}
}

//GetFilters get all filters from database
func GetFilters(c *gin.Context) {
	filtersMap := map[string]*[]string{
		"cod": {},
		"soa": {},
		"floodtypei": {},
		"floodtypeii": {},
		"modeltypei": {},
		"modeltypeii": {},
		"modeltypeiii": {},
		"eis": {},
	}

	err := database.GetFiltersMap(filtersMap)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, filtersMap)
	}
}

//GetModelByID Get user by id
func GetModelByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var model models.FloodModel
	err := database.GetModelByID(&model, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, model)
	}
}

// GetMissingModels gets all missing models from database
func GetMissingModels(c *gin.Context) {
	var missingModels []models.MissingModel
	err := database.GetMissingModels(&missingModels)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, missingModels)
	}
}
