package main

import (
	"github.com/gorilla/mux"
	"mux/internal/database"
	"mux/internal/handlers"
	"mux/internal/messagesService"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/update", handler.UpdateMessageHandler).Methods("PUT")
	router.HandleFunc("/api/delete", handler.DeleteMessageHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
