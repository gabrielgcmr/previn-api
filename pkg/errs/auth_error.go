package errs

import "errors"

var (
	ErrDuplicateEmail = errors.New("este e-mail já está em uso")
	ErrInvalidLogin   = errors.New("e-mail ou senha inválidos")
	ErrHashFailure    = errors.New("erro ao criptografar a senha")
	ErrCreateUser     = errors.New("não foi possível criar o usuário")
)
