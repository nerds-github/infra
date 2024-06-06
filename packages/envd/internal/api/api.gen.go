//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.1-0.20240519200907-da9077bb5ffe DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// FilePath defines model for FilePath.
type FilePath = string

// User defines model for User.
type User = string

// BadRequest defines model for BadRequest.
type BadRequest = Error

// DirectoryPathError defines model for DirectoryPathError.
type DirectoryPathError = Error

// FileNotFound defines model for FileNotFound.
type FileNotFound = Error

// InternalServerError defines model for InternalServerError.
type InternalServerError = Error

// NotEnoughDiskSpace defines model for NotEnoughDiskSpace.
type NotEnoughDiskSpace = Error

// GetFilesPathParams defines parameters for GetFilesPath.
type GetFilesPathParams struct {
	// Username User used for setting the owner, or resolving relative paths.
	Username User `form:"username" json:"username"`
}

// PutFilesPathMultipartBody defines parameters for PutFilesPath.
type PutFilesPathMultipartBody struct {
	File *openapi_types.File `json:"file,omitempty"`
}

// PutFilesPathParams defines parameters for PutFilesPath.
type PutFilesPathParams struct {
	// Username User used for setting the owner, or resolving relative paths.
	Username User `form:"username" json:"username"`
}

// PutFilesPathMultipartRequestBody defines body for PutFilesPath for multipart/form-data ContentType.
type PutFilesPathMultipartRequestBody PutFilesPathMultipartBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Download a file
	// (GET /files/{path})
	GetFilesPath(w http.ResponseWriter, r *http.Request, path FilePath, params GetFilesPathParams)
	// Upload a file and ensure the parent directories exist. If the file exists, it will be overwritten.
	// (PUT /files/{path})
	PutFilesPath(w http.ResponseWriter, r *http.Request, path FilePath, params PutFilesPathParams)
	// Ensure the time and metadata is synced with the host
	// (POST /sync)
	PostSync(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetFilesPath operation middleware
func (siw *ServerInterfaceWrapper) GetFilesPath(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "path" -------------
	var path FilePath

	err = runtime.BindStyledParameterWithOptions("simple", "path", r.PathValue("path"), &path, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "path", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFilesPathParams

	// ------------- Required query parameter "username" -------------

	if paramValue := r.URL.Query().Get("username"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "username"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "username", r.URL.Query(), &params.Username)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "username", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFilesPath(w, r, path, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutFilesPath operation middleware
func (siw *ServerInterfaceWrapper) PutFilesPath(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "path" -------------
	var path FilePath

	err = runtime.BindStyledParameterWithOptions("simple", "path", r.PathValue("path"), &path, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "path", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PutFilesPathParams

	// ------------- Required query parameter "username" -------------

	if paramValue := r.URL.Query().Get("username"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "username"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "username", r.URL.Query(), &params.Username)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "username", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutFilesPath(w, r, path, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostSync operation middleware
func (siw *ServerInterfaceWrapper) PostSync(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostSync(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       *http.ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m *http.ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/files/{path}", wrapper.GetFilesPath)
	m.HandleFunc("PUT "+options.BaseURL+"/files/{path}", wrapper.PutFilesPath)
	m.HandleFunc("POST "+options.BaseURL+"/sync", wrapper.PostSync)

	return m
}
