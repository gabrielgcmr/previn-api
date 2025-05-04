package dto

import "time"

// Response é o JSON que devolvemos ao cliente (sem senha!).
type Response struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	CPF       string    `json:"cpf"`
	CNS       *string   `json:"cns,omitempty"`
	Email     string    `json:"email"`
	Phone     *string   `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     *string   `json:"token,omitempty"` // se você emitir JWT no login/register
}
