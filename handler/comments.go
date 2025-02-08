package handler

import (
	Structures "server/core/structures"
)

type commentModel struct {
	Structures.BaseModelController
}

func (c *commentModel) Read(filter map[string]interface{}) (interface{}, error) {
	return c.Read(filter)
}

func commentController(model Structures.Model) {
	methods := map[string]Structures.Methods{
		model.CollectionName: &commentModel{
			BaseModelController: *Structures.GetController(model.CollectionName),
		},
	}
	Structures.SaveMethods(methods)
}
