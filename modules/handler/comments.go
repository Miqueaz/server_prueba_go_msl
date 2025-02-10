package handler

import (
	"fmt"
	base "server/core/base"

	"go.mongodb.org/mongo-driver/mongo"
)

type commentModel struct {
	base.BaseModelController
}

func (c *commentModel) Read(filter map[string]interface{}) (*mongo.Cursor, error) {
	fmt.Println("Comment")

	return c.BaseModelController.Read(filter)
}

func CommentControllerInit(model base.Model) {
	commentModel := &commentModel{
		BaseModelController: base.BaseModelController{
			Model: model,
		}}
	base.NewControllerWithMethods(commentModel.Model, commentModel)
}
