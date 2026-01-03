// 1. REPOSITORY INTERFACE
// File: internal/domain/repository/student_repository.go

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
)

type StudentRepository interface {
	Create(ctx context.Context, student *entity.Student) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Student, error)
	FindByNIM(ctx context.Context, nim string) (*entity.Student, error)
	FindAll(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*entity.Student, int64, error)
	Update(ctx context.Context, student *entity.Student) error
	Delete(ctx context.Context, id uuid.UUID) error
}
