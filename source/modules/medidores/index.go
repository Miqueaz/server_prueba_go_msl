package medidores

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	medidores_model "main/source/modules/medidores/model"
)

var Service = base_service.NewService[base_service.Default[medidores_model.MedidoresStruct]](*medidores_model.Model)

func Init() {
	print("Medidores Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/medidores")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
