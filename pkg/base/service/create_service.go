package base_service

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
func NewServices[T any](model models.Model[T], methods Methods[T]) *Service[T] {
	var service *Service[T]

	if methods != nil {
		service = &Service[T]{Model: model, Methods: methods}
		SaveService(service)
		return service
	} else {
		service = &Service[T]{Model: model, Methods: &Service[T]{}}
		SaveService(service)
		return service
	}
}

// Guardar un controlador en el mapa global
func SaveService[T any](service *Service[T]) {
	helpers.SaveStructure(service, &services)
}

// InitGeneric inicializa un controlador genérico con un modelo y métodos opcionales.
func Init[T any](c Methods[T]) *Service[T] {
	if c == nil {
		m, _ := models.GetModel[T]()
		if m == nil {
			return nil
		}
		return NewServices(*m, nil)
	}
	return NewServices(c.GetModel(), c)
}
