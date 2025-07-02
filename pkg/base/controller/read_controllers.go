package base_controller

import (
	"log"
	helpers "main/pkg/base/helpers"
	base_service "main/pkg/base/service"
)

// Obtener un modelo usando type assertion
func GetController[T any](name string) (*Controller[T], bool) {

	if value, ok := helpers.LoadStructure[Controller[T]](&controllers); ok {
		return value, true
	}

	// Si no se encuentra el controlador, imprimir un mensaje de error
	// if newController, ok := GeneratorController[T](); ok {
	// 	return newController, true
	// }

	return nil, false
}

// // Generador de controlladores
func GeneratorController[T any]() (*Controller[T], bool) {
	log.Print("Generando controllador")
	if service, ok := base_service.GetService[T](""); ok {
		newController := NewController(*service, nil)
		return newController, true
	}

	return nil, false
}
