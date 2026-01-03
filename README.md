# 🎓 Go Academic Service

**Production-ready REST API untuk Sistem Akademik dengan Clean Architecture, JWT Authentication & Role-Based Access Control**

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat)](https://gin-gonic.com/)

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [API Documentation](#api-documentation)
- [Database Schema](#database-schema)
- [Testing](#testing)
- [Security](#security)
- [Deployment](#deployment)
- [Project Structure](#project-structure)

---

## Overview

Go Academic Service adalah RESTful API untuk sistem manajemen akademik yang dibangun dengan Go (Golang) menggunakan Clean Architecture. Service ini menyediakan endpoint untuk mengelola data mahasiswa, dosen, mata kuliah, dan KRS dengan authentication dan authorization yang aman.

**Key Highlights:**

- Clean Architecture dengan 4-layer separation
- JWT Authentication dengan bcrypt password hashing
- Role-Based Access Control (Admin, Staff, Student)
- Complete CRUD operations dengan pagination & filtering
- PostgreSQL database dengan GORM
- Docker support untuk easy deployment

---

## Features

| Feature | Status | Description |
|---------|--------|-------------|
| JWT Authentication | Completed | Secure token-based authentication |
| Students CRUD | Completed | Complete dengan pagination & filtering |
| Lecturers CRUD | Completed | Department, position, specialization management |
| Courses Management | Database Ready | Schema tersedia untuk mata kuliah |
| Enrollments (KRS) | Database Ready | Schema untuk grade tracking |
| Role-Based Access | Completed | Admin, Staff, Student permissions |
| Advanced Filters | Completed | Search, pagination, sorting |
| Input Validation | Completed | Comprehensive request validation |

---

## Tech Stack

**Core Technologies:**

- **Language:** Go 1.21+
- **Web Framework:** Gin
- **Database:** PostgreSQL 15
- **ORM:** GORM
- **Authentication:** JWT (golang-jwt/jwt)
- **Security:** Bcrypt password hashing
- **Validation:** go-playground/validator
- **Containerization:** Docker & Docker Compose

**Key Libraries:**

```go
github.com/gin-gonic/gin           // Web framework
gorm.io/gorm                        // ORM
gorm.io/driver/postgres            // PostgreSQL driver
github.com/golang-jwt/jwt/v5       // JWT authentication
golang.org/x/crypto/bcrypt         // Password hashing
github.com/go-playground/validator // Request validation
github.com/joho/godotenv           // Environment variables
github.com/google/uuid             // UUID generation
```

---

## Architecture

Project ini menggunakan Clean Architecture dengan layer separation yang jelas:

```
┌─────────────────────────────────────┐
│   HTTP Layer (Handlers/DTOs)       │  ← API Endpoints & Request/Response
├─────────────────────────────────────┤
│   Use Case Layer (Business Logic)  │  ← Validation & Business Rules
├─────────────────────────────────────┤
│   Repository Layer (Database)      │  ← CRUD Operations & Queries
├─────────────────────────────────────┤
│   Domain Layer (Entities)          │  ← Core Models & Interfaces
└─────────────────────────────────────┘
```

**Layer Responsibilities:**

1. **Domain Layer** - Defines entities and repository interfaces
2. **Repository Layer** - Database operations implementation
3. **Use Case Layer** - Business logic and validation
4. **Delivery Layer** - HTTP handlers, DTOs, and middleware

**Benefits:**

- Testable and maintainable code
- Independent of external frameworks
- Easy to scale and modify
- Clear separation of concerns

---

## Quick Start

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher (or Docker)
- Git

### Installation

**1. Clone Repository**

```bash
git clone https://github.com/haninhammoud01/go-academic-service.git
cd go-academic-service
```

**2. Setup Environment**

```bash
cp .env.example .env
```

Edit `.env` file sesuai kebutuhan:

```env
APP_NAME=go-academic-service
APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=academic_db
DB_SSLMODE=disable

JWT_SECRET=your-super-secret-jwt-key-change-this
JWT_EXPIRED=24h
```

**3. Install Dependencies**

```bash
go mod download
go mod tidy
```

**4. Start Database**

Using Docker (Recommended):

```bash
docker-compose up -d postgres
```

Or use local PostgreSQL and create database:

```sql
CREATE DATABASE academic_db;
```
 
**5. Run Application**

```bash
go run cmd/api/main.go
```

Server will start at `http://localhost:8080`

**6. Verify Installation**

```bash
curl http://localhost:8080/health
# Expected: {"status":"ok","message":"Service is running"}
```

---

## API Documentation

### Base URL

```
http://localhost:8080/api/v1
```

### Authentication

All endpoints (except `/auth/*`) require JWT token in header:

```
Authorization: Bearer <your-jwt-token>
```

---

### Authentication Endpoints

#### Register User

```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "admin",
  "email": "admin@academic.com",
  "password": "admin123",
  "role": "admin"
}
```

**Available Roles:** `admin`, `staff`, `student`

#### Login

```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "admin@academic.com",
  "password": "admin123"
}
```

**Response:**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "uuid",
      "username": "admin",
      "email": "admin@academic.com",
      "role": "admin"
    }
  }
}
```

---

### Students Endpoints

#### Create Student

```http
POST /api/v1/students
Authorization: Bearer <token>
Content-Type: application/json

{
  "nim": "23001",
  "name": "Lin Patra",
  "email": "lin@student.ac.id",
  "phone": "08123456789",
  "gender": "male",
  "major": "Computer Science",
  "enrollment_year": 2024,
  "status": "active"
}
```

**Required Role:** `admin`, `staff`

#### Get All Students

```http
GET /api/v1/students?page=1&page_size=10&major=Computer%20Science&status=active&search=lin
Authorization: Bearer <token>
```

**Query Parameters:**

- `page` - Page number (default: 1)
- `page_size` - Items per page (default: 10, max: 100)
- `major` - Filter by major
- `status` - Filter by status (active, inactive, graduated, dropped)
- `search` - Search by name or NIM

#### Get Student by ID

```http
GET /api/v1/students/{id}
Authorization: Bearer <token>
```

#### Update Student

```http
PUT /api/v1/students/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Lin Patra Updated",
  "gpa": 3.75
}
```

**Required Role:** `admin`, `staff`

#### Delete Student

```http
DELETE /api/v1/students/{id}
Authorization: Bearer <token>
```

**Required Role:** `admin`

---

### Lecturers Endpoints

Similar structure to Students with the following endpoints:

```
POST   /api/v1/lecturers          [admin, staff]
GET    /api/v1/lecturers           [authenticated]
GET    /api/v1/lecturers/{id}      [authenticated]
PUT    /api/v1/lecturers/{id}      [admin, staff]
DELETE /api/v1/lecturers/{id}      [admin]
```

**Example Create Lecturer:**

```json
{
  "nip": "199001011",
  "name": "Dr. Jane Smith",
  "email": "jane@academic.com",
  "phone": "08123456789",
  "department": "Computer Science",
  "position": "Associate Professor",
  "specialization": "Artificial Intelligence",
  "gender": "female"
}
```

---

### Response Format

**Success Response:**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    // response data
  }
}
```

**Error Response:**

```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error description"
}
```

**Pagination Response:**

```json
{
  "success": true,
  "data": {
    "data": [...],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 50,
      "total_page": 5
    }
  }
}
```

---

## Database Schema

### Entity Relationship Diagram

```
┌─────────────┐       ┌──────────────┐       ┌─────────────┐
│    Users    │       │   Students   │       │  Lecturers  │
├─────────────┤       ├──────────────┤       ├─────────────┤
│ id (PK)     │◄─────┤│ id (PK)      │       │ id (PK)     │
│ username    │       │ user_id (FK) │       │ user_id (FK)│
│ email       │       │ nim          │       │ nip         │
│ password    │       │ name         │       │ name        │
│ role        │       │ email        │       │ email       │
│ is_active   │       │ major        │       │ department  │
└─────────────┘       │ gpa          │       │ position    │
                      └──────────────┘       └─────────────┘
                             │                      │
                             └──────────┬───────────┘
                                        │
                                  ┌─────────────┐
                                  │   Courses   │
                                  ├─────────────┤
                                  │ id (PK)     │
                                  │ code        │
                                  │ name        │
                                  │ credits     │
                                  │ lecturer_id │
                                  └─────────────┘
                                        │
                                        ▼
                                 ┌──────────────┐
                                 │ Enrollments  │
                                 ├──────────────┤
                                 │ id (PK)      │
                                 │ student_id   │
                                 │ course_id    │
                                 │ grade        │
                                 │ status       │
                                 └──────────────┘
```

### Tables Overview

**Users** - Authentication and role management  
**Students** - Student data with NIM, major, GPA tracking  
**Lecturers** - Lecturer data with department and specialization  
**Courses** - Course information with credits and semester  
**Enrollments** - Student-course relationship with grades (KRS)

---

## Testing

### Manual Testing

**Using Thunder Client / Postman:**

1. Register an admin user
2. Login to get JWT token
3. Add token to Authorization header
4. Test all CRUD endpoints

**Using cURL:**

```bash
# Register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","email":"admin@academic.com","password":"admin123","role":"admin"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@academic.com","password":"admin123"}'

# Create Student (replace YOUR_TOKEN)
curl -X POST http://localhost:8080/api/v1/students \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"nim":"23001","name":"Lin Patra","email":"lin@student.ac.id","major":"Computer Science","enrollment_year":2024}'
```

---

## Security

### Authentication & Authorization
- JWT token-based authentication
- Token expiration (configurable via JWT_EXPIRED)
- Role-based access control (RBAC)
- Middleware for route protection

### Data Protection
- Password hashing with bcrypt (cost: 10)
- Passwords never stored in plain text
- Sensitive data excluded from API responses
- SQL injection prevention via GORM parameterized queries

### Input Validation
- Request validation on all endpoints
- Type checking and format validation
- Business rule validation in use case layer

---

## Deployment

### Docker Deployment

```bash
# Build and run all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Production Checklist

- [ ] Change JWT_SECRET to a strong random string (min 32 characters)
- [ ] Set APP_ENV=production
- [ ] Use strong database password
- [ ] Enable HTTPS/TLS
- [ ] Configure proper CORS settings
- [ ] Add rate limiting middleware
- [ ] Setup monitoring and logging
- [ ] Regular database backups

---

## Project Structure

```
go-academic-service/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── domain/
│   │   ├── entity/                 # Domain entities
│   │   │   ├── user.go
│   │   │   ├── student.go
│   │   │   ├── lecturer.go
│   │   │   ├── course.go
│   │   │   └── enrollment.go
│   │   └── repository/             # Repository interfaces
│   │       ├── user_repository.go
│   │       ├── student_repository.go
│   │       └── lecturer_repository.go
│   ├── repository/
│   │   └── postgres/               # Repository implementations
│   │       ├── user_repository_impl.go
│   │       ├── student_repository_impl.go
│   │       └── lecturer_repository_impl.go
│   ├── usecase/                    # Business logic
│   │   ├── auth_usecase.go
│   │   ├── student_usecase.go
│   │   └── lecturer_usecase.go
│   ├── delivery/
│   │   └── http/
│   │       ├── handler/            # HTTP handlers
│   │       │   ├── auth_handler.go
│   │       │   ├── student_handler.go
│   │       │   └── lecturer_handler.go
│   │       ├── middleware/         # Middleware
│   │       │   └── auth_middleware.go
│   │       └── dto/                # Data Transfer Objects
│   │           ├── request/
│   │           └── response/
│   └── pkg/                        # Shared utilities
│       ├── jwt/                    # JWT helper
│       └── password/               # Password helper
├── database/
│   └── migrations/                 # SQL migrations
├── docs/
│   └── swagger/                    # API documentation
├── .env.example                    # Environment template
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Author

**Hanin Hammoud**

Email: 230411100005@student.trunojoyo.ac.id  
GitHub: [@haninhammoud01](https://github.com/haninhammoud01)

---

## Acknowledgments

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) by Robert C. Martin
- [Gin Web Framework](https://gin-gonic.com/) - High-performance HTTP web framework
- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- [Go Best Practices](https://go.dev/doc/effective_go)

---

## Support

If you have any questions or issues, please:
- Open an issue on GitHub
- Contact via email

---

**Built with Go**
