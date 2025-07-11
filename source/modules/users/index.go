package users

import (
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	user_service "main/source/modules/users/services"
)

func Init() {
	print("Users Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/users")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", user_service.Service.Read)
	// r.POST("/", user_service.Service.Insert)
	r.GET("/:id", user_service.Service.ReadOne)
	r.PUT("/:id", user_service.Service.Update)
	r.DELETE("/:id", user_service.Service.Delete)
}
