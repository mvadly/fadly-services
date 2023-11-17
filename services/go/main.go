package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Golang!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	fmt.Println("Server is listening on :9999...")
	http.ListenAndServe(":9999", r)
}
