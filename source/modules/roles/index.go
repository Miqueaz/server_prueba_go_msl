package roles

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	roles_model "main/source/modules/roles/model"
)

var Service = base_service.NewService[base_service.Default[roles_model.RolesStruct]](*roles_model.Model)

func Init() {
	print("Roles Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/roles")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
