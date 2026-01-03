// File: internal/domain/entity/lecturer.go
package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lecturer struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	NIP            string         `gorm:"uniqueIndex;not null;size:20" json:"nip"`
	Name           string         `gorm:"not null;size:100" json:"name"`
	Email          string         `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Phone          string         `gorm:"size:20" json:"phone"`
	Address        string         `gorm:"type:text" json:"address"`
	DateOfBirth    *time.Time     `json:"date_of_birth,omitempty"`
	Gender         string         `gorm:"size:10;check:gender IN ('male', 'female')" json:"gender"`
	Department     string         `gorm:"not null;size:100" json:"department"`
	Position       string         `gorm:"size:50" json:"position"`
	Specialization string         `gorm:"size:100" json:"specialization"`
	EducationLevel string         `gorm:"size:50" json:"education_level"`
	Status         string         `gorm:"size:20;default:'active';check:status IN ('active', 'inactive', 'retired')" json:"status"`
	UserID         *uuid.UUID     `gorm:"type:uuid" json:"user_id,omitempty"`
	User           *User          `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL" json:"user,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Lecturer) TableName() string {
	return "lecturers"
}
