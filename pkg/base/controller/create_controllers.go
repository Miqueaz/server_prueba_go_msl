package base_controller

import (
	helpers "main/pkg/base/helpers"
	models "main/pkg/base/models"
)

// Crear un nuevo modelo y guardarlo
// func NewController[T any](model models.Model[T]) (*Controller[T], bool) {
// 	controller := &Controller[T]{
// 		Model: model,
// 	}
// 	SaveController(controller)
// 	return controller, true
// }

// Crear un nuevo controlador con métodos personalizados
func NewController[T any](model models.Model[T], methods Methods[T]) *Controller[T] {
	controller := &Controller[T]{Model: model, Methods: methods}
	SaveController(controller)
	return controller
}

// Guardar un controlador en el mapa global
func SaveController[T any](controller *Controller[T]) {
	helpers.SaveStructure(controller, &controllers)
}

// InitGeneric inicializa un controlador genérico con un modelo y métodos opcionales.
func Init[T any](c Methods[T]) *Controller[T] {
	if c == nil {
		m, _ := models.GetModel[T]()
		if m == nil {
			return nil
		}
		return NewController(*m, nil)
	}
	return NewController(c.GetModel(), c)
}
