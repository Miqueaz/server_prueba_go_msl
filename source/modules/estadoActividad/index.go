package estadoActividad

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	estado_actividad_model "main/source/modules/estadoActividad/model"
)

var Service = base_service.NewService[base_service.Default[estado_actividad_model.EstadoActividadStruct]](*estado_actividad_model.Model)

func Init() {
	print("EstadoActividad Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/estadoActividad")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
