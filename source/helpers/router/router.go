package router

import (
	"main/pkg/base/router"
	"main/pkg/client"
	"net/http"

	"github.com/rs/cors"
)

var rout = router.Router()

func NewRoute(path string) *router.GroupRouter {
	return rout.Group(path)
}

func Router() *router.AppRouter {
	return rout
}

func init() {
	rout.GET("/", func(w http.ResponseWriter, r *http.Request) {
		handlerRouter(IndexRoute).ServeHTTP(w, r)
	})
}

func handlerRouter(handler http.HandlerFunc) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT"},
		AllowedHeaders: []string{"Authorization", "Content-Type", "application/json"},
	}).Handler(handler)
}

func IndexRoute(res http.ResponseWriter, req *http.Request) {
	ctx, err := client.Init(res, req)
	if err != nil {
		ctx.InternalServerError(err)
	}

	ctx.Success("Welcome to my API REST with GO", nil)
}
