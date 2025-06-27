package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AppRouter struct {
	Router     *mux.Router
	middleware []func(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request)
}

func Router() *AppRouter {
	r := mux.NewRouter()
	return &AppRouter{Router: r}
}

func (ar *AppRouter) Use(mw func(http.ResponseWriter, *http.Request) func(http.ResponseWriter, *http.Request)) {
	ar.middleware = append(ar.middleware, mw)
}

func (ar *AppRouter) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	ar.Router.HandleFunc(path, http.HandlerFunc(handler)).Methods("GET")
}

func (ar *AppRouter) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	ar.Router.HandleFunc(path, http.HandlerFunc(handler)).Methods("POST")
}

func (ar *AppRouter) PUT(path string, handler func(http.ResponseWriter, *http.Request)) {
	ar.Router.HandleFunc(path, http.HandlerFunc(handler)).Methods("PUT")
}
