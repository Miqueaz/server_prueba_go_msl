package base_module

import (
	controller "main/pkg/base/controller"
	"sync"
)

// Controller agrupa un modelo con sus métodos
type Module[T any] struct {
	controller.Controller[T]
}

// Mapa global de modelos (uso de sync.Map para concurrencia y tipos mixtos)
var Modules sync.Map // key: string, value: *Model[any]
