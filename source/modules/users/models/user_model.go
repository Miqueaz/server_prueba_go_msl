package user_model

import base_models "main/pkg/base/models"

type UserStruct struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

var Model = base_models.NewModel[UserStruct]("user", "users")
