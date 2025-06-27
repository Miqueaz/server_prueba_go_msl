package users

import (
	user_controller "main/modules/users/controller"
	routes "main/modules/users/routes"
	base_controller "main/pkg/base/controller"
)

func Init() {
	print("Users Module Initialized\n")

	controller := base_controller.Init(user_controller.UserController{}.Methods)

	routes.Init(*controller)
}
