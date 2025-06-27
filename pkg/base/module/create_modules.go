package base_module

import (
	controller "main/pkg/base/controller"
	models "main/pkg/base/models"
)

// Crear un nuevo modelo y guardarlo
func NewModule[T any]() Module[T] {
	newModel := models.NewModel[T]("Module", "modules")
	newController := controller.NewController(*newModel, nil)
	newModule := Module[T]{
		Controller: *newController,
	}

	return newModule
}
