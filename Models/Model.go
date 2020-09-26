package Models

import (
	"github.com/MattRighetti/fdm-repository-backend/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(model *[]FloodModel) (err error) {
	if err = Config.DB.Find(model).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(model *FloodModel, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(model).Error; err != nil {
		return err
	}
	return nil
}
