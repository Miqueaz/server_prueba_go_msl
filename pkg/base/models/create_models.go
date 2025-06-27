package base_models

import (
	"log"
	helpers "main/pkg/base/helpers"
)

// Crear un nuevo modelo y guardarlo
func NewModel[T any](name string, collectionName string) *Model[T] {

	model := &Model[T]{
		Name:           name,
		CollectionName: collectionName,
	}
	helpers.SaveStructure(model, &models)

	if m, ok := helpers.LoadStructure[Model[T]](&models); ok {
		log.Printf("Modelo '%s' creado y almacenado con éxito.\n", m.Name)
		return model
	}

	return nil
}
