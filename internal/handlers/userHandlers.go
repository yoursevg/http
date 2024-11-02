package handlers

import (
	"context"
	"mux/internal/userService"
	"mux/internal/web/users"
	"strconv"
)

type UserHandler struct {
	Service *userService.UserService
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	// Получение всех сообщений из сервиса
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    usr.Email,
			Password: usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{Email: userRequest.Email, Password: userRequest.Password}
	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    createdUser.Email,
		Password: createdUser.Password,
	}
	return response, nil
}

func (h *UserHandler) PutUsers(_ context.Context, request users.PutUsersRequestObject) (users.PutUsersResponseObject, error) {
	userRequest := request.Body
	//Парсим пришедший нам в URL id
	msgIdU64, err := strconv.ParseUint(request.Params.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	userToUpdate := userService.User{Email: *userRequest.Email, Password: *userRequest.Password}
	//Выполняем обработку
	updatedUser, err := h.Service.UpdateUserByID(uint(msgIdU64), userToUpdate)
	if err != nil {
		return users.PutUsers404Response{}, err
	}
	response := users.PutUsers200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    updatedUser.Email,
		Password: updatedUser.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsers(_ context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	//Парсим пришедший нам в URL id
	msgIdU64, err := strconv.ParseUint(request.Params.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	//Выполняем обработку
	err = h.Service.DeleteUserByID(uint(msgIdU64))
	if err != nil {
		return users.DeleteUsers404Response{}, err
	}
	return users.DeleteUsers204Response{}, err
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
