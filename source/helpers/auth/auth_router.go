package auth

import (
	"main/source/helpers/router"
)

func AuthRouter() {
	r := router.NewRoute("/auth")

	r.GET("/sign-in", SignIn)
	r.POST("/sign-up", SignUp)
}
