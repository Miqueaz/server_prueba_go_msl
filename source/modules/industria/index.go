package industria

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	industria_model "main/source/modules/industria/model"
)

var Service = base_service.NewService[base_service.Default[industria_model.IndustriaStruct]](*industria_model.Model)

func Init() {
	print("Industria Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/industria")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
