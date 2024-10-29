package handlers

import (
	"context"
	"mux/internal/messagesService" // Импортируем наш сервис
	"mux/internal/web/messages"
	"strconv"
)

type Handler struct {
	Service *messagesService.MessageService
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}
	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:   &msg.ID,
			Text: &msg.Text,
		}
		response = append(response, message)
	}

	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToCreate := messagesService.Message{Text: *messageRequest.Text}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)
	if err != nil {
		return nil, err
	}
	response := messages.PostMessages201JSONResponse{
		Id:   &createdMessage.ID,
		Text: &createdMessage.Text,
	}
	return response, nil
}

func (h *Handler) PutMessages(ctx context.Context, request messages.PutMessagesRequestObject) (messages.PutMessagesResponseObject, error) {
	//Парсим пришедший нам в URL id
	msgIdU64, err := strconv.ParseUint(request.Params.Id, 10, 32)
	if err != nil {
		return nil, err
	}

	messageToUpdate := messagesService.Message{Text: *request.Body.Text}
	//Выполняем обработку
	updatedMessage, err := h.Service.UpdateMessageByID(int(msgIdU64), messageToUpdate)
	if err != nil {
		return nil, err
	}

	response := messages.PutMessages200JSONResponse{
		Id:   &updatedMessage.ID,
		Text: &updatedMessage.Text,
	}

	return response, nil
}

func (h *Handler) DeleteMessages(_ context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	//Парсим пришедший нам в URL id
	msgIdU64, err := strconv.ParseUint(request.Params.Id, 10, 32)
	if err != nil {
		return nil, err
	}

	//Выполняем обработку
	err = h.Service.DeleteMessageByID(int(msgIdU64))
	if err != nil {
		return nil, err
	}

	return messages.DeleteMessages204Response{}, err
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}
