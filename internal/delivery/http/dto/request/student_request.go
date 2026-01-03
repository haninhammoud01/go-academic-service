// 4. DTOs - REQUEST
// File: internal/delivery/http/dto/request/student_request.go

package request

import "time"

type CreateStudentRequest struct {
	NIM            string     `json:"nim" binding:"required"`
	Name           string     `json:"name" binding:"required"`
	Email          string     `json:"email" binding:"required,email"`
	Phone          string     `json:"phone"`
	Address        string     `json:"address"`
	DateOfBirth    *time.Time `json:"date_of_birth"`
	Gender         string     `json:"gender" binding:"oneof=male female"`
	Major          string     `json:"major" binding:"required"`
	EnrollmentYear int        `json:"enrollment_year" binding:"required,min=2000"`
	Status         string     `json:"status" binding:"oneof=active inactive graduated dropped"`
}

type UpdateStudentRequest struct {
	Name           string     `json:"name"`
	Email          string     `json:"email" binding:"omitempty,email"`
	Phone          string     `json:"phone"`
	Address        string     `json:"address"`
	DateOfBirth    *time.Time `json:"date_of_birth"`
	Gender         string     `json:"gender" binding:"omitempty,oneof=male female"`
	Major          string     `json:"major"`
	EnrollmentYear int        `json:"enrollment_year" binding:"omitempty,min=2000"`
	Status         string     `json:"status" binding:"omitempty,oneof=active inactive graduated dropped"`
	GPA            float64    `json:"gpa" binding:"omitempty,min=0,max=4"`
}
