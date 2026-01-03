// File: internal/domain/entity/course.go
package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Code        string         `gorm:"uniqueIndex;not null;size:20" json:"code"`
	Name        string         `gorm:"not null;size:200" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Credits     int            `gorm:"not null;check:credits > 0" json:"credits"`
	Semester    int            `gorm:"not null;check:semester > 0" json:"semester"`
	Department  string         `gorm:"not null;size:100" json:"department"`
	CourseType  string         `gorm:"size:50;check:course_type IN ('mandatory', 'elective')" json:"course_type"`
	MaxStudents int            `gorm:"default:40" json:"max_students"`
	LecturerID  *uuid.UUID     `gorm:"type:uuid" json:"lecturer_id,omitempty"`
	Lecturer    *Lecturer      `gorm:"foreignKey:LecturerID;constraint:OnDelete:SET NULL" json:"lecturer,omitempty"`
	Status      string         `gorm:"size:20;default:'active';check:status IN ('active', 'inactive')" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Course) TableName() string {
	return "courses"
}
