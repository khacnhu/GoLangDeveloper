package services

import (
	"errors"
	internal "go-tutorial/internals/models"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func (a *AuthService) InitService(database *gorm.DB) {
	a.db = database
	a.db.AutoMigrate(&internal.User{})
}

func (a *AuthService) Login(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email can not be null")
	}

	if password == nil {
		return nil, errors.New("password can't be null")
	}

	var user internal.User

	user.Email = *email
	user.Password = *password

	if err := a.db.Where("email=?", email).Where("password =?", password).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (a *AuthService) Register(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email can not be null")
	}

	if password == nil {
		return nil, errors.New("password can't be null")
	}

	var user internal.User

	user.Email = *email
	user.Password = *password

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}
