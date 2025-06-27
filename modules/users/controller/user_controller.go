package user_controller

import (
	user_model "main/modules/users/model"
	base_controller "main/pkg/base/controller"
	"net/http"
)

type UserController struct {
	base_controller.Controller[user_model.Struct]
}

func (user UserController) Read(res http.ResponseWriter, req *http.Request) {
}
