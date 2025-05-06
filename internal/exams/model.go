package exam

import (
	"time"

	"gorm.io/gorm"
)

// CodeSystem define os sistemas de codificação de exame
type CodeSystem string

const (
	CodeSystemLOINC CodeSystem = "LOINC"
	CodeSystemTUSS  CodeSystem = "TUSS"
	CodeSystemSUS   CodeSystem = "SUS"
)

// Exam representa um exame clínico associado a um paciente
// Inclui dados estruturados e referência a arquivos raw (texto e imagem/PDF)
type Exam struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID        uint      `gorm:"not null;index" json:"patient_id"`
	RequestingDoctor *string   `gorm:"size:100" json:"doctor,omitempty"`
	DateCollected    time.Time `gorm:"type:date;not null" json:"date_collected"`
	Key              string    `gorm:"size:100;not null" json:"key"`          // Nome completo do exame
	Label            *string   `gorm:"size:100" json:"label,omitempty"`       // Rótulo amigável
	Tags             *string   `gorm:"size:100" json:"tags,omitempty"`        // Ex: LipidProfile
	Abbreviation     *string   `gorm:"size:50" json:"abbreviation,omitempty"` // Ex: LDL

	Unit           *string     `gorm:"size:20" json:"unit,omitempty"`                 // Unidade de medida
	Method         *string     `gorm:"size:100" json:"method,omitempty"`              // Método de realização
	ReferenceRange *string     `gorm:"size:50" json:"reference_range,omitempty"`      // Intervalo de referência
	Date           *time.Time  `gorm:"type:date" json:"date,omitempty"`               // Data do exame
	RawText        *string     `gorm:"type:text" json:"raw_text,omitempty"`           // Texto bruto extraído (OCR)
	FileLink       *string     `gorm:"size:255" json:"file_link,omitempty"`           // URL/Path do arquivo de imagem ou PDF
	Code           *string     `gorm:"size:50" json:"code,omitempty"`                 // Código do exame
	CodeSystem     *CodeSystem `gorm:"type:varchar(20)" json:"code_system,omitempty"` // Sistema de codificação

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

// AnalitoResult representa cada “linha” do exame, ou seja, um analito e seu valor
type AnalitoResult struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	ExamID       uint    `gorm:"not null;index" json:"exam_id"`
	Name         string  `gorm:"size:100;not null" json:"name"`
	Abbreviation *string `gorm:"size:50" json:"abbreviation,omitempty"` // Ex: LDL
	// Para suportar valor numérico ou textual
	ValueString  *string  `gorm:"size:50" json:"value_string,omitempty"`
	ValueNumeric *float64 `json:"value_numeric,omitempty"`

	Unit     *string  `gorm:"size:20" json:"unit,omitempty"`
	MinValue *float64 `json:"min_value,omitempty"`
	MaxValue *float64 `json:"max_value,omitempty"`
	UnitRef  *string  `gorm:"size:20" json:"unit_ref,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
