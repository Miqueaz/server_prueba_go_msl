package base

import (
	"log"
	"server/helpers"
	"sync"
)

// Modelo genérico
type Model[T any] struct {
	Name           string `json:"name"`
	CollectionName string `json:"collectionName"`
	Structure      T      `json:"structure"`
}

// Mapa global de modelos (uso de sync.Map para concurrencia y tipos mixtos)
var models sync.Map // key: string, value: *Model[any]

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

// Obtener un modelo usando type assertion
func GetModel[T any]() (*Model[T], bool) {
	log.Print("Cargando Modelo")

	if value, ok := helpers.LoadStructure[Model[T]](&models); ok {
		log.Print("Modelo Obtenido: ", value.Name)
		return value, true
	}
	log.Print("No se encontro el modelo")
	return nil, false
}
