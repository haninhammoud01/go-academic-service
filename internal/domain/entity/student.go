// File: internal/domain/entity/student.go
package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	NIM            string         `gorm:"uniqueIndex;not null;size:20" json:"nim"`
	Name           string         `gorm:"not null;size:100" json:"name"`
	Email          string         `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Phone          string         `gorm:"size:20" json:"phone"`
	Address        string         `gorm:"type:text" json:"address"`
	DateOfBirth    *time.Time     `json:"date_of_birth,omitempty"`
	Gender         string         `gorm:"size:10;check:gender IN ('male', 'female')" json:"gender"`
	Major          string         `gorm:"not null;size:100" json:"major"`
	EnrollmentYear int            `gorm:"not null" json:"enrollment_year"`
	Status         string         `gorm:"size:20;default:'active';check:status IN ('active', 'inactive', 'graduated', 'dropped')" json:"status"`
	GPA            float64        `gorm:"type:decimal(3,2);default:0.00" json:"gpa"`
	UserID         *uuid.UUID     `gorm:"type:uuid" json:"user_id,omitempty"`
	User           *User          `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL" json:"user,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Student) TableName() string {
	return "students"
}
