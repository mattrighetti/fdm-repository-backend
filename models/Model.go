package models

import (
	"github.com/MattRighetti/fdm-repository-backend/config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(model *[]FloodModel) (err error) {
	if err = config.DB.Find(model).Error; err != nil {
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
