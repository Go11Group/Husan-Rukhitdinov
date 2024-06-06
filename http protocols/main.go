package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greeting", greeting)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
func greeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Thist method isn't GET"))
		return
	}
	fmt.Println(r.URL)
	fmt.Println(r.Host)
	fmt.Println(r.Method)
	n, err := w.Write([]byte("hey dear what's up"))
	if err != nil {
		fmt.Println(err,n)
	}
}


