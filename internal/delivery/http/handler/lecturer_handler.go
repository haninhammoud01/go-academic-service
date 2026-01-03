// File: internal/delivery/http/handler/lecturer_handler.go
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

type LecturerHandler struct {
	useCase usecase.LecturerUseCase
}

func NewLecturerHandler(useCase usecase.LecturerUseCase) *LecturerHandler {
	return &LecturerHandler{useCase: useCase}
}

func (h *LecturerHandler) Create(c *gin.Context) {
	var req request.CreateLecturerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err))
		return
	}

	lecturer := &entity.Lecturer{
		NIP:            req.NIP,
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		Department:     req.Department,
		Position:       req.Position,
		Specialization: req.Specialization,
		DateOfBirth:    req.DateOfBirth,
		Gender:         req.Gender,
		Status:         "active",
	}

	if err := h.useCase.Create(c.Request.Context(), lecturer); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Failed to create lecturer", err))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse("Lecturer created successfully", response.ToLecturerResponse(lecturer)))
}

func (h *LecturerHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid ID", err))
		return
	}

	lecturer, err := h.useCase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("Lecturer not found", err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Lecturer retrieved successfully", response.ToLecturerResponse(lecturer)))
}

func (h *LecturerHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	filters := make(map[string]interface{})
	if department := c.Query("department"); department != "" {
		filters["department"] = department
	}
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if search := c.Query("search"); search != "" {
		filters["search"] = search
	}

	lecturers, total, err := h.useCase.GetAll(c.Request.Context(), page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get lecturers", err))
		return
	}

	var lecturerResponses []response.LecturerResponse
	for _, lecturer := range lecturers {
		lecturerResponses = append(lecturerResponses, response.ToLecturerResponse(lecturer))
	}

	totalPage := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPage++
	}

	result := map[string]interface{}{
		"data": lecturerResponses,
		"pagination": response.PaginationMeta{
			Page:      page,
			PageSize:  pageSize,
			Total:     total,
			TotalPage: totalPage,
		},
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Lecturers retrieved successfully", result))
}

func (h *LecturerHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid ID", err))
		return
	}

	var req request.UpdateLecturerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request", err))
		return
	}

	lecturer := &entity.Lecturer{
		Name:           req.Name,
		Email:          req.Email,
		Phone:          req.Phone,
		Department:     req.Department,
		Position:       req.Position,
		Specialization: req.Specialization,
		DateOfBirth:    req.DateOfBirth,
		Gender:         req.Gender,
		Status:         req.Status,
	}

	if err := h.useCase.Update(c.Request.Context(), id, lecturer); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Failed to update lecturer", err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Lecturer updated successfully", response.ToLecturerResponse(lecturer)))
}

func (h *LecturerHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid ID", err))
		return
	}

	if err := h.useCase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse("Failed to delete lecturer", err))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Lecturer deleted successfully", nil))
}
