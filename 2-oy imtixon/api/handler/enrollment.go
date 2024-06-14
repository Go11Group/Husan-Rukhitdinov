package handler

import (
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kubun/pp/v3"
)

func (h *Handler) CreateEnrollments(c *gin.Context) {
	var enroll models.Enrollments
	pp.Print(&enroll)

	// Bind JSON data
	err := c.BindJSON(&enroll)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if enroll.CourseID == "" {
		c.String(http.StatusBadRequest, "Course ID is required")
		return
	}
	if enroll.UserID == "" {
		c.String(http.StatusBadRequest, "User ID is required")
		return
	}

	enroll.ID = uuid.New().String()
	pp.Print(&enroll)

	// Create the enrollment
	err = h.Enrollment.CreateEnrollment(&enroll)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Return the created enrollment
	c.JSON(http.StatusOK, "SECCESS")
}
func (h *Handler) ReadEnrollmentById(c *gin.Context) {
	id := c.Param("enrollment_id")
	if id == "" {
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}

	enroll, err := h.Enrollment.ReadEnrollment(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, enroll)

}

func (h *Handler) UpdateEnrollments(c *gin.Context) {
	var enroll models.Enrollments
	err := c.Bind(&enroll)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.Enrollment.UpdateEnrollment(&enroll, c.Param("enrollment_id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(200, "SUCCESS")

}

func (h *Handler) DeleteEnrollments(c *gin.Context) {
	id := c.Param("enrollment_id")
	if id == "" {
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}

	err := h.Enrollment.DeleteEnrollment(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(200, "SUCCESS")

}

func (h *Handler) GetFilterAllEnrollments(c *gin.Context) {
	enrollmentFilter := models.EnrollmentF{}
	enrollmentFilter.ID = c.Param("enrollment_id")
	enrollmentFilter.UserID = c.Param("user_id")
	enrollmentFilter.CourseID = c.Param("course_id")
	enrollmentFilter.EnrollmentDate = c.Param("enrollment_date")
	en, err := h.Enrollment.FilterAllEnrollment(enrollmentFilter)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

	}
	c.JSON(http.StatusOK, en)
}
