package handler

import (
	"net/http"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kubun/pp/v3"
)

func (h *Handler) CreateLessons(c *gin.Context) {
	var lesson models.Lessons

	// Bind the JSON request body to the lesson struct
	err := c.BindJSON(&lesson)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Check if course_id is empty
	if lesson.CourseID == "" {
		c.String(http.StatusBadRequest, "Course ID is required")
		return
	}

	// Generate a new UUID for lesson_id (after binding)
	lesson.ID = uuid.New().String()
	pp.Print(&lesson)

	err = h.Lesson.CreateLesson(&lesson)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "SUCCESS")
}

func (h *Handler) ReadLessonById(c *gin.Context) {
	id := c.Param("lesson_id")
	if id == "" {
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik bor")
	}

	lesson, err := h.Lesson.ReadLesson(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, lesson)

}
func (h *Handler) UpdateLessons(c *gin.Context) {
	var lesson models.Lessons
	err := c.Bind(&lesson)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.Lesson.UpdateLesson(&lesson, c.Param("lesson_id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(200, "SUCCESS")

}

func (h *Handler) DeleteLessons(c *gin.Context) {
	id := c.Param("lesson_id")
	if id == "" {
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik bor")
	}

	err := h.Lesson.DeleteLesson(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(200, "SUCCESS")

}

func (h *Handler) GetAllFilterLessons(c *gin.Context) {
	lessonFilter := models.LessonF{}
	lessonFilter.ID = c.Param("lesson_id")
	lessonFilter.CourseID = c.Param("course_id")
	lessonFilter.Title = c.Param("title")
	lessonFilter.Content = c.Param("description")
	les, err := h.Lesson.ReadAllFilterLesson(lessonFilter)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

	}
	c.JSON(http.StatusOK, les)

}

// 2-TASK
func (h *Handler) GetLessonsByCoursesId(c *gin.Context) {
	courseId, lesssons, err := h.Lesson.GetLessonByCourseId(c.Param("lesson_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(lesssons); i++ {
			c.JSON(200, gin.H{
				"message":   "Success",
				"status":    http.StatusOK,
				"time":      time.Now(),
				"course_id": courseId,
				"Lessons": gin.H{
					"lesson_id": lesssons[i].ID,
					"course_id": lesssons[i].CourseID,
					"title":     lesssons[i].Title,
					"content":   lesssons[i].Content,
				},
			})
		}
	}
}
