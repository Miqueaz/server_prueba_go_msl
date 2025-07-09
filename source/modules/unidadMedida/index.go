package unidadMedida

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	unidad_medida_model "main/source/modules/unidadMedida/model"
)

var Service = base_service.NewService[base_service.Default[unidad_medida_model.UnidadMedidaStruct]](*unidad_medida_model.Model)

func Init() {
	print("UnidadMedida Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/unidadMedida")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
