// 5. AUTH USE CASE
// File: internal/usecase/auth_usecase.go

package usecase

import (
	"context"
	"errors"

	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
	"github.com/haninhammoud01/go-academic-service/internal/domain/repository"
	"github.com/haninhammoud01/go-academic-service/internal/pkg/jwt"
	"github.com/haninhammoud01/go-academic-service/internal/pkg/password"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	Register(ctx context.Context, user *entity.User, plainPassword string) error
	Login(ctx context.Context, email, plainPassword string) (string, *entity.User, error)
}

type authUseCaseImpl struct {
	userRepo   repository.UserRepository
	jwtService *jwt.JWTService
}

func NewAuthUseCase(userRepo repository.UserRepository, jwtService *jwt.JWTService) AuthUseCase {
	return &authUseCaseImpl{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (uc *authUseCaseImpl) Register(ctx context.Context, user *entity.User, plainPassword string) error {
	// Check if email exists
	existing, err := uc.userRepo.FindByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existing != nil {
		return errors.New("email already exists")
	}

	// Check if username exists
	existing, err = uc.userRepo.FindByUsername(ctx, user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existing != nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := password.Hash(plainPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Set default values
	if user.Role == "" {
		user.Role = "student"
	}
	user.IsActive = true

	return uc.userRepo.Create(ctx, user)
}

func (uc *authUseCaseImpl) Login(ctx context.Context, email, plainPassword string) (string, *entity.User, error) {
	// Find user by email
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, errors.New("invalid email or password")
		}
		return "", nil, err
	}

	// Check if user is active
	if !user.IsActive {
		return "", nil, errors.New("user account is inactive")
	}

	// Verify password
	if !password.Verify(plainPassword, user.Password) {
		return "", nil, errors.New("invalid email or password")
	}

	// Generate token
	token, err := uc.jwtService.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
