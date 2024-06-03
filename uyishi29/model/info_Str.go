package model

type Users struct {
	Id       int
	Username string
	Email    string
	Password string
}

type Product struct {
	Id             int
	Name           string
	Discription    string
	Price          float32
	StockQuantity int
}

type UserProducts struct {
    Id        int
    UserId    int
    ProductId int
}
