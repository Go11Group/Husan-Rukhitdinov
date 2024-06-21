package handler

import (
	"fmt"
	"net/http"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourse(c *gin.Context) {
	var course models.Courses
	err := c.BindJSON(&course)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.Course.CreateCourse(&course)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusCreated, course)
}

func (h *Handler) ReadCourseById(c *gin.Context) {
	id := c.Param("course_id")
	if id == "" {
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik bor")
	}

	course, err := h.Course.ReadCourse(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, course)

}
func (h *Handler) UpdateCourse(c *gin.Context) {
	var course models.Courses
	err := c.Bind(&course)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.Course.UpdateCourse(&course,c.Param("course_id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(200, "SUCCESS")

}
func (h *Handler) DeleteCourse(c *gin.Context) {
	id := c.Param("course_id")
	fmt.Println(id)
	if id == "" {
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik bor")
	}

	err := h.Course.DeleteCourse(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(200, "SUCCESS")

}

func (h *Handler) GetFilterAllCourses(c *gin.Context) {
	courseFilter := models.CourseF{}
	courseFilter.ID = c.Param("course_id")
	courseFilter.Title = c.Param("title")
	courseFilter.Description = c.Param("description")
	us,err := h.Course.ReadFilterAllCourse(courseFilter)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, us)

}

// 3-TASK
func (h *Handler) GetEnrollUsersByCourses(c *gin.Context) {
	courseId, users, err := h.Course.GetEnrollUserByCourse(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {

		for i := 0; i < len(users); i++ {
			c.JSON(http.StatusOK, gin.H{
				"message":   "Success",
				"status":    http.StatusOK,
				"time":      time.Now(),
				"course_id": courseId,
				"User": gin.H{
					"message": "Success",
					"status":  http.StatusOK,
					"user_id": users[i].ID,
					"name":    users[i].Name,
					"email":   users[i].Email,
				},
			})
		}
	}
}

// 5-TASK
func (h *Handler) GetPopularCourses(c *gin.Context) {
	startTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2024, 12, 31, 23, 59, 59, 999, time.UTC)

	time1 := c.Query("start_time")
	time2 := c.Query("end_time")
	if time1 != "" {
		parsedStartTime, err := time.Parse("2006-01-02", time1)
		if err == nil {
			startTime = parsedStartTime
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid start_time format. Use YYYY-MM-DD.",
				"status":  http.StatusBadRequest,
			})
			return
		}
	}
	if time2 != "" {
		parsedEndTime, err := time.Parse("2006-01-02", time2)
		if err == nil {
			endTime = parsedEndTime
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid end_time format. Use YYYY-MM-DD.",
				"status":  http.StatusBadRequest,
			})
			return
		}
	}

	count, course, err := h.Course.GetPopularyCourse(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"status":  http.StatusOK,
		"time":    time.Now(),
		"id":      course.ID,
		"title":   course.Title,
		"count":   count,
	})
}
