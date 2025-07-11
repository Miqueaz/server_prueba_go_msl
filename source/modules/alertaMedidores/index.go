package alertaMedidores

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	alerta_medidores_model "main/source/modules/alertaMedidores/model"
)

var Service = base_service.NewService[base_service.Default[alerta_medidores_model.AlertaMedidoresStruct]](*alerta_medidores_model.Model)

func Init() {
	print("AlertaMedidores Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/alertaMedidores")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
