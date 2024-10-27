package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MessageStruct struct {
	Text string `json:"text"`
}

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	var msgs []Message
	DB.Find(&msgs)

	// Конвертируем данные в JSON и отправляем в ответ
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(msgs); err != nil {
		http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
	}
}

func PostMessageHandler(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintln(w, "New message: ", message.Text)
}

func PatchMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	//Парсим пришедший нам в URL id
	msgId := r.URL.Query().Get("id")
	u64, err := strconv.ParseUint(msgId, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	msgIdUint := uint(u64)
	message.ID = msgIdUint
	//Находим нашу запись в бд
	DB.First(&message)

	//Декодируем, что прислал клиент в JSONе
	var newMessageStruct MessageStruct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newMessageStruct)
	if err != nil {
		panic(err)
	}
	//Текст на замену от клиента
	newMessage := newMessageStruct.Text

	//Сохраняем изменения в БД
	message.Text = newMessage
	DB.Save(&message)

	//Отправляем клиенту в ответ изменненную структуру
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
	}
}

func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	//Парсим пришедший нам в URL id
	msgId := r.URL.Query().Get("id")
	u64, err := strconv.ParseUint(msgId, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	msgIdUint := uint(u64)
	message.ID = msgIdUint
	//Удаляем в БД найденную запись
	DB.Delete(&message)
	//Отправляем клиенту ответ
	fmt.Fprintf(w, "Deleted message with id: %s", msgId)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/message", GetMessageHandler).Methods("GET")
	router.HandleFunc("/api/message", PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/message", PatchMessageHandler).Methods("PATCH")
	router.HandleFunc("/api/message", DeleteMessageHandler).Methods("DELETE")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
