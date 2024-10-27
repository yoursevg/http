package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MessageStruct struct {
	Text string `json:"text"`
}

func GetHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(DB.Find(&MessageStruct{}))
}

func PostHelloHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var messageStruct MessageStruct
	err := decoder.Decode(&messageStruct)
	if err != nil {
		panic(err)
	}

	message := &Message{
		Text: messageStruct.Text,
	}

	//Добавляем запись в БД
	result := DB.Create(&message)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Fprintln(w, "New message: ", message)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHelloHandler).Methods("GET")
	router.HandleFunc("/api/hello", PostHelloHandler).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
