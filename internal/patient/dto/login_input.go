package dto

// LoginInput é o payload esperado no endpoint POST /users/login
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
