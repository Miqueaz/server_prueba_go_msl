package tipoBitacora

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	tipo_bitacora_model "main/source/modules/tipoBitacora/model"
)

var Service = base_service.NewService[base_service.Default[tipo_bitacora_model.TipoBitacoraStruct]](*tipo_bitacora_model.Model)

func Init() {
	print("TipoBitacora Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/tipoBitacora")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
