package users

import (
	routes "main/modules/users/routes"
)

func Init() {
	print("Users Module Initialized\n")
	routes.Init()
}
