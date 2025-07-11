package actividadesUsuario

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	actividades_usuario_model "main/source/modules/actividadesUsuario/model"
)

var Service = base_service.NewService[base_service.Default[actividades_usuario_model.ActividadesUsuarioStruct]](*actividades_usuario_model.Model)

func Init() {
	print("ActividadesUsuario Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/actividadesUsuario")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
