package users

import (
	"github.com/google/wire"
)

type UserService interface {
	CreateOne(*User) (*User, error)
	FindMany() ([]User, error)
	FindOne(id int) (*User, error)
	UpdateOne(id int, user *User) (*User, error)
	DeleteOne(id int) (*User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{repository: userRepository}
}

var SetService = wire.NewSet(NewUserService)

func (service *userService) CreateOne(user *User) (*User, error) {
	result, err := service.repository.Create(*user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *userService) FindMany() ([]User, error) {
	users, err := service.repository.Find()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *userService) FindOne(id int) (*User, error) {
	result, err := service.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *userService) UpdateOne(id int, user *User) (*User, error) {
	result, err := service.repository.UpdateOne(id, *user)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (service *userService) DeleteOne(id int) (*User, error) {
	user, err := service.repository.DeleteOne(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
