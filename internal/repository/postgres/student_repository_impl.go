// 2. REPOSITORY IMPLEMENTATION
// File: internal/repository/postgres/student_repository_impl.go

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
	"github.com/haninhammoud01/go-academic-service/internal/domain/repository"
	"gorm.io/gorm"
)

type studentRepositoryImpl struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) repository.StudentRepository {
	return &studentRepositoryImpl{db: db}
}

func (r *studentRepositoryImpl) Create(ctx context.Context, student *entity.Student) error {
	return r.db.WithContext(ctx).Create(student).Error
}

func (r *studentRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Student, error) {
	var student entity.Student
	if err := r.db.WithContext(ctx).First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepositoryImpl) FindByNIM(ctx context.Context, nim string) (*entity.Student, error) {
	var student entity.Student
	if err := r.db.WithContext(ctx).First(&student, "nim = ?", nim).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepositoryImpl) FindAll(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*entity.Student, int64, error) {
	var students []*entity.Student
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Student{})

	// Apply filters
	if major, ok := filters["major"].(string); ok && major != "" {
		query = query.Where("major ILIKE ?", "%"+major+"%")
	}
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if search, ok := filters["search"].(string); ok && search != "" {
		query = query.Where("name ILIKE ? OR nim ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&students).Error; err != nil {
		return nil, 0, err
	}

	return students, total, nil
}

func (r *studentRepositoryImpl) Update(ctx context.Context, student *entity.Student) error {
	return r.db.WithContext(ctx).Save(student).Error
}

func (r *studentRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Student{}, "id = ?", id).Error
}
