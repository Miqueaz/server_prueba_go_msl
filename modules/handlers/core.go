package handlers

import (
	"server/core/base"
	"sync"
)

type Controller struct {
	base.Controller[any]
}

var controllers sync.Map

func pushController(name string, controller any) {
	controllers.Store(name, controller)
}

func GetController(name string) (any, bool) {
	if value, ok := controllers.Load(name); ok {

		if controller, ok := value.(Controller); ok {
			return controller, true
		}
	}
	var zero Controller
	return zero, false
}
