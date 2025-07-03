package auth

import (
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
)

func AuthRouter() {
	r := router.NewRoute("/auth")

	r.GET("/sign-in", SignIn, middleware.JWTMiddleware(), middleware.ValidatorMiddleware[AuthDTO]())
	r.POST("/sign-up", SignUp, middleware.ValidatorMiddleware[AuthDTO]())
}
