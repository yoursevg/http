package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"mux/internal/database"
	"mux/internal/handlers"
	"mux/internal/taskService"
	"mux/internal/userService"
	"mux/internal/web/tasks"
	"mux/internal/web/users"
)

func main() {
	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
