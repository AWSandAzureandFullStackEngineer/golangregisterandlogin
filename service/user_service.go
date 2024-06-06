package service

import (
	"golangservices/entity"
	"golangservices/repository"
	"golangservices/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user entity.User) (string, error) {
	existingUser, err := repository.FindByUsername(user.Username)
	if err == nil && existingUser.ID != 0 {
					return "", errors.New("username already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	err = repository.SaveUser(user)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Authenticate(username, password string) (string, error) {
	user, err := repository.FindByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
