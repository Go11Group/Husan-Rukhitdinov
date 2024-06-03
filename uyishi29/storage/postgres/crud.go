package postgres

import (
	"database/sql"
	"go_run/model"
)

type UserRepo struct {
	DB *sql.DB
}

func NewInfoRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateNewUser(user model.Users) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err = tx.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) CreateNewProduct(userID int, product model.Product) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var productID int
	query := `INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(query, product.Name, product.Discription, product.Price, product.StockQuantity).Scan(&productID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO user_products (user_id, product_id) VALUES ($1, $2)`, userID, productID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) ReadAllUsers() ([]model.Users, error) {
	rows, err := r.DB.Query(`SELECT id, username, email, password FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		var user model.Users
		err = rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) ReadAllProducts() ([]model.Product, error) {
	rows, err := r.DB.Query(`SELECT id, name, description, price, stock_quantity FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Discription, &product.Price, &product.StockQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *UserRepo) UpdateUser(user model.Users) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UpdateProduct(product model.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5`
	_, err := r.DB.Exec(query, product.Name, product.Discription, product.Price, product.StockQuantity, product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteUser(userId int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteProduct(productId int) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM user_products WHERE product_id = $1`, productId)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM products WHERE id = $1`, productId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
