package dto

// RegisterInput payload para criar um usuário
type RegisterInput struct {
	FullName string `json:"full_name"  validate:"required"`
	CPF      string `json:"cpf"        validate:"required,len=14"`
	CNS      string `json:"cns,omitempty" validate:"omitempty,len=15"`
	Email    string `json:"email"      validate:"required,email"`
	Password string `json:"password"   validate:"required,min=8"`
	Phone    string `json:"phone,omitempty" validate:"omitempty"`
}
