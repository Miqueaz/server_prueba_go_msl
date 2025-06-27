package user_routes

import (
	"main/modules/core/handler/client"
	router "main/modules/core/router"
	user_model "main/modules/users/model"
	base_controller "main/pkg/base/controller"
	"net/http"
)

func Init(c base_controller.Controller[user_model.Struct]) {
	var r = router.Router()
	r.GET("/users", http.HandlerFunc(c.Read))
}

func userHandler(res http.ResponseWriter, req *http.Request) {
	ctx, err := client.Init(res, req)
	if err != nil {
		ctx.InternalServerError(err)
		return
	}

	// Here you would typically fetch users from a database or service
	users := []string{"Alice", "Bob", "Charlie"}

	// Convert []string to []any
	usersAny := make([]any, len(users))
	for i, v := range users {
		usersAny[i] = v
	}

	ctx.Success("Users fetched successfully", usersAny)
}
