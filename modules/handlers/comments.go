package handlers

import (
	"server/core/base"
	"server/modules/models"
)

type comment[T any] struct {
	base.Controller[T]
}

var ctrl, _ = base.GetController[models.Comment]()
var CommentController = comment[models.Comment]{
	Controller: *ctrl,
}
