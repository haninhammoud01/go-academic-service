// File: internal/domain/entity/enrollment.go
package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID                   uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	StudentID            uuid.UUID      `gorm:"type:uuid;not null" json:"student_id"`
	Student              *Student       `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE" json:"student,omitempty"`
	CourseID             uuid.UUID      `gorm:"type:uuid;not null" json:"course_id"`
	Course               *Course        `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE" json:"course,omitempty"`
	AcademicYear         string         `gorm:"not null;size:10" json:"academic_year"`
	Semester             int            `gorm:"not null;check:semester > 0" json:"semester"`
	EnrollmentDate       time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"enrollment_date"`
	Status               string         `gorm:"size:20;default:'enrolled';check:status IN ('enrolled', 'completed', 'dropped', 'failed')" json:"status"`
	Grade                string         `gorm:"size:2;check:grade IN ('A', 'AB', 'B', 'BC', 'C', 'D', 'E')" json:"grade,omitempty"`
	Score                *float64       `gorm:"type:decimal(5,2)" json:"score,omitempty"`
	AttendancePercentage *float64       `gorm:"type:decimal(5,2)" json:"attendance_percentage,omitempty"`
	Remarks              string         `gorm:"type:text" json:"remarks,omitempty"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Enrollment) TableName() string {
	return "enrollments"
}
