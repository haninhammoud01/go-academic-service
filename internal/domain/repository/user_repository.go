// 3. USER REPOSITORY INTERFACE
// File: internal/domain/repository/user_repository.go

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}
