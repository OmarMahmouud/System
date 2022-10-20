package auth

import (
	"System/System"
	"System/models"
)

func getUserByToken(Token string) (bool, *models.User) {
	if Token == "" {
		return false, nil
	}
	var row models.User
	System.DB.Where("token = ?", Token).First(&row)
	if row.ID == 0 {
		return false, nil
	}
	return true, &row
}
