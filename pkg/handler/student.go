package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mkrs2404/sre-bootcamp/pkg/config"
	apierrors "github.com/mkrs2404/sre-bootcamp/pkg/errors"
	"github.com/mkrs2404/sre-bootcamp/pkg/middleware"
	"github.com/mkrs2404/sre-bootcamp/pkg/repo/student"
	"github.com/mkrs2404/sre-bootcamp/pkg/service"
	"github.com/sirupsen/logrus"
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

// NewStudentHandler creates a new StudentHandler.
func NewStudentHandler(ss service.StudentService) StudentHandler {
	return &studentHandler{
		studentService: ss,
	}
}

// RegisterRoutes registers the API routes with a Gin router.
func (h *studentHandler) RegisterRoutes(router *gin.RouterGroup) {
	apiKey := config.Get().APIKey

	// route to get all students
	router.GET("/students", h.getAll)

	// route to get a single student
	router.GET("/students/:id", h.getByID)

	// route to add a new student
	router.POST("/students", middleware.Auth(apiKey), h.add)

	// route to update an existing student
	router.PUT("/students/:id", middleware.Auth(apiKey), h.update)

	// route to delete an existing student
	router.DELETE("/students/:id", middleware.Auth(apiKey), h.delete)
}

func (h *studentHandler) getAll(c *gin.Context) {
	students, err := h.studentService.GetAll()
	if err != nil {
		c.JSON(apierrors.StatusCode(err), apierrors.Response(err))
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *studentHandler) getByID(c *gin.Context) {
	studentID := c.Param("id")

	sID, err := strconv.ParseInt(studentID, 10, 32)
	if err != nil {
		logrus.Errorf("error parsing the student id %s : %v", studentID, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": apierrors.ErrInvalidID.Error(),
		})
		return
	}

	students, err := h.studentService.GetByID(int(sID))
	if err != nil {
		if err == apierrors.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(apierrors.StatusCode(err), apierrors.Response(err))
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *studentHandler) add(c *gin.Context) {
	var student student.Student
	if err := c.BindJSON(&student); err != nil {
		logrus.Errorf("error unmarshalling the student: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.studentService.Create(&student); err != nil {
		if err == apierrors.ErrEmailAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(apierrors.StatusCode(err), apierrors.Response(err))
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *studentHandler) update(c *gin.Context) {
	studentID := c.Param("id")

	var student student.Student
	if err := c.BindJSON(&student); err != nil {
		logrus.Errorf("error unmarshalling the student: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	sID, err := strconv.ParseUint(studentID, 10, 32)
	if err != nil {
		logrus.Errorf("error parsing the student id %s : %v", studentID, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": apierrors.ErrInvalidID.Error(),
		})
		return
	}
	student.ID = uint(sID)

	if err := h.studentService.Update(&student); err != nil {
		switch err {
		case apierrors.ErrEmailAlreadyExists:
			c.JSON(http.StatusConflict, gin.H{
				"message": err.Error(),
			})
			return
		case apierrors.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(apierrors.StatusCode(err), apierrors.Response(err))
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *studentHandler) delete(c *gin.Context) {
	studentID := c.Param("id")

	sID, err := strconv.ParseInt(studentID, 10, 32)
	if err != nil {
		logrus.Errorf("error parsing the student id %s : %v", studentID, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": apierrors.ErrInvalidID.Error(),
		})
		return
	}

	if err := h.studentService.Delete(int(sID)); err != nil {
		if err == apierrors.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(apierrors.StatusCode(err), apierrors.Response(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "student deleted successfully",
	})
}
