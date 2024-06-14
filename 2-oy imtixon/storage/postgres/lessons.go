package postgres

import (
	"database/sql"
	"fmt"
	"project/models"
	"time"
)

type Lessons struct {
	db *sql.DB
}

func NewLesson(db *sql.DB) *Lessons {
	return &Lessons{db}
}

func (l *Lessons) CreateLesson(lesson *models.Lessons) error {
	_, err := l.db.Exec("INSERT INTO lessons(lesson_id, course_id, title, content, created_at, update_at) VALUES ($1, $2, $3, $4, $5, $6)",
		lesson.ID, // Include lesson.ID here
		lesson.CourseID,
		lesson.Title,
		lesson.Content,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (l *Lessons) ReadLesson(id string) (*models.Lessons, error) {
	row := l.db.QueryRow("select * from lessons where lesson_id = $1", id)
	var lesson models.Lessons
	err := row.Scan(
		&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdateAt, &lesson.DeleteAt)
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (l *Lessons) UpdateLesson(lesson *models.Lessons, id string) error {

	_, err := l.db.Exec("update lessons set title = $1, content = $2 where lesson_id = $3",
		lesson.Title, lesson.Content, id)
	if err != nil {
		fmt.Println("dk------mf", err)
		return err
	}
	return nil
}
func (l *Lessons) DeleteLesson(id string) error {
	_, err := l.db.Exec("delete from lessons where lesson_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Lessons) ReadAllFilterLesson(f models.LessonF) ([]models.LessonF, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select lesson_id,course_id,title,content
	from lessons `
	filter := `where true `

	if len(f.ID) > 0 {
		params["lesson_id"] = f.ID
		filter += " and lesson_id = :lesson_id "
	}
	if len(f.ID) > 0 {
		params["course_id"] = f.ID
		filter += " and course_id = :course_id "
	}
	if len(f.Title) > 0 {
		params["title"] = f.ID
		filter += "and title = :title "
	}

	if len(f.Content) > 0 {
		params["content"] = f.ID
		filter += " and content = :content "
	}

	query = query + filter

	query, arr = ReplaceQueryParam(query, params)
	rows, err := c.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []models.LessonF
	for rows.Next() {
		var lesson models.LessonF
		err = rows.Scan(&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content)
		if err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

// 2-TASK
func (c *Lessons) GetLessonByCourseId(courseId string) (string, []models.Lessons, error) {
	rows, err := c.db.Query("select course_id,title,content from lessons where course_id = $1", courseId)
	if err != nil {
		return "", nil, err
	}
	defer rows.Close()

	var lessons []models.Lessons
	for rows.Next() {
		lesson := models.Lessons{}
		err := rows.Scan(&lesson.CourseID, &lesson.Title, &lesson.Content)
		if err != nil {
			return "", nil, err
		}
		lessons = append(lessons, lesson)
	}
	return courseId, lessons, nil
}
