package client

import (
	"encoding/json"
)

func (_http *HttpContext) Json(status int32, message string, data []any, err error) error {
	_http.Response.Status = status
	_http.Response.Message = message
	_http.Response.Data = data
	if err != nil {
		_http.Response.Error = err.Error()
	} else {
		_http.Response.Error = ""
	}

	_http.Handler.Res.Header().Set("Content-Type", "application/json")
	_http.Handler.Res.WriteHeader(int(status))
	return json.NewEncoder(_http.Handler.Res).Encode(_http.Response)
}

func (_http *HttpContext) Success(message string, data []any) error {
	return _http.Json(200, message, data, nil)
}

func (_http *HttpContext) Created(message string, data []any) error {
	return _http.Json(201, message, data, nil)
}

func (_http *HttpContext) Error(message string, err error) error {
	return _http.Json(400, message, nil, err)
}

func (_http *HttpContext) Unauthorized(err error) error {
	return _http.Json(401, "Unauthorized", nil, err)
}

func (_http *HttpContext) Forbidden(err error) error {
	return _http.Json(403, "Forbidden", nil, err)
}

func (_http *HttpContext) NotFound(err error) error {
	return _http.Json(404, "NotFound", nil, err)
}

func (_http *HttpContext) InternalServerError(err error) error {
	return _http.Json(500, "Internal Server Error", nil, err)
}
