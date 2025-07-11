package tipoMedidor

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	tipo_medidor_model "main/source/modules/tipoMedidor/model"
)

var Service = base_service.NewService[base_service.Default[tipo_medidor_model.TipoMedidorStruct]](*tipo_medidor_model.Model)

func Init() {
	print("TipoMedidor Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/tipoMedidor")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
