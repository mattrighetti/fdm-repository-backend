package database

import (
	"database/sql"
	"fmt"
	"github.com/mattrighetti/fdm-repository-backend/models"
	"strings"
)

//GetAllModels Fetch all user data
func GetAllModels(m *[]models.FloodModel) error {
	return queryWithContext(
		"SELECT id, name, acronym, version, cod, soa, floodtypei, floodtypeii, modeltypei, modeltypeii, modeltypeiii, eis FROM models",
		func(rows *sql.Rows) error {
			var err error
			defer rows.Close()
			for rows.Next() {
				var model models.FloodModel
				err = rows.Scan(
					model.ID,
					model.Name,
					model.Acronym,
					model.Version,
					model.Cod,
					model.Soa,
					model.Floodtypei,
					model.Floodtypeii,
					model.Modeltypei,
					model.Modeltypeii,
					model.Modeltypeiii,
					model.Eis,
				)
				if err != nil {
					return fmt.Errorf("could not scan row: %v", err)
				}
				*m = append(*m, model)
			}
			return nil
		})
}

func GetFilters(filters *[]string, name string) error {
	return queryWithContext(
		"SELECT DISTINCT ? FROM models",
		func(rows *sql.Rows) error {
			defer rows.Close()
			var tmp string
			for rows.Next() {
				if err := rows.Scan(&tmp); err != nil {
					return nil
				}

				if tmp == "" {
					continue
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
		},
		name)
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

func GetFiltersMap(filtersMap map[string]*[]string) error {
	for k := range filtersMap {
		err := GetFilters(filtersMap[k], k)
		if err != nil {
			return err
		}
	}
	return nil
}

//GetModelByID returns a model with the given ID
func GetModelByID(model *models.FloodModel, id string) error {
	return queryWithContext(
		"SELECT * FROM user WHERE id=?",
		func(rows *sql.Rows) error {
			var err error
			defer rows.Close()
			for rows.Next() {
				var model models.FloodModel
				err = rows.Scan(
					model.ID,
					model.Name,
					model.Acronym,
					model.Version,
					model.Cod,
					model.Soa,
					model.Floodtypei,
					model.Floodtypeii,
					model.Modeltypei,
					model.Modeltypeii,
					model.Modeltypeiii,
					model.Eis,
				)
				if err != nil {
					return fmt.Errorf("could not get model by id: %v", err)
				}
			}
			return nil
		},
		id)
}

// GetMissingModels fetches all missing models from database
func GetMissingModels(missingModels *[]models.MissingModel) error {
	return queryWithContext(
		"SELECT * FROM missing_models",
		func(rows *sql.Rows) error {
			var err error
			defer rows.Close()
			for rows.Next() {
				var model models.MissingModel
				err = rows.Scan(&model.ID, &model.Name, &model.Bibliography)
				if err != nil {
					return fmt.Errorf("could not query missing models: %v", err)
				}
				*missingModels = append(*missingModels, model)
			}
			return nil
		})
}
