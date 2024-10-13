package services

import (
	"errors"
	internal "go-tutorial/internals/models"
	"go-tutorial/internals/utils"

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

	if err := a.db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.Email == "" {
		return nil, errors.New("no user found with email")
	}

	if utils.CheckPasswordHash(*password, user.Password) == false {
		return nil, errors.New("password is not correct")
	}

	return &user, nil

}

func (a *AuthService) CheckUserExistOrNot(email *string) bool {
	var user internal.User

	if err := a.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return false

	}

	if user.Email != "" {
		return true
	}

	return false

}

func (a *AuthService) Register(email *string, password *string, role *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email can not be null")
	}

	if password == nil {
		return nil, errors.New("password can't be null")
	}

	if a.CheckUserExistOrNot(email) {
		return nil, errors.New("email have already existed by others")
	}

	hashPassword, err := utils.HashPassword(*password)

	if err != nil {
		return nil, err
	}

	var user internal.User

	user.Email = *email
	user.Password = hashPassword

	if *role == "" {
		user.Role = "User"
	} else {
		user.Role = *role
	}

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil

}
