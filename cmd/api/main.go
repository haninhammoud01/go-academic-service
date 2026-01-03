package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/haninhammoud01/go-academic-service/internal/config"
	"github.com/haninhammoud01/go-academic-service/internal/delivery/http/handler"
	"github.com/haninhammoud01/go-academic-service/internal/delivery/http/middleware"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
	"github.com/haninhammoud01/go-academic-service/internal/pkg/jwt"
	postgresRepo "github.com/haninhammoud01/go-academic-service/internal/repository/postgres"
	"github.com/haninhammoud01/go-academic-service/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @title Go Academic Service API
// @version 1.0
// @description Academic Management Service REST API
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	if err := runMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Initialize JWT Service
	jwtService := jwt.NewJWTService(cfg.JWT.Secret, cfg.JWT.Expired)

	// Initialize Repositories
	userRepo := postgresRepo.NewUserRepository(db)
	studentRepo := postgresRepo.NewStudentRepository(db)
	lecturerRepo := postgresRepo.NewLecturerRepository(db)

	// Initialize Use Cases
	authUseCase := usecase.NewAuthUseCase(userRepo, jwtService)
	studentUseCase := usecase.NewStudentUseCase(studentRepo)
	lecturerUseCase := usecase.NewLecturerUseCase(lecturerRepo)

	// Initialize Handlers
	authHandler := handler.NewAuthHandler(authUseCase)
	studentHandler := handler.NewStudentHandler(studentUseCase)
	lecturerHandler := handler.NewLecturerHandler(lecturerUseCase)

	// Initialize Middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Service is running"})
	})

	// API v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(authMiddleware.Authenticate())
		{
			// Students routes
			students := protected.Group("/students")
			{
				students.POST("", authMiddleware.RequireRole("admin", "staff"), studentHandler.Create)
				students.GET("", studentHandler.GetAll)
				students.GET("/:id", studentHandler.GetByID)
				students.PUT("/:id", authMiddleware.RequireRole("admin", "staff"), studentHandler.Update)
				students.DELETE("/:id", authMiddleware.RequireRole("admin"), studentHandler.Delete)
			}

			// Lecturers routes
			lecturers := protected.Group("/lecturers")
			{
				lecturers.POST("", authMiddleware.RequireRole("admin", "staff"), lecturerHandler.Create)
				lecturers.GET("", lecturerHandler.GetAll)
				lecturers.GET("/:id", lecturerHandler.GetByID)
				lecturers.PUT("/:id", authMiddleware.RequireRole("admin", "staff"), lecturerHandler.Update)
				lecturers.DELETE("/:id", authMiddleware.RequireRole("admin"), lecturerHandler.Delete)
			}
		}
	}

	addr := fmt.Sprintf(":%s", cfg.App.Port)
	log.Printf("Starting %s on %s", cfg.App.Name, addr)
	log.Println("")
	log.Println("📚 Available Endpoints:")
	log.Println("   GET    /health")
	log.Println("")
	log.Println("🔐 Authentication (Public):")
	log.Println("   POST   /api/v1/auth/register")
	log.Println("   POST   /api/v1/auth/login")
	log.Println("")
	log.Println("👥 Students (Protected):")
	log.Println("   POST   /api/v1/students          [admin, staff]")
	log.Println("   GET    /api/v1/students          [authenticated]")
	log.Println("   GET    /api/v1/students/:id      [authenticated]")
	log.Println("   PUT    /api/v1/students/:id      [admin, staff]")
	log.Println("   DELETE /api/v1/students/:id      [admin]")
	log.Println("")
	log.Println("👨‍🏫 Lecturers (Protected):")
	log.Println("   POST   /api/v1/lecturers         [admin, staff]")
	log.Println("   GET    /api/v1/lecturers         [authenticated]")
	log.Println("   GET    /api/v1/lecturers/:id     [authenticated]")
	log.Println("   PUT    /api/v1/lecturers/:id     [admin, staff]")
	log.Println("   DELETE /api/v1/lecturers/:id     [admin]")

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.Database.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	log.Println("Database connected successfully")
	return db, nil
}

func runMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Student{},
		&entity.Lecturer{},
		&entity.Course{},
		&entity.Enrollment{},
	); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	log.Println("Migrations completed")
	return nil
}
