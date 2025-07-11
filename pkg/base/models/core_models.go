package base_models

import (
	query_postgres "main/pkg/sql"
	"sync"
)

// Modelo genérico
type Model[T any] struct {
	Name           string
	CollectionName string
	Structure      T
	query_postgres.QueryBuilder[T]
}

// Mapa global de modelos (uso de sync.Map para concurrencia y tipos mixtos)
var models sync.Map // key: string, value: *Model[any]
