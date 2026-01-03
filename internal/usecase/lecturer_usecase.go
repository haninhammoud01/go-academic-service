// File: internal/usecase/lecturer_usecase.go
package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
	"github.com/haninhammoud01/go-academic-service/internal/domain/repository"
	"gorm.io/gorm"
)

type LecturerUseCase interface {
	Create(ctx context.Context, lecturer *entity.Lecturer) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Lecturer, error)
	GetAll(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*entity.Lecturer, int64, error)
	Update(ctx context.Context, id uuid.UUID, lecturer *entity.Lecturer) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type lecturerUseCaseImpl struct {
	repo repository.LecturerRepository
}

func NewLecturerUseCase(repo repository.LecturerRepository) LecturerUseCase {
	return &lecturerUseCaseImpl{repo: repo}
}

func (uc *lecturerUseCaseImpl) Create(ctx context.Context, lecturer *entity.Lecturer) error {
	if lecturer.NIP == "" || lecturer.Name == "" || lecturer.Email == "" || lecturer.Department == "" {
		return errors.New("required fields are missing")
	}
	return uc.repo.Create(ctx, lecturer)
}

func (uc *lecturerUseCaseImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Lecturer, error) {
	lecturer, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lecturer not found")
		}
		return nil, err
	}
	return lecturer, nil
}

func (uc *lecturerUseCaseImpl) GetAll(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*entity.Lecturer, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return uc.repo.FindAll(ctx, page, pageSize, filters)
}

func (uc *lecturerUseCaseImpl) Update(ctx context.Context, id uuid.UUID, lecturer *entity.Lecturer) error {
	existing, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("lecturer not found")
		}
		return err
	}
	lecturer.ID = existing.ID
	return uc.repo.Update(ctx, lecturer)
}

func (uc *lecturerUseCaseImpl) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("lecturer not found")
		}
		return err
	}
	return uc.repo.Delete(ctx, id)
}
