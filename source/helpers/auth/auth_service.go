package auth

import (
	"context"
	"errors"
	database "main/connection/db/mongo"
	"main/pkg/crypto"
	key "main/security/token"
	"time"
)

type AuthService struct{}

func (s *AuthService) signIn(ctx context.Context, username string, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// user, err := reposditory.FindUser(ctx, username, email)
	// if err != nil {
	// 	return "", err
	// }

	user := struct {
		ID       string
		Username string
		Email    string
		Password string
	}{
		ID:       "exampleUserID",
		Username: username,
		Email:    email,
		Password: "$2a$10$EIX/5z1Zb75",
	}

	if err := crypto.CheckPassword(user.Password, password); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := key.GenerateJWT(user.ID)
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

	// user, err := repository.InsertUser(ctx, username, email, hashedPassword)
	user := struct {
		ID       string
		Username string
		Email    string
		Password string
	}{
		ID:       "newUserID",
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	if err != nil {
		return "", err
	}

	err = database.CreateIndex("users", "email", true)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
