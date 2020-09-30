package controllers

import (
	"net/http"

	"github.com/MattRighetti/fdm-repository-backend/models"

	"github.com/gin-gonic/gin"
)

//GetModels gets all users from database
func GetModels(c *gin.Context) {
	var floodModel []models.FloodModelNoDescription
	err := models.GetAllModelsNoMarkdown(&floodModel)

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

	err := models.GetFiltersMap(filtersMap)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, filtersMap)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var model models.FloodModel
	err := models.GetUserByID(&model, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, model)
	}
}

// GetMissingModels gets all missing models from database
func GetMissingModels(c *gin.Context) {
	var missingModels []models.MissingModel
	err := models.GetMissingModels(&missingModels)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, missingModels)
	}
}
