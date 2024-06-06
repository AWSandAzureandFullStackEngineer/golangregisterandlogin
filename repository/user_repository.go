package repository

import (
	"golangservices/config"
	"golangservices/entity"
)

func FindByUsername(username string) (entity.User, error) {
	var user entity.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func SaveUser(user entity.User) error {
	err := config.DB.Create(&user).Error
	return err
}
