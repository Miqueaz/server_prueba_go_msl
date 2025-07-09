package tipoArchivo

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	tipo_archivo_model "main/source/modules/tipoArchivo/model"
)

var Service = base_service.NewService[base_service.Default[tipo_archivo_model.TipoArchivoStruct]](*tipo_archivo_model.Model)

func Init() {
	print("TipoArchivo Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/tipoArchivo")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
