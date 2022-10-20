package transformers

import (
	"System/exchanges/user"
	"System/models"
)

func TransformUserLogin(row *models.User) user.UserResponseLogin {
	return user.UserResponseLogin{
		ID:    row.ID,
		Name:  row.Name,
		Role:  row.Role,
		Email: row.Email,
		Token: row.Token,
	}
}

func TransformUserSignUp(row *models.User) user.UserResponseSignUp {
	return user.UserResponseSignUp{
		ID:    row.ID,
		Name:  row.Name,
		Role:  row.Role,
		Email: row.Email,
	}
}

func TranformJoin(join *models.StudentCourse) user.JoinResponse {
	return user.JoinResponse{
		UserID:   join.UserID,
		CourseID: join.CourseID,
	}
}
