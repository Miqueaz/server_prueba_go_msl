package base_controller

import (
	helpers "main/pkg/base/helpers"
	base_service "main/pkg/base/service"
)

// Crear un nuevo controlador
func NewController[T any](s base_service.Service[T]) *Controller[T] {
	controller := &Controller[T]{Service: s}
	SaveController(controller)
	return controller
}

// Guardar un controlador en el mapa global
func SaveController[T any](controller *Controller[T]) {
	helpers.SaveStructure(controller, &controllers)
}
