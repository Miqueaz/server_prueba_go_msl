package controller

import (
	"server/core/connection/proto"
	"server/core/controller/methods"
)

type server struct {
	proto.UnimplementedBaseServer
}
type BaseController struct{ *methods.BaseMethods }

func (super *BaseController) Base() {
	filter := map[string]interface{}{"name": "John"}
	collectionName := "users"
	super.Load(filter, collectionName)
}

func (server *server) Controller() {
	controller := &BaseController{&methods.BaseMethods{}}
	filter := map[string]interface{}{"name": "John"}
	collectionName := "users"
	controller.Load(filter, collectionName)
}
