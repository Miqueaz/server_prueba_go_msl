package base_controller

import (
	base_models "main/pkg/base/models"
	base_service "main/pkg/base/service"
	"net/http"
	"sync"
)

// Controller agrupa un modelo con sus métodos
type Controller[T any] struct {
	base_service.Service[T]
	Methods[T]
}

// Methods define los métodos CRUD
type Methods[T any] interface {
	Read(res http.ResponseWriter, req *http.Request)
	Insert(res http.ResponseWriter, req *http.Request)
	Update(res http.ResponseWriter, req *http.Request)
	Delete(res http.ResponseWriter, req *http.Request)
	GetModel() base_models.Model[T]
}

// Mapa global de modelos (uso de sync.Map para concurrencia y tipos mixtos)
var controllers sync.Map // key: string, value: *Model[any]
