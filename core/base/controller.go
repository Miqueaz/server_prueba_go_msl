package base

import (
	"log"
	"server/helpers"
	"sync"
)

// Controller agrupa un modelo con sus métodos
type Controller[T any] struct {
	BaseModelController[T]
}

// Mapa global de modelos (uso de sync.Map para concurrencia y tipos mixtos)
var controllers sync.Map // key: string, value: *Model[any]

// Crear un nuevo modelo y guardarlo
func NewController[T any](model Model[T]) (*Controller[T], bool) {
	controller := &Controller[T]{
		BaseModelController: BaseModelController[T]{
			Model: model,
		},
	}
	helpers.SaveStructure(controller, &controllers)
	return controller, true
}

// Obtener un modelo usando type assertion
func GetController[T any]() (*Controller[T], bool) {

	if value, ok := helpers.LoadStructure[Controller[T]](&controllers); ok {
		return value, true
	}

	// Si no se encuentra el controlador, imprimir un mensaje de error
	if newController, ok := GeneratorController[T](); ok {
		return newController, true
	}

	return nil, false
}

// Generador de controlladores
func GeneratorController[T any]() (*Controller[T], bool) {
	log.Print("Generando controllador")
	if model, ok := GetModel[T](); ok {
		if newController, ok := NewController(*model); ok {
			return newController, true
		}
	}

	return nil, false
}
