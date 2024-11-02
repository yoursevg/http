// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	Id     *uint  `json:"id,omitempty"`
	IsDone *bool  `json:"is_done,omitempty"`
	Text   string `json:"text"`
}

// DeleteTasksParams defines parameters for DeleteTasks.
type DeleteTasksParams struct {
	// Id ID of the task to delete
	Id string `form:"id" json:"id"`
}

// PutTasksJSONBody defines parameters for PutTasks.
type PutTasksJSONBody struct {
	// IsDone Updated is_done flag
	IsDone *bool `json:"is_done,omitempty"`

	// Text The updated text of the message
	Text *string `json:"text,omitempty"`
}

// PutTasksParams defines parameters for PutTasks.
type PutTasksParams struct {
	// Id ID of the task to update
	Id string `form:"id" json:"id"`
}

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody = Task

// PutTasksJSONRequestBody defines body for PutTasks for application/json ContentType.
type PutTasksJSONRequestBody PutTasksJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete existing task by ID
	// (DELETE /tasks)
	DeleteTasks(ctx echo.Context, params DeleteTasksParams) error
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx echo.Context) error
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx echo.Context) error
	// Update task ID
	// (PUT /tasks)
	PutTasks(ctx echo.Context, params PutTasksParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteTasks converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTasks(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteTasksParams
	// ------------- Required query parameter "id" -------------

	err = runtime.BindQueryParameter("form", true, true, "id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTasks(ctx, params)
	return err
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx)
	return err
}

// PostTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PostTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTasks(ctx)
	return err
}

// PutTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PutTasks(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PutTasksParams
	// ------------- Required query parameter "id" -------------

	err = runtime.BindQueryParameter("form", true, true, "id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutTasks(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/tasks", wrapper.DeleteTasks)
	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.POST(baseURL+"/tasks", wrapper.PostTasks)
	router.PUT(baseURL+"/tasks", wrapper.PutTasks)

}

type DeleteTasksRequestObject struct {
	Params DeleteTasksParams
}

type DeleteTasksResponseObject interface {
	VisitDeleteTasksResponse(w http.ResponseWriter) error
}

type DeleteTasks204Response struct {
}

func (response DeleteTasks204Response) VisitDeleteTasksResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteTasks404Response struct {
}

func (response DeleteTasks404Response) VisitDeleteTasksResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetTasksRequestObject struct {
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type GetTasks200JSONResponse []Task

func (response GetTasks200JSONResponse) VisitGetTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasksRequestObject struct {
	Body *PostTasksJSONRequestBody
}

type PostTasksResponseObject interface {
	VisitPostTasksResponse(w http.ResponseWriter) error
}

type PostTasks201JSONResponse Task

func (response PostTasks201JSONResponse) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PutTasksRequestObject struct {
	Params PutTasksParams
	Body   *PutTasksJSONRequestBody
}

type PutTasksResponseObject interface {
	VisitPutTasksResponse(w http.ResponseWriter) error
}

type PutTasks200JSONResponse Task

func (response PutTasks200JSONResponse) VisitPutTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutTasks404Response struct {
}

func (response PutTasks404Response) VisitPutTasksResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Delete existing task by ID
	// (DELETE /tasks)
	DeleteTasks(ctx context.Context, request DeleteTasksRequestObject) (DeleteTasksResponseObject, error)
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error)
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error)
	// Update task ID
	// (PUT /tasks)
	PutTasks(ctx context.Context, request PutTasksRequestObject) (PutTasksResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// DeleteTasks operation middleware
func (sh *strictHandler) DeleteTasks(ctx echo.Context, params DeleteTasksParams) error {
	var request DeleteTasksRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTasks(ctx.Request().Context(), request.(DeleteTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTasksResponseObject); ok {
		return validResponse.VisitDeleteTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTasks operation middleware
func (sh *strictHandler) GetTasks(ctx echo.Context) error {
	var request GetTasksRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasks(ctx.Request().Context(), request.(GetTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksResponseObject); ok {
		return validResponse.VisitGetTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTasks operation middleware
func (sh *strictHandler) PostTasks(ctx echo.Context) error {
	var request PostTasksRequestObject

	var body PostTasksJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTasks(ctx.Request().Context(), request.(PostTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTasksResponseObject); ok {
		return validResponse.VisitPostTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PutTasks operation middleware
func (sh *strictHandler) PutTasks(ctx echo.Context, params PutTasksParams) error {
	var request PutTasksRequestObject

	request.Params = params

	var body PutTasksJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutTasks(ctx.Request().Context(), request.(PutTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutTasksResponseObject); ok {
		return validResponse.VisitPutTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}