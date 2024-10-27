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

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if message == "" {
			fmt.Fprintln(w, "Hello, world!")
		} else {
			fmt.Fprintf(w, "Hello, %s!\n", message)
		}
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var messageJSON MessageStruct
		err := decoder.Decode(&messageJSON)
		if err != nil {
			panic(err)
		}
		message = messageJSON.Message
		fmt.Println(message)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET", "POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
