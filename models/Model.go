package models

import (
	"fmt"
	"github.com/MattRighetti/fdm-repository-backend/config"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//GetAllUsers Fetch all user data
func GetAllModels(model *[]FloodModel) (err error) {
	if err = config.DB.Find(model).Error; err != nil {
		return err
	}
	return nil
}

func GetFilters(filters *[]string, name string) (err error) {
	rows, err := config.BareDB.Query(fmt.Sprintf("SELECT DISTINCT %s FROM models", name))
	if err != nil {
		return err
	}

	defer rows.Close()

	var tmp string
	for rows.Next() {
		if err := rows.Scan(&tmp); err != nil {
			return nil
		}


		if strings.Contains(tmp, "/") && name != "cod" {
			stringSlice := strings.Split(tmp, "/")
			for _, v := range stringSlice {
				if _, found := find(stringSlice, v); !found {
					*filters = append(*filters, v)
				}
			}
		} else {
			*filters = append(*filters, tmp)
		}
	}

	return nil
}

// find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func GetFiltersMap(filtersMap map[string]*[]string) (err error) {
	for k := range filtersMap {
		err := GetFilters(filtersMap[k], k)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAllModelsNoMarkdown(model *[]FloodModelNoDescription) (err error) {
	if err = config.DB.Select("id, name, acronym, version, cod, soa, floodtypei, floodtypeii, modeltypei, modeltypeii, modeltypeiii, eis").Find(model).Error; err != nil {
		return err
	}

	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(model *FloodModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(model).Error; err != nil {
		return err
	}
	return nil
}

// GetMissingModels fetches all missing models from database
func GetMissingModels(missingModels *[]MissingModel) (err error) {
	if err = config.DB.Find(missingModels).Error; err != nil {
		return err
	}

	return nil
}
