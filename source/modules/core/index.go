package modules

import "main/source/modules/users"

func init() {
	NewModule(users.Init)
}
