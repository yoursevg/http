package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MessageStruct struct {
	Message string `json:"message"`
}

var message string

func GetHelloHandler(w http.ResponseWriter, r *http.Request) {
	if message == "" {
		fmt.Fprintln(w, "Hello, world!")
	} else {
		fmt.Fprintf(w, "Hello, %s!\n", message)
	}
}

func PostHelloHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var messageJSON MessageStruct
	err := decoder.Decode(&messageJSON)
	if err != nil {
		panic(err)
	}
	message = messageJSON.Message
	fmt.Fprintln(w, "New message: ", message)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHelloHandler).Methods("GET")
	router.HandleFunc("/api/hello", PostHelloHandler).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
