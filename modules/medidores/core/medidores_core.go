// Archivo generado automáticamente para el módulo medidores (core)
package medidores_core

import (
	user_model "main/modules/medidores/model"
	base_controller "main/pkg/base/controller"
	base_service "main/pkg/base/service"
)

var Controller = base_controller.NewController(
	*base_service.Init[user_model.Struct](nil),
)
var Service = base_service.Init[user_model.Struct](nil)
