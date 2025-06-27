package base_models

import "sync"

// Modelo genérico
type Model[T any] struct {
	Name           string `json:"name"`
	CollectionName string `json:"collectionName"`
	Structure      T      `json:"structure"`
}

// Mapa global de modelos (uso de sync.Map para concurrencia y tipos mixtos)
var models sync.Map // key: string, value: *Model[any]
