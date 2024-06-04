package main

import (
	"database/sql"
	"fmt"
	"github.com/go-faker/faker/v3"
	_ "github.com/go-faker/faker/v3"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"math/rand"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:3333@localhost:5432/connectdb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 1000000; i++ {
		_, err = db.Exec("insert into product (id, name, category, cost) values ($1, $2, $3, $4)",
			uuid.NewString(), faker.FirstName(), faker.LastName(), rand.Intn(100000))
		if err != nil {
			fmt.Println(err)
		}
		if i%100000 == 0 {
			fmt.Println(i)
		}
	}
}
