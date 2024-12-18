package handlers

import (
	"context"
	"mux/internal/taskService"
	"mux/internal/web/tasks"
	"strconv"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех сообщений из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Text:   tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{Text: taskRequest.Text, IsDone: *taskRequest.IsDone}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Text:   createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) PutTasks(_ context.Context, request tasks.PutTasksRequestObject) (tasks.PutTasksResponseObject, error) {
	taskRequest := request.Body
	//Парсим пришедший нам в URL id
	msgIdU64, err := strconv.ParseUint(request.Params.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	taskToUpdate := taskService.Task{Text: *taskRequest.Text, IsDone: *taskRequest.IsDone}
	//Выполняем обработку
	updatedTask, err := h.Service.UpdateTaskByID(uint(msgIdU64), taskToUpdate)
	if err != nil {
		return tasks.PutTasks404Response{}, err
	}
	response := tasks.PutTasks200JSONResponse{
		Id:     &updatedTask.ID,
		Text:   updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasks(_ context.Context, request tasks.DeleteTasksRequestObject) (tasks.DeleteTasksResponseObject, error) {
	//Парсим пришедший нам в URL id
	msgIdU64, err := strconv.ParseUint(request.Params.Id, 10, 32)
	if err != nil {
		return nil, err
	}
	//Выполняем обработку
	err = h.Service.DeleteTaskByID(uint(msgIdU64))
	if err != nil {
		return tasks.DeleteTasks404Response{}, err
	}
	return tasks.DeleteTasks204Response{}, err
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}
