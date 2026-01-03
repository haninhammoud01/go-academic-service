// 7. HANDLER
// File: internal/delivery/http/handler/student_handler.go

package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/delivery/http/dto/request"
	"github.com/haninhammoud01/go-academic-service/internal/delivery/http/dto/response"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
	"github.com/haninhammoud01/go-academic-service/internal/usecase"
)

type StudentHandler struct {
	useCase usecase.StudentUseCase
}

func NewStudentHandler(useCase usecase.StudentUseCase) *StudentHandler {
	return &StudentHandler{useCase: useCase}
}

// Create godoc
// @Summary Create new student
// @Tags students
// @Accept json
// @Produce json
// @Param student body request.CreateStudentRequest true "Student data"
// @Success 201 {object} response.BaseResponse
// @Router /students [post]
func (h *StudentHandler) Create(c *gin.Context) {
	var req request.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err))
		return
	}

	student := &entity.Student{
		NIM:            req.NIM,
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		Address:        req.Address,
		DateOfBirth:    req.DateOfBirth,
		Gender:         req.Gender,
		Major:          req.Major,
		EnrollmentYear: req.EnrollmentYear,
		Status:         req.Status,
	}

	if student.Status == "" {
		student.Status = "active"
	}

	if err := h.useCase.Create(c.Request.Context(), student); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Failed to create student", err))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse("Student created successfully", response.ToStudentResponse(student)))
}

// GetByID godoc
// @Summary Get student by ID
// @Tags students
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} response.BaseResponse
// @Router /students/{id} [get]
func (h *StudentHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid student ID", err))
		return
	}

	student, err := h.useCase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("Student not found", err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Student retrieved successfully", response.ToStudentResponse(student)))
}

// GetAll godoc
// @Summary Get all students
// @Tags students
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param major query string false "Filter by major"
// @Param status query string false "Filter by status"
// @Param search query string false "Search by name or NIM"
// @Success 200 {object} response.BaseResponse
// @Router /students [get]
func (h *StudentHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	filters := make(map[string]interface{})
	if major := c.Query("major"); major != "" {
		filters["major"] = major
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}

	students, total, err := h.useCase.GetAll(c.Request.Context(), page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get students", err))
		return
	}

	var studentResponses []response.StudentResponse
	for _, student := range students {
		studentResponses = append(studentResponses, response.ToStudentResponse(student))
	}

	totalPage := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPage++
	}

	result := response.StudentListResponse{
		Data: studentResponses,
		Pagination: response.PaginationMeta{
			Page:      page,
			PageSize:  pageSize,
			Total:     total,
			TotalPage: totalPage,
		},
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Students retrieved successfully", result))
}

// Update godoc
// @Summary Update student
// @Tags students
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Param student body request.UpdateStudentRequest true "Student data"
// @Success 200 {object} response.BaseResponse
// @Router /students/{id} [put]
func (h *StudentHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid student ID", err))
		return
	}

	var req request.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err))
		return
	}

	student := &entity.Student{
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		Address:        req.Address,
		DateOfBirth:    req.DateOfBirth,
		Gender:         req.Gender,
		Major:          req.Major,
		EnrollmentYear: req.EnrollmentYear,
		Status:         req.Status,
		GPA:            req.GPA,
	}

	if err := h.useCase.Update(c.Request.Context(), id, student); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Failed to update student", err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Student updated successfully", response.ToStudentResponse(student)))
}

// Delete godoc
// @Summary Delete student
// @Tags students
// @Produce json
// @Param id path string true "Student ID"
// @Success 200 {object} response.BaseResponse
// @Router /students/{id} [delete]
func (h *StudentHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid student ID", err))
		return
	}

	if err := h.useCase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("Failed to delete student", err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Student deleted successfully", nil))
}
