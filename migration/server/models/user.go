package models

import "time"

type User struct {
	Id        string
	Name      string
	Age       int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Filter struct {
	Name          string
	Age           int
	Email         string
	Limit, Offset int
}
