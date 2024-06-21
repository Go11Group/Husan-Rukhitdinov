package main

import (
	"fmt"
	"project/api/handler"
	"project/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connect!!!")

	user := postgres.NewUser(db)
	course := postgres.NewCourse(db)
	lesson := postgres.NewLesson(db)
	enrollment := postgres.NewEnrollments(db)

	l := handler.NewGin(db, user, course, lesson, enrollment)
	l.Run("localhost:8080")
}
