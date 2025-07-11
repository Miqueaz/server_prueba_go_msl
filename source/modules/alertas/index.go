package alertas

import (
	base_service "main/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	alertas_model "main/source/modules/alertas/model"
)

var Service = base_service.NewService[base_service.Default[alertas_model.AlertasStruct]](*alertas_model.Model)

func Init() {
	print("Alertas Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/alertas")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
