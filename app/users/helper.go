package users

import (
	"System/System"
	"System/exchanges/user"
	"System/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func buildToken(id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	Token, _ := token.SignedString([]byte("1o2m3a4r5z6i7z8o9"))
	return Token
}

func setHelper(req *user.SignupRequest) *models.User {
	return &models.User{
		Name:     req.UserName,
		Role:     req.Role,
		Email:    req.Email,
		Password: req.Password,
	}
}

func logoutHelper(row *models.User) {
	row.Token = ""
	System.DB.Save(&row)
}
