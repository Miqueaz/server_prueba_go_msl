package modulosRol

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	modulos_rol_model "main/source/modules/modulosRol/model"
)

var Service = base_service.NewService[base_service.Default[modulos_rol_model.ModulosRolStruct]](*modulos_rol_model.Model)

func Init() {
	print("PermisosRol Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/modulosRol")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
