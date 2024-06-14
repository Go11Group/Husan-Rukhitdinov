package postgres

import (
	"database/sql"
	"fmt"
	"project/models"
	"time"

	"github.com/google/uuid"
)

type Courses struct {
	db *sql.DB
}

func NewCourse(db *sql.DB) *Courses {
	return &Courses{db}
}

func (c *Courses) CreateCourse(course *models.Courses) error {
	course.ID = uuid.NewString()
	_, err := c.db.Exec("insert into courses (course_id,title,description) VALUES ($1,$2,$3)",
		course.ID, course.Title, course.Description)
	if err != nil {
		return err
	}
	return nil
}

func (c *Courses) ReadCourse(id string) (*models.Courses, error) {
	row := c.db.QueryRow("select * from courses where course_id = $1", id)
	var course models.Courses
	err := row.Scan(
		&course.ID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdateAt, &course.DeleteAt)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (c *Courses) UpdateCourse(course *models.Courses, id string) error {

	_, err := c.db.Exec("update courses set title = $1, description = $2 where course_id = $3",
		course.Title, course.Description, id)
	if err != nil {
		return err
	}
	return nil
}
func (c *Courses) DeleteCourse(id string) error {
	_, err := c.db.Exec("delete from courses where course_id = $1", id)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

// 3-TASK
func (c *Courses) GetEnrollUserByCourse(courseId string) (string, []models.Users, error) {
	var userID string
	err := c.db.QueryRow("select  user_id from enrollments where course_id = $1", courseId).Scan(&userID)
	if err != nil {
		return "", nil, err
	}
	rows, err := c.db.Query("select user_id,name,email from  users where id=$1", &userID)
	if err != nil {
		return "", nil, err
	}
	users := []models.Users{}
	for rows.Next() {
		user := models.Users{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return "", nil, err
		}
		users = append(users, user)
	}
	return courseId, users, nil
}

// 5-TASK
func (c *Courses) GetPopularyCourse(startTime, endTime time.Time) (int, *models.Courses, error) {
	course := models.Courses{}
	var enrollmentCount int
	err := c.db.QueryRow(`select c.course_id, c.title, COUNT(*)from courses c inner join  enrollments e ON c.course_id = e.course_id where e.enrollment_date BETWEEN $1 AND $2 group by  c.course_id, c.title order by  count(*) desc limit 1`, startTime, endTime).Scan(&course.ID, &course.Title, &enrollmentCount)
	if err != nil {
		return 0, nil, err
	}
	return enrollmentCount, &course, nil
}

func (c *Courses) ReadFilterAllCourse(f models.CourseF) ([]models.Courses, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select course_id,title,description
	from courses `
	filter := `where true `

	if len(f.ID) > 0 {
		params["course_id"] = f.ID
		filter += " and course_id = :course_id "
	}
	if len(f.Title) > 0 {
		params["title"] = f.ID
		filter += "and title = :title "
	}

	if len(f.Description) > 0 {
		params["description"] = f.ID
		filter += " and description = :description "
	}

	query = query + filter

	query, arr = ReplaceQueryParam(query, params)
	rows, err := c.db.Query(query, arr...)
	if err != nil {
		fmt.Println("fffff....",err)
		return nil, err
	}
	defer rows.Close()

	var courses []models.Courses
	for rows.Next() {
		var course models.Courses
		err = rows.Scan(&course.ID, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}
	return courses, nil
}
