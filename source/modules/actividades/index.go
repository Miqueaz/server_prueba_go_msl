package actividades

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	actividades_model "main/source/modules/actividades/model"
)

var Service = base_service.NewService[base_service.Default[actividades_model.ActividadesStruct]](*actividades_model.Model)

func Init() {
	print("Actividades Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/actividades")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
