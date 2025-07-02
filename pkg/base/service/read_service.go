package base_service

import (
	"log"
	helpers "main/pkg/base/helpers"
	models "main/pkg/base/models"
)

// Obtener un modelo usando type assertion
func GetService[T any](name string) (*Service[T], bool) {

	if value, ok := helpers.LoadStructure[Service[T]](&services); ok {
		return value, true
	}

	// Si no se encuentra el controlador, imprimir un mensaje de error
	if newController, ok := GeneratorService[T](); ok {
		return newController, true
	}

	return nil, false
}

// Generador de controlladores
func GeneratorService[T any]() (*Service[T], bool) {
	log.Print("Generando controllador")
	if model, ok := models.GetModel[T](); ok {
		newService := NewServices(*model, nil)
		return newService, true
	}

	return nil, false
}
