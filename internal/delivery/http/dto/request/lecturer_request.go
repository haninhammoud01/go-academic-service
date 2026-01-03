// File: internal/delivery/http/dto/request/lecturer_request.go
package request

import "time"

type CreateLecturerRequest struct {
	NIP            string     `json:"nip" binding:"required"`
	Name           string     `json:"name" binding:"required"`
	Email          string     `json:"email" binding:"required,email"`
	Phone          string     `json:"phone"`
	Department     string     `json:"department" binding:"required"`
	Position       string     `json:"position"`
	Specialization string     `json:"specialization"`
	DateOfBirth    *time.Time `json:"date_of_birth"`
	Gender         string     `json:"gender" binding:"omitempty,oneof=male female"`
}

type UpdateLecturerRequest struct {
	Name           string     `json:"name"`
	Email          string     `json:"email" binding:"omitempty,email"`
	Phone          string     `json:"phone"`
	Department     string     `json:"department"`
	Position       string     `json:"position"`
	Specialization string     `json:"specialization"`
	DateOfBirth    *time.Time `json:"date_of_birth"`
	Gender         string     `json:"gender" binding:"omitempty,oneof=male female"`
	Status         string     `json:"status" binding:"omitempty,oneof=active inactive retired"`
}
