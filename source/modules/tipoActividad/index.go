package tipoActividad

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	tipo_actividad_model "main/source/modules/tipoActividad/model"
)

var Service = base_service.NewService[base_service.Default[tipo_actividad_model.TipoActividadStruct]](*tipo_actividad_model.Model)

func Init() {
	print("TipoActividad Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/tipoActividad")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
