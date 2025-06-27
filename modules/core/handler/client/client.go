package client

import "net/http"

type HttpContext struct {
	Request  HttpRequest
	Response HttpResponse
	Handler  HttpHandler
}

type HttpHandler struct {
	Req *http.Request
	Res http.ResponseWriter
}

type HttpRequest struct {
	Method string
	Url    string
	Header http.Header
	Params map[string][]string
	Body   []byte
}

type HttpResponse struct {
	Status  int32  `json:"status"`
	Message string `json:"message,omitempty"`
	Data    []any  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
