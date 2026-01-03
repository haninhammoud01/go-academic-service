// 6. AUTH DTOs - REQUEST
// File: internal/delivery/http/dto/request/auth_request.go

package request

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=admin staff student"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
