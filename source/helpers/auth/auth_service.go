package auth

import (
	"context"
	"errors"
	"main/pkg/crypto"
	key "main/security/token"
	user_model "main/source/modules/users/models"
	user_service "main/source/modules/users/services"
	"time"
)

type AuthService struct{}

func (s *AuthService) signIn(ctx context.Context, matricula string, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	users, err := user_model.Model.Find.Where("matricula", "=", matricula).Exec(ctx)
	user := users[0]
	if err != nil {
		return "", err
	}

	if err := crypto.CheckPassword(user.Contrasena, password); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := key.GenerateJWT(user.Matricula)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthService) signUp(ctx context.Context, username string, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	hashedPassword, err := crypto.EncryptPassword(password)
	if err != nil {
		return "", err
	}

	user, err := user_service.Service.Insert(user_model.UserStruct{
		PrimerNombre:    username,
		SegundoNombre:   nil,
		PrimerApellido:  username,
		SegundoApellido: nil,
		Matricula:       username,
		Correo:          email,
		Contrasena:      hashedPassword,
		Rol:             1, // Assuming 1 is the default role for new users
	})

	if err != nil {
		return "", err
	}

	return user.Matricula, nil
}
