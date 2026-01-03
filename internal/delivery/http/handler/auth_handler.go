// 8. AUTH HANDLER
// File: internal/delivery/http/handler/auth_handler.go

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haninhammoud01/go-academic-service/internal/delivery/http/dto/request"
	"github.com/haninhammoud01/go-academic-service/internal/delivery/http/dto/response"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
	"github.com/haninhammoud01/go-academic-service/internal/usecase"
)

type AuthHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

// Register godoc
// @Summary Register new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body request.RegisterRequest true "User data"
// @Success 201 {object} response.BaseResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err))
		return
	}

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
	}

	if err := h.authUseCase.Register(c.Request.Context(), user, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Failed to register", err))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse("User registered successfully", response.ToUserResponse(user)))
}

// Login godoc
// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body request.LoginRequest true "Login credentials"
// @Success 200 {object} response.BaseResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err))
		return
	}

	token, user, err := h.authUseCase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse("Login failed", err))
		return
	}

	authResp := response.AuthResponse{
		Token: token,
		User:  response.ToUserResponse(user),
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Login successful", authResp))
}
