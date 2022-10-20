package user

type JoinResponse struct {
	UserID   uint `json:"user_id" `
	CourseID uint `json:"course_id"`
}

type Student struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type UserResponseSignUp struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

type UserResponseLogin struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Token string `json:"token" `
}
