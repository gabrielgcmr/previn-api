package patient

import (
	"time"
)

type Patient struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	CPF          string     `gorm:"size:14;not null;uniqueIndex" json:"cpf"`
	CNS          *string    `gorm:"size:15" json:"cns,omitempty"`
	FullName     string     `gorm:"size:255;not null" json:"full_name"`
	PasswordHash string     `gorm:"size:255;not null" json:"-"`
	Phone        *string    `gorm:"size:20" json:"phone,omitempty"`
	Email        *string    `gorm:"size:100;uniqueIndex" json:"email,omitempty"`
	CreatedAt    *time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
