package areasIndustriales

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	areas_industriales_model "main/source/modules/areasIndustriales/model"
)

var Service = base_service.NewService[base_service.Default[areas_industriales_model.AreasIndustrialesStruct]](*areas_industriales_model.Model)

func Init() {
	print("AreasIndustriales Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/areasIndustriales")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
