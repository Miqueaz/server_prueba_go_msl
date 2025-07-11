package tipoAlerta

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	tipo_alerta_model "main/source/modules/tipoAlerta/model"
)

var Service = base_service.NewService[base_service.Default[tipo_alerta_model.TipoAlertaStruct]](*tipo_alerta_model.Model)

func Init() {
	print("TipoAlerta Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/tipoAlerta")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
