package patient

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// ----------------------
// Operações CRUD
// ----------------------

// Create: Cadastra um novo usuário
func (r *Repository) Create(user *Patient) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}
	return nil
}

// FindByID: Busca usuário por ID
func (r *Repository) FindByID(id int) (*Patient, error) {
	var user Patient
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}
	return &user, nil
}

// FindByCPF: Busca usuário por CPF (útil para login)
func (r *Repository) FindByCPF(cpf string) (*Patient, error) {
	var user Patient
	if err := r.db.Where("cpf = ?", cpf).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Busca usuário pelo email
func (r *Repository) FindByEmail(email string) (*Patient, error) {
	var user Patient
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update: Atualiza um usuário existente
func (r *Repository) Update(user *Patient) error {
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}
	return nil
}

// Delete: Remove um usuário (soft delete se você usar gorm.DeletedAt)
func (r *Repository) Delete(id int) error {
	if err := r.db.Delete(&Patient{}, id).Error; err != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}
	return nil
}

// ----------------------
// Métodos adicionais (opcionais)
// ----------------------

// Exemplo: Listar todos os usuários com paginação
func (r *Repository) List(limit, offset int) ([]Patient, error) {
	var users []Patient
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
