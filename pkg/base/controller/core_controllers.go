package base_controller

import (
	models "main/pkg/base/models"
	"main/pkg/hooks"
	"net/http"
	"sync"
)

// Controller agrupa un modelo con sus métodos
type Controller[T any] struct {
	models.Model[T]
	Methods[T]
	hooks.Hookable
	hooks.Cleaners
}

// Methods define los métodos CRUD
type Methods[T any] interface {
	Read(res http.ResponseWriter, req *http.Request)
	Insert(data T) error
	Update(filter map[string]any, data T) error
	Delete(filter map[string]any) error
	GetModel() models.Model[T]
}

// Mapa global de modelos (uso de sync.Map para concurrencia y tipos mixtos)
var controllers sync.Map // key: string, value: *Model[any]
