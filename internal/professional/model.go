package professional

import (
	"time"
)

// Profession representa o papel do usuário no sistema.
type Profession string

const (
	ProfessionMedico     Profession = "medico"
	ProfessionEnfermeiro Profession = "enfermeiro"
	ProfessionACS        Profession = "acs" // agente comunitário de saúde
	ProfessionTecEnf     Profession = "tecnico_enfermagem"
)

type Professional struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	CPF          string     `gorm:"size:14;not null;uniqueIndex" json:"cpf"` // ex: “123.456.789-00”
	CNS          *string    `gorm:"size:15" json:"cns,omitempty"`
	FullName     string     `gorm:"size:255;not null" json:"full_name"`
	PasswordHash string     `gorm:"size:255;not null" json:"-"`
	Profession   Profession `gorm:"type:varchar(30);not null" json:"profession"`
	Phone        *string    `gorm:"size:20" json:"phone,omitempty"`              // Telefone de contato
	Email        *string    `gorm:"size:100;uniqueIndex" json:"email,omitempty"` // E-mail de contato
	CreatedAt    *time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
