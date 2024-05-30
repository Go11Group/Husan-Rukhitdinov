package main

import (
	"fmt"
	"go_modul/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	st := postgres.NewStudentRepo(db)

	id := "441bf2db-0467-4869-985b-fba9adc65f32"
	err = st.DeleteStudent(id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ochirildi")


}
