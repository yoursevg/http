// Package messages provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package messages

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Message defines model for Message.
type Message struct {
	Id   *uint   `json:"id,omitempty"`
	Text *string `json:"text,omitempty"`
}

// DeleteMessagesParams defines parameters for DeleteMessages.
type DeleteMessagesParams struct {
	// Id ID of the message to delete
	Id string `form:"id" json:"id"`
}

// PutMessagesJSONBody defines parameters for PutMessages.
type PutMessagesJSONBody struct {
	// Text The updated text of the message
	Text *string `json:"text,omitempty"`
}

// PutMessagesParams defines parameters for PutMessages.
type PutMessagesParams struct {
	// Id ID of the message to update
	Id string `form:"id" json:"id"`
}

// PostMessagesJSONRequestBody defines body for PostMessages for application/json ContentType.
type PostMessagesJSONRequestBody = Message

// PutMessagesJSONRequestBody defines body for PutMessages for application/json ContentType.
type PutMessagesJSONRequestBody PutMessagesJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete existing message by ID
	// (DELETE /messages)
	DeleteMessages(ctx echo.Context, params DeleteMessagesParams) error
	// Get all messages
	// (GET /messages)
	GetMessages(ctx echo.Context) error
	// Create a new message
	// (POST /messages)
	PostMessages(ctx echo.Context) error
	// Update message text by ID
	// (PUT /messages)
	PutMessages(ctx echo.Context, params PutMessagesParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteMessages converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteMessages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteMessagesParams
	// ------------- Required query parameter "id" -------------

	err = runtime.BindQueryParameter("form", true, true, "id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteMessages(ctx, params)
	return err
}

// GetMessages converts echo context to params.
func (w *ServerInterfaceWrapper) GetMessages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMessages(ctx)
	return err
}

// PostMessages converts echo context to params.
func (w *ServerInterfaceWrapper) PostMessages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostMessages(ctx)
	return err
}

// PutMessages converts echo context to params.
func (w *ServerInterfaceWrapper) PutMessages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PutMessagesParams
	// ------------- Required query parameter "id" -------------

	err = runtime.BindQueryParameter("form", true, true, "id", ctx.QueryParams(), &params.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutMessages(ctx, params)
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

	router.DELETE(baseURL+"/messages", wrapper.DeleteMessages)
	router.GET(baseURL+"/messages", wrapper.GetMessages)
	router.POST(baseURL+"/messages", wrapper.PostMessages)
	router.PUT(baseURL+"/messages", wrapper.PutMessages)

}

type DeleteMessagesRequestObject struct {
	Params DeleteMessagesParams
}

type DeleteMessagesResponseObject interface {
	VisitDeleteMessagesResponse(w http.ResponseWriter) error
}

type DeleteMessages204Response struct {
}

func (response DeleteMessages204Response) VisitDeleteMessagesResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteMessages404Response struct {
}

func (response DeleteMessages404Response) VisitDeleteMessagesResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetMessagesRequestObject struct {
}

type GetMessagesResponseObject interface {
	VisitGetMessagesResponse(w http.ResponseWriter) error
}

type GetMessages200JSONResponse []Message

func (response GetMessages200JSONResponse) VisitGetMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostMessagesRequestObject struct {
	Body *PostMessagesJSONRequestBody
}

type PostMessagesResponseObject interface {
	VisitPostMessagesResponse(w http.ResponseWriter) error
}

type PostMessages201JSONResponse Message

func (response PostMessages201JSONResponse) VisitPostMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PutMessagesRequestObject struct {
	Params PutMessagesParams
	Body   *PutMessagesJSONRequestBody
}

type PutMessagesResponseObject interface {
	VisitPutMessagesResponse(w http.ResponseWriter) error
}

type PutMessages200JSONResponse Message

func (response PutMessages200JSONResponse) VisitPutMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutMessages404Response struct {
}

func (response PutMessages404Response) VisitPutMessagesResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Delete existing message by ID
	// (DELETE /messages)
	DeleteMessages(ctx context.Context, request DeleteMessagesRequestObject) (DeleteMessagesResponseObject, error)
	// Get all messages
	// (GET /messages)
	GetMessages(ctx context.Context, request GetMessagesRequestObject) (GetMessagesResponseObject, error)
	// Create a new message
	// (POST /messages)
	PostMessages(ctx context.Context, request PostMessagesRequestObject) (PostMessagesResponseObject, error)
	// Update message text by ID
	// (PUT /messages)
	PutMessages(ctx context.Context, request PutMessagesRequestObject) (PutMessagesResponseObject, error)
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

// DeleteMessages operation middleware
func (sh *strictHandler) DeleteMessages(ctx echo.Context, params DeleteMessagesParams) error {
	var request DeleteMessagesRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteMessages(ctx.Request().Context(), request.(DeleteMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteMessagesResponseObject); ok {
		return validResponse.VisitDeleteMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetMessages operation middleware
func (sh *strictHandler) GetMessages(ctx echo.Context) error {
	var request GetMessagesRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetMessages(ctx.Request().Context(), request.(GetMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetMessagesResponseObject); ok {
		return validResponse.VisitGetMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostMessages operation middleware
func (sh *strictHandler) PostMessages(ctx echo.Context) error {
	var request PostMessagesRequestObject

	var body PostMessagesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostMessages(ctx.Request().Context(), request.(PostMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostMessagesResponseObject); ok {
		return validResponse.VisitPostMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PutMessages operation middleware
func (sh *strictHandler) PutMessages(ctx echo.Context, params PutMessagesParams) error {
	var request PutMessagesRequestObject

	request.Params = params

	var body PutMessagesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutMessages(ctx.Request().Context(), request.(PutMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutMessagesResponseObject); ok {
		return validResponse.VisitPutMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
