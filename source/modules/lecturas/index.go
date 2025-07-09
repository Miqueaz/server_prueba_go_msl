package lecturas

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	lecturas_model "main/source/modules/lecturas/model"
)

var Service = base_service.NewService[base_service.Default[lecturas_model.LecturasStruct]](*lecturas_model.Model)

func Init() {
	print("Lecturas Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/lecturas")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
