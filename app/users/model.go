package users

import (
	"System/System"
	"System/exchanges/user"
	"System/models"
)

func getUserByEmail(email string) (bool, models.User) {
	row := models.User{}
	System.DB.First(&row, "email = ? ", email)
	if row.ID == 0 {
		return false, models.User{}
	}
	return true, row
}

func updateRow(row *models.User) {
	System.DB.Save(&row)
}

func store(req *user.SignupRequest) *models.User {
	row := setHelper(req)
	System.DB.Create(&row)
	return row
}
