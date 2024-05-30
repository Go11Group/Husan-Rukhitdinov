package postgres

import (
	"database/sql"
	"go_modul/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) CreateCourse(course *model.Course) error {
	_,err := c.DB.Exec(`insert into coourse(name,field)values($1,$2)`,&course.Name,&course.Field)
	if err != nil{
		return err
	}
	return nil
}
func (c *CourseRepo) UpdateCourse(course *model.Course) error {
	_, err := c.DB.Exec(`
	UPDATE course
	SET name = $1
	WHERE id = $2`, &course.Name,&course.Id)
	if err != nil{
		return err
	}
	return nil
}

func (c *CourseRepo) DeleteCourse(id string) error {
	_, err := c.DB.Exec(`
	delete from course
	WHERE id = $1`,id)
	if err != nil{
		return err
	}
	return nil
}
