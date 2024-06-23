package handler

import (
	"database/sql"
	"project/storage/postgres"

	"github.com/gin-gonic/gin"
)

func NewGin(db *sql.DB, user *postgres.User, course *postgres.Courses, lesson *postgres.Lessons, enrollment *postgres.Enrollments) *gin.Engine {
	k := gin.Default()

	handler := NewHandler(db, user, course, lesson, enrollment)

	// Users
	k.GET("getAllUsers", handler.GetFilterAllUsers)
	k.GET("getUser/:user_id", handler.GetUserById)
	k.GET("getUserCourse/:user_id", handler.GetCourseByUsers)
	k.GET("getSearchUser", handler.SearchUsers)
	k.PUT("updateUser/:user_id", handler.UpdateUser)
	k.POST("createUser", handler.CreateUser)
	k.DELETE("deleteUser/:user_id", handler.DeleteUsers)

	// COURSES
	k.GET("getCourse/:course_id", handler.ReadCourseById)
	k.GET("getFilterAllCourse", handler.GetFilterAllCourses)
	k.GET("getEnrollUsersByCourse", handler.GetEnrollUsersByCourses)
	k.GET("getPopularCourse", handler.GetPopularCourses)
	k.PUT("UpdateCourse/:course_id", handler.UpdateCourse)
	k.POST("createCourse", handler.CreateCourse)
	k.DELETE("deleteCourse/:course_id", handler.DeleteCourse)
	// LESSON
	k.GET("getLesson/:lesson_id", handler.ReadLessonById)
	k.GET("getAllFilterLesson", handler.GetAllFilterLessons)
	k.GET("getLessonByCourseID", handler.GetLessonsByCoursesId)
	k.PUT("updateLesson/:lesson_id", handler.UpdateLessons)
	k.POST("createLesson", handler.CreateLessons)
	k.DELETE("deleteLesson/:lesson_id", handler.DeleteLessons)
	// ENROLLMENT
	k.GET("getEnrollment/:enrollment_id", handler.ReadEnrollmentById)
	k.GET("getFilterAllEnrollment", handler.GetFilterAllEnrollments)
	k.PUT("updateEnrollment/:enrollment_id", handler.UpdateEnrollments)
	k.POST("createEnrollment", handler.CreateEnrollments)
	k.DELETE("deleteEnrollment", handler.DeleteEnrollments)

	return k
}
