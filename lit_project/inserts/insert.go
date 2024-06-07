package Insert_test

import (
	"encoding/json"
	"io"
	"lit_project/model"
	"net/http"
)

func InsertUser(w http.ResponseWriter, r *http.Request) (model.Users, error) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Bu so'rov get so'rovi emas?"))
	}

	byt, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Users{}, err
	}

	var user model.Users
	err = json.Unmarshal(byt, &user)
	if err != nil {
		return model.Users{}, err
	}

	return user, nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) (model.Users, error) {
	if r.Method != http.MethodPut {
		w.Write([]byte("Bu so'rov get so'rovi emas?"))
	}

	byt, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Users{}, err
	}

	var userUpdate model.Users
	err = json.Unmarshal(byt, &userUpdate)
	if err != nil {
		return model.Users{}, err
	}

	return userUpdate, nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) (model.Users, error) {
	if r.Method != http.MethodDelete {
		w.Write([]byte("Bu so'rov get so'rovi emas?"))
	}

	byt, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Users{}, err
	}

	var userDelete model.Users
	err = json.Unmarshal(byt, &userDelete)
	if err != nil {
		return model.Users{}, err
	}

	return userDelete, nil
}

func InsertProduct(w http.ResponseWriter, r *http.Request) (model.Products, error) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Bu so'rov get so'rovi emas?"))
	}

	byt, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Products{}, err
	}

	var product model.Products
	err = json.Unmarshal(byt, &product)
	if err != nil {
		return model.Products{}, err
	}

	return product, nil
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) (model.Products, error) {
	if r.Method != http.MethodPut {
		w.Write([]byte("Bu so'rov get so'rovi emas?"))
	}

	byt, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Products{}, err
	}

	var productUpdate model.Products
	err = json.Unmarshal(byt, &productUpdate)
	if err != nil {
		return model.Products{}, err
	}

	return productUpdate, nil
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) (model.Products, error) {
	if r.Method != http.MethodDelete {
		w.Write([]byte("Bu so'rov get so'rovi emas?"))
	}

	byt, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Products{}, err
	}

	var productDelete model.Products
	err = json.Unmarshal(byt, &productDelete)
	if err != nil {
		return model.Products{}, err
	}

	return productDelete, nil
}
