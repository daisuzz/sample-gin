package usecase

import (
	"../infrastructure"
)

type UserService interface {
	GetUsers() []string
	GetUser(name string) []string
}

type userServiceImpl struct {
	repo infrastructure.UserRepository
}

func CreateUserService(repository infrastructure.UserRepository) UserService {
	return &userServiceImpl{repository}
}

func (service *userServiceImpl) GetUsers() []string {
	return service.repo.GetUsers()
}

func (service *userServiceImpl) GetUser(name string) []string {
	return service.repo.GetUser(name)
}
