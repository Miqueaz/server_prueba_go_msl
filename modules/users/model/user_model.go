package user_model

import (
	base "main/pkg/base/controller"
	models "main/pkg/base/models"
)

type Struct struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func init() {
	base.NewController(*models.NewModel[Struct]("User", "users"), nil)
}
