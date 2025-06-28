package user_routes

import (
	"main/modules/core/router"
	user_core "main/modules/users/core"
	"net/http"
)

func Init() {
	var r = router.Router()
	r.GET("/users", http.HandlerFunc(user_core.Controller.Read))
}
