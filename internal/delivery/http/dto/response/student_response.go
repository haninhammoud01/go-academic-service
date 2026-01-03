// 5. DTOs - RESPONSE
// File: internal/delivery/http/dto/response/student_response.go

package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
)

type StudentResponse struct {
	ID             uuid.UUID  `json:"id"`
	NIM            string     `json:"nim"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone,omitempty"`
	Address        string     `json:"address,omitempty"`
	DateOfBirth    *time.Time `json:"date_of_birth,omitempty"`
	Gender         string     `json:"gender,omitempty"`
	Major          string     `json:"major"`
	EnrollmentYear int        `json:"enrollment_year"`
	Status         string     `json:"status"`
	GPA            float64    `json:"gpa"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type StudentListResponse struct {
	Data       []StudentResponse `json:"data"`
	Pagination PaginationMeta    `json:"pagination"`
}

type PaginationMeta struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

func ToStudentResponse(student *entity.Student) StudentResponse {
	return StudentResponse{
		ID:             student.ID,
		NIM:            student.NIM,
		Name:           student.Name,
		Email:          student.Email,
		Phone:          student.Phone,
		Address:        student.Address,
		DateOfBirth:    student.DateOfBirth,
		Gender:         student.Gender,
		Major:          student.Major,
		EnrollmentYear: student.EnrollmentYear,
		Status:         student.Status,
		GPA:            student.GPA,
		CreatedAt:      student.CreatedAt,
		UpdatedAt:      student.UpdatedAt,
	}
}
