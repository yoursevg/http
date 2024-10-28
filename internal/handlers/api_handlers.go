package handlers

import (
	"encoding/json"
	"fmt"
	"mux/internal/messagesService" // Импортируем наш сервис
	"net/http"
	"strconv"
)

type Handler struct {
	Service *messagesService.MessageService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

func (h *Handler) UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message

	//Парсим пришедший нам в URL id
	msgId := r.URL.Query().Get("id")
	msgIdU64, err := strconv.ParseUint(msgId, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	msgIdUint := uint(msgIdU64)
	message.ID = msgIdUint

	//Заполняем нашу структуру message тем, что пришло от клиента (message.text)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//Выполняем обработку
	updatedMessage, err := h.Service.UpdateMessageByID(int(msgIdU64), message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//Отдаём ответ
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updatedMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message

	//Парсим пришедший нам в URL id
	msgId := r.URL.Query().Get("id")
	msgIdU64, err := strconv.ParseUint(msgId, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	msgIdUint := uint(msgIdU64)
	message.ID = msgIdUint

	//Выполняем обработку
	err = h.Service.DeleteMessageByID(int(msgIdU64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//Ответ
	fmt.Fprintf(w, "Successfully deleted message with id %s\n", msgId)
}
