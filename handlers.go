package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mattrighetti/fdm-repository-backend/database"

	"github.com/mattrighetti/fdm-repository-backend/models"
)

// getModels gets all users from database
func getModels(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var floodModel []models.FloodModel
	err := database.GetAllModels(&floodModel)
	if err != nil {
		sendStatus(w, http.StatusNotFound)
	} else {
		sendJSON(w, http.StatusOK, floodModel)
	}
}

// getNumModels gets the total number of models in the database
func getNumModels(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var numModels int64
	err := database.CountModels(&numModels)
	if err != nil {
		sendStatus(w, http.StatusInternalServerError)
	} else {
		sendJSON(w, http.StatusOK, numModels)
	}
}

// getFilters gets all filters from database
func getFilters(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	filtersMap := map[string]*[]string{
		"cod":          {},
		"soa":          {},
		"floodtypei":   {},
		"floodtypeii":  {},
		"modeltypei":   {},
		"modeltypeii":  {},
		"modeltypeiii": {},
		"eis":          {},
	}

	err := database.GetFiltersMap(filtersMap)
	if err != nil {
		sendStatus(w, http.StatusNotFound)
	} else {
		sendJSON(w, http.StatusOK, filtersMap)
	}
}

// getModelByID gets user by id
func getModelByID(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	var model models.FloodModel
	err := database.GetModelByID(&model, id)
	if err != nil {
		sendStatus(w, http.StatusNotFound)
	} else {
		sendJSON(w, http.StatusOK, model)
	}
}

// getMissingModels gets all missing models from database
func getMissingModels(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var missingModels []models.MissingModel
	err := database.GetMissingModels(&missingModels)
	if err != nil {
		sendStatus(w, http.StatusNotFound)
	} else {
		sendJSON(w, http.StatusOK, missingModels)
	}
}
