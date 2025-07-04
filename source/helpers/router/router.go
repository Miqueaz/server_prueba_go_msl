package router

import (
	"main/pkg/base/router"
	"main/pkg/client"
	"net/http"

	"github.com/gin-contrib/cors"
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
		IndexRoute(w, r)
	})
	rout.SetTrustedProxies([]string{"0.0.0.0/0"})
	rout.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT"},
		AllowHeaders: []string{"Authorization", "Content-Type", "application/json"},
	}))

}

func IndexRoute(res http.ResponseWriter, req *http.Request) {
	ctx, err := client.Init(res, req)
	if err != nil {
		ctx.InternalServerError(err)
	}

	ctx.Success("Welcome to my API REST with GO", nil)
}
