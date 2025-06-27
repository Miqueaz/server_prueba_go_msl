package auth

import (
	"main/modules/core/middleware"

	"github.com/gorilla/mux"
)

func AuthRouter(router *mux.Router) {
	r := router.PathPrefix("/auth").Subrouter()
	protected := r
	protected.Use(middleware.JWTMiddleware)

	r.HandleFunc("/sign-in", middleware.ValidatorMiddleware[AuthDTO](SignIn)).Methods("POST")
	protected.HandleFunc("/sign-up", middleware.ValidatorMiddleware[AuthDTO](SignUp)).Methods("POST")
}
