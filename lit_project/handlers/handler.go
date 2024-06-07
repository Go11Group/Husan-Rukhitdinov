package handler

import (
	crudproduct "lit_project/crudProducts"
	crudusers "lit_project/crudUsers"
	"net/http"
)

type HandlerUser struct {
	User    *crudusers.CrudUsersRepo
	Product *crudproduct.CrudProductRepo
}

type HandlerProduct struct {
	Product *crudproduct.CrudProductRepo
}

func NewHandler(handler HandlerUser) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/insertuser", handler.User.InsertUsers)
	mux.HandleFunc("/updateuser", handler.User.UpdateUsers)
	mux.HandleFunc("/deleteuser", handler.User.DeleteUsers)
	mux.HandleFunc("/readuser", handler.User.ReadUsers)
	mux.HandleFunc("/insertproduct", handler.Product.InsertProduct)
	mux.HandleFunc("/updateproduct", handler.Product.UpdateProduct)
	mux.HandleFunc("/dalateproduct", handler.Product.DeleteProduct)
	mux.HandleFunc("/readproduct", handler.Product.ReadProduct)

	return &http.Server{Handler: mux}
}
