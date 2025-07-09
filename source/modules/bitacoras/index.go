package bitacoras

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	bitacoras_model "main/source/modules/bitacoras/model"
)

var Service = base_service.NewService[base_service.Default[bitacoras_model.BitacorasStruct]](*bitacoras_model.Model)

func Init() {
	print("Bitacoras Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/bitacoras")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
