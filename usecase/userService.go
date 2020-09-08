package usecase

import (
	"../infrastructure"
	"log"
)

type UserService interface {
	GetUsers() []UserResponse
	GetUser(name string) UserResponse
}

type userServiceImpl struct {
	repo infrastructure.UserRepository
}

func CreateUserService(repository infrastructure.UserRepository) UserService {
	return &userServiceImpl{repository}
}

func (service *userServiceImpl) GetUsers() []UserResponse {
	users, err := service.repo.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	var userResponses []UserResponse
	for _, u := range users {
		res := UserResponse{
			Id:   u.Id,
			Name: u.Name,
			Age:  u.Age,
		}
		userResponses = append(userResponses, res)
	}

	return userResponses
}

func (service *userServiceImpl) GetUser(name string) UserResponse {

	user, err := service.repo.GetUser(name)
	if err != nil {
		log.Fatal(err)
	}

	return UserResponse{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
	}
}
