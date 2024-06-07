package model

type Users struct {
	Id        int    `json:"id"`
	Firstname string `firstname`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age`
}

type Products struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
