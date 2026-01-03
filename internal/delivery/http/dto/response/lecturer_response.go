// File: internal/delivery/http/dto/response/lecturer_response.go
package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
)

type LecturerResponse struct {
	ID             uuid.UUID `json:"id"`
	NIP            string    `json:"nip"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone,omitempty"`
	Department     string    `json:"department"`
	Position       string    `json:"position,omitempty"`
	Specialization string    `json:"specialization,omitempty"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

func ToLecturerResponse(lecturer *entity.Lecturer) LecturerResponse {
	return LecturerResponse{
		ID:             lecturer.ID,
		NIP:            lecturer.NIP,
		Name:           lecturer.Name,
		Email:          lecturer.Email,
		Phone:          lecturer.Phone,
		Department:     lecturer.Department,
		Position:       lecturer.Position,
		Specialization: lecturer.Specialization,
		Status:         lecturer.Status,
		CreatedAt:      lecturer.CreatedAt,
	}
}
