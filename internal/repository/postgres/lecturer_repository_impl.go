// File: internal/repository/postgres/lecturer_repository_impl.go
package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
	"github.com/haninhammoud01/go-academic-service/internal/domain/repository"
	"gorm.io/gorm"
)

type lecturerRepositoryImpl struct {
	db *gorm.DB
}

func NewLecturerRepository(db *gorm.DB) repository.LecturerRepository {
	return &lecturerRepositoryImpl{db: db}
}

func (r *lecturerRepositoryImpl) Create(ctx context.Context, lecturer *entity.Lecturer) error {
	return r.db.WithContext(ctx).Create(lecturer).Error
}

func (r *lecturerRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Lecturer, error) {
	var lecturer entity.Lecturer
	if err := r.db.WithContext(ctx).First(&lecturer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &lecturer, nil
}

func (r *lecturerRepositoryImpl) FindAll(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*entity.Lecturer, int64, error) {
	var lecturers []*entity.Lecturer
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Lecturer{})

	if department, ok := filters["department"].(string); ok && department != "" {
		query = query.Where("department ILIKE ?", "%"+department+"%")
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if search, ok := filters["search"].(string); ok && search != "" {
		query = query.Where("name ILIKE ? OR nip ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&lecturers).Error; err != nil {
		return nil, 0, err
	}

	return lecturers, total, nil
}

func (r *lecturerRepositoryImpl) Update(ctx context.Context, lecturer *entity.Lecturer) error {
	return r.db.WithContext(ctx).Save(lecturer).Error
}

func (r *lecturerRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Lecturer{}, "id = ?", id).Error
}
