package crudproduct

import (
	"lit_project/inserts"
	"database/sql"
	"encoding/json"
	"lit_project/model"
	"net/http"
)

type CrudProductRepo struct {
	db *sql.DB
}

func NewCrudProductRepo(db *sql.DB) *CrudProductRepo {
	return &CrudProductRepo{db: db}
}

func (i *CrudProductRepo) InsertProduct(w http.ResponseWriter, r *http.Request) {
	product, err := Insert_test.InsertProduct(w, r)
	if err != nil {
		panic(err)
	}

	_, err = i.db.Exec("insert into products(id,name,description,price,stock_quantity) values($1,$2,$3,$4)",
		&product.Id, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		panic(err)
	}
}

func (u *CrudProductRepo) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	up, err := Insert_test.UpdateProduct(w, r)
	if err != nil {
		panic(err)
	}

	_, err = u.db.Exec("update products set name=$1, description=$2,price=$3,stock_quantity=&4  where id=$5",
		&up.Name, &up.Price, &up.Quantity, &up.Id)

	if err != nil {
		panic(err)
	}
}

func (d *CrudProductRepo) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	del, err := Insert_test.DeleteProduct(w, r)
	if err != nil {
		panic(err)
	}

	_, err = d.db.Exec("delete from products where id=$1", &del.Id)
	if err != nil {
		panic(err)
	}
}

func (s *CrudProductRepo) ReadProduct(w http.ResponseWriter, r *http.Request) {
	row, err := s.db.Query("select * from products")
	if err != nil {
		http.Error(w, "Unable to scan the row", http.StatusInternalServerError)
		return
	}
	defer row.Close()

	var products []model.Products
	for row.Next() {
		var product model.Products
		err = row.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			http.Error(w, "Unable to scan the row", http.StatusInternalServerError)
			return
		}

		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		panic(err)
	}
}
