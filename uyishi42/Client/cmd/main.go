package main

import "pro/handlers"

func main() {
	h := handlers.NewHandler()
	server := handlers.CreateServer(h)
	server.ListenAndServe()
}
