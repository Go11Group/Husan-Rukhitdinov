package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "info_persons"
	password = "3333"
)

type Info struct {
	first_name   string
	second_name  string
	age          int
	gender       string
	phone_number string
}

func main() {

	con := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`
    SELECT
      first_name,
	  second_name,
	  age,
	  gender,
	  phone_number
    FROM
      info_peoples
  `)
	if err != nil {
		log.Println(err)
	}

	infos := []Info{}

	for rows.Next() {
		info := Info{}
		rows.Scan(
			&info.first_name,
			&info.second_name,
			&info.age,
			&info.gender,
			&info.phone_number,
		)
		infos = append(infos, info)
	}

	count := 0
	for _, v := range infos {
		fmt.Println(v)
		count++
	}
	fmt.Println(count)

}
