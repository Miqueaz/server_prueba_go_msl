package reportes

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	reportes_model "main/source/modules/reportes/model"
)

var Service = base_service.NewService[base_service.Default[reportes_model.ReportesStruct]](*reportes_model.Model)

func Init() {
	print("Reportes Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/reportes")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
