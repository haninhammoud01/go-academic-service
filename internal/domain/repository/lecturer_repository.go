// File: internal/domain/repository/lecturer_repository.go
package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
)

type LecturerRepository interface {
	Create(ctx context.Context, lecturer *entity.Lecturer) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Lecturer, error)
	FindAll(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*entity.Lecturer, int64, error)
	Update(ctx context.Context, lecturer *entity.Lecturer) error
	Delete(ctx context.Context, id uuid.UUID) error
}
