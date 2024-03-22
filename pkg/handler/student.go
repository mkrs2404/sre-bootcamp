package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkrs2404/sre-bootcamp/pkg/config"
	"github.com/mkrs2404/sre-bootcamp/pkg/errors"
	"github.com/mkrs2404/sre-bootcamp/pkg/middleware"
	"github.com/mkrs2404/sre-bootcamp/pkg/repo/student"
	"github.com/mkrs2404/sre-bootcamp/pkg/service"
)

type StudentHandler interface {
	RegisterRoutes(router *gin.RouterGroup)
	getAll(c *gin.Context)
	getByID(c *gin.Context)
	add(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
}

type studentHandler struct {
	studentService service.StudentService
}

// NewStudentHandler creates a new AnimalHandler.
func NewStudentHandler(ss service.StudentService) StudentHandler {
	return &studentHandler{
		studentService: ss,
	}
}

// RegisterRoutes registers the API routes with a Gin router.
func (h *studentHandler) RegisterRoutes(router *gin.RouterGroup) {
	apiKey := config.Get().APIKey

	// route to get all students
	router.GET("/student", h.getAll)

	// route to get a single student
	router.GET("/students/:id", h.getByID)

	// route to add a new student
	router.POST("/students", middleware.Auth(apiKey), h.add)

	// route to update an existing student
	router.PUT("/students/:id", middleware.Auth(apiKey), h.update)

	// route to delete an existing student
	router.DELETE("/students/:id", middleware.Auth(apiKey), h.delete)
}

// getAll handles a request to get all students.
// @Summary Get students
// @Description Get all students
// @Tags Students
// @Produce application/json
// @Success 200 {object} []student.Student
// @Failure 400 {object} errors.HTTPError
// @Failure 401 {object} errors.HTTPError
// @Failure 409 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/students [get]
func (h *studentHandler) getAll(c *gin.Context) {
	students, err := h.studentService.GetAll()
	if err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}

	c.JSON(http.StatusOK, students)
}

// getByID handles a request to get a student by ID.
// @Summary Get student
// @Description Get one student by ID
// @Tags Students
// @Produce application/json
// @Success 200 {object} student.Student
// @Failure 400 {object} errors.HTTPError
// @Failure 401 {object} errors.HTTPError
// @Failure 409 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/students/{id} [get]
func (h *studentHandler) getByID(c *gin.Context) {
	studentID := c.Param("id")

	sID, err := strconv.ParseInt(studentID, 10, 32)
	if err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}

	students, err := h.studentService.GetByID(int(sID))
	if err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}

	c.JSON(http.StatusOK, students)
}

// add handles a request to add a new student.
// @Summary Add student
// @Description Add new student
// @Tags Students
// @Param "Authorization" header string true "Enter JSON Web Token (JWT)"
// @Produce application/json
// @Param data body student.Student true "Required data for adding new student"
// @Success 201 {object} student.Student
// @Failure 400 {object} errors.HTTPError
// @Failure 401 {object} errors.HTTPError
// @Failure 409 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/students [post]
func (h *studentHandler) add(c *gin.Context) {
	var student student.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}

	if err := h.studentService.Create(&student); err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}

	c.JSON(http.StatusCreated, student)
}

// update handles a request to update an existing student.
// @Summary Update student
// @Description Update existing student by id
// @Tags Students
// @Produce application/json
// @Param "Authorization" header string true "Enter JSON Web Token (JWT)"
// @Param id path string true "id"
// @Param data body student.Student true "Required data for updating student"
// @Success 200 {object} student.Student
// @Failure 400 {object} errors.HTTPError
// @Failure 401 {object} errors.HTTPError
// @Failure 409 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/students/{id} [put]
func (h *studentHandler) update(c *gin.Context) {
	studentID := c.Param("id")

	var student student.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
	}

	sID, err := strconv.ParseUint(studentID, 10, 32)
	if err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}
	student.ID = uint(sID)

	if err := h.studentService.Update(&student); err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}

	c.JSON(http.StatusOK, student)
}

// Delete handles a request to delete an existing student.
// @Summary Delete student
// @Description Delete existing student by id
// @Tags Students
// @Produce application/json
// @Param "Authorization" header string true "Enter JSON Web Token (JWT)"
// @Param id path string true "id"
// @Success 200 {object} student.Student
// @Failure 400 {object} errors.HTTPError
// @Failure 401 {object} errors.HTTPError
// @Failure 409 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/students/{id} [delete]
func (h *studentHandler) delete(c *gin.Context) {
	studentID := c.Param("id")

	sID, err := strconv.ParseUint(studentID, 10, 32)
	if err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}
	var student student.Student
	student.ID = uint(sID)

	if err := h.studentService.Delete(&student); err != nil {
		c.JSON(errors.StatusCode(err), errors.Response(err))
		return
	}

	c.JSON(http.StatusOK, student)
}
