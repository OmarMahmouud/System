package user

type LoginRequest struct {
	Email    string `json:"email" validate:"required , email" `
	Password string `json:"password" validate:"required , min=6"`
}

type SignupRequest struct {
	Email    string `json:"email" validate:"required , email" `
	Password string `json:"password" validate:"required , min=6"`
	UserName string `json:"user_name" validate:"required"`
	Role     string `json:"role" validate:"eq=Student"`
}

type Join struct {
	UserID   uint `json:"user_id" `
	CourseID uint `json:"course_id"`
}
type LogoutRequest struct {
	Id uint `json:"id"`
}
