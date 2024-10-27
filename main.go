package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
