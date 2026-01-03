// 3. USE CASE
// File: internal/usecase/student_usecase.go

package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
	"github.com/haninhammoud01/go-academic-service/internal/domain/repository"
	"gorm.io/gorm"
)

type StudentUseCase interface {
	Create(ctx context.Context, student *entity.Student) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Student, error)
	GetAll(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*entity.Student, int64, error)
	Update(ctx context.Context, id uuid.UUID, student *entity.Student) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type studentUseCaseImpl struct {
	repo repository.StudentRepository
}

func NewStudentUseCase(repo repository.StudentRepository) StudentUseCase {
	return &studentUseCaseImpl{repo: repo}
}

func (uc *studentUseCaseImpl) Create(ctx context.Context, student *entity.Student) error {
	// Check if NIM already exists
	existing, err := uc.repo.FindByNIM(ctx, student.NIM)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check existing NIM: %w", err)
	}
	if existing != nil {
		return errors.New("NIM already exists")
	}

	// Validate required fields
	if student.NIM == "" || student.Name == "" || student.Email == "" || student.Major == "" {
		return errors.New("required fields are missing")
	}

	return uc.repo.Create(ctx, student)
}

func (uc *studentUseCaseImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Student, error) {
	student, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}
	return student, nil
}

func (uc *studentUseCaseImpl) GetAll(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*entity.Student, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return uc.repo.FindAll(ctx, page, pageSize, filters)
}

func (uc *studentUseCaseImpl) Update(ctx context.Context, id uuid.UUID, student *entity.Student) error {
	// Check if student exists
	existing, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("student not found")
		}
		return err
	}

	// Update fields
	student.ID = existing.ID
	return uc.repo.Update(ctx, student)
}

func (uc *studentUseCaseImpl) Delete(ctx context.Context, id uuid.UUID) error {
	// Check if student exists
	_, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("student not found")
		}
		return err
	}

	return uc.repo.Delete(ctx, id)
}
