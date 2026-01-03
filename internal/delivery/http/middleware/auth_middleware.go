// 9. AUTH MIDDLEWARE
// File: internal/delivery/http/middleware/auth_middleware.go

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/haninhammoud01/go-academic-service/internal/delivery/http/dto/response"
	"github.com/haninhammoud01/go-academic-service/internal/pkg/jwt"
)

type AuthMiddleware struct {
	jwtService *jwt.JWTService
}

func NewAuthMiddleware(jwtService *jwt.JWTService) *AuthMiddleware {
	return &AuthMiddleware{jwtService: jwtService}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse("Missing authorization header", nil))
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse("Invalid authorization format", nil))
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := m.jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse("Invalid or expired token", err))
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse("User role not found", nil))
			c.Abort()
			return
		}

		roleStr := userRole.(string)
		allowed := false
		for _, role := range roles {
			if roleStr == role {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, response.ErrorResponse("Insufficient permissions", nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
