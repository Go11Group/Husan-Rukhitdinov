package postgres

import (
	"database/sql"
	"fmt"
	"project/models"
	"time"

	"github.com/google/uuid"
)

type Enrollments struct {
	db *sql.DB
}

func NewEnrollments(db *sql.DB) *Enrollments {
	return &Enrollments{db}
}
func (e *Enrollments) CreateEnrollment(enroll *models.Enrollments) error {
	enroll.ID = uuid.NewString()

	// Parse the EnrollmentDate string to a time.Time object
	enrollmentDate, err := time.Parse("2006-01-02", enroll.EnrollmentDate) // Use the correct format
	if err != nil {
		return fmt.Errorf("invalid enrollment date format: %w", err)
	}

	enroll.DeleteAt = 0
	_, err = e.db.Exec("insert into enrollments (enrollment_id,user_id,course_id,enrollment_date) VALUES ($1,$2,$3,$4)",
		enroll.ID, enroll.UserID, enroll.CourseID, enrollmentDate)
	if err != nil {
		return err
	}
	return nil
}
func (e *Enrollments) ReadEnrollment(id string) (*models.Enrollments, error) {
	row := e.db.QueryRow("select * from enrollments where enrollment_id = $1", id)
	var enroll models.Enrollments
	err := row.Scan(
		&enroll.ID, &enroll.UserID, &enroll.CourseID, &enroll.EnrollmentDate, &enroll.CreatedAt, &enroll.UpdateAt, &enroll.DeleteAt)
	if err != nil {
		return nil, err
	}
	return &enroll, nil
}

func (e *Enrollments) UpdateEnrollment(enroll *models.Enrollments, id string) error {

	_, err := e.db.Exec("update enrollment set user_id = $1, course_id = $2, enrollment_date = $3 where enrollment_id = $4",
		enroll.UserID, enroll.CourseID, enroll.EnrollmentDate, id)
	if err != nil {
		return err
	}
	return nil
}
func (e *Enrollments) DeleteEnrollment(id string) error {
	_, err := e.db.Exec("delete from enrollments where enrollment_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Enrollments) FilterAllEnrollment(f models.EnrollmentF) ([]models.EnrollmentF, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	query := `select enrollment_id, user_id, course_id, enrollment_date
	from enrollments `
	filter := `where true `

	if len(f.ID) > 0 {
		params["enrollment_id"] = f.ID
		filter += " and enrollment_id = :enrollment_id "
	}
	if len(f.UserID) > 0 {
		params["user_id"] = f.ID
		filter += " and user_id = :user_id "
	}
	if len(f.CourseID) > 0 {
		params["course_id"] = f.ID
		filter += " and course_id = :course_id "
	}
	if len(f.EnrollmentDate) > 0 {
		params["enrollment_id"] = f.ID
		filter += "and enrollment_id = :enrollment_id "
	}

	query = query + filter

	query, arr = ReplaceQueryParam(query, params)
	rows, err := c.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []models.EnrollmentF
	for rows.Next() {
		var enroll models.EnrollmentF
		err = rows.Scan(&enroll.ID, &enroll.UserID, &enroll.CourseID, &enroll.EnrollmentDate)
		if err != nil {
			return nil, err
		}

		enrollments = append(enrollments, enroll)
	}
	return enrollments, nil
}
