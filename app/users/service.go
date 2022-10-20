package users

import (
	"System/exchanges/user"
	"System/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func signupService(req *user.SignupRequest) *models.User {
	// line
	req.Password, _ = HashPassword(req.Password)
	return store(req)

}
func loginService(row *models.User) {
	row.Token = buildToken(row.ID)
	updateRow(row)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func logoutService(row *models.User) {
	logoutHelper(row)
}
