package infrastructure

import "errors"

type Users []User

type UserRepository interface {
	GetUsers() (Users, error)
	GetUser(name string) (User, error)
}

type userRepositoryImpl struct {
	users Users
}

func CreateUserRepository() UserRepository {
	return &userRepositoryImpl{
		users: []User{
			{
				Id:   "1",
				Name: "Ichiro",
				Age:  20,
			},
			{
				Id:   "2",
				Name: "Jiro",
				Age:  30,
			},
		},
	}
}

func (u *userRepositoryImpl) GetUsers() (Users, error) {
	return u.users, nil
}

func (u *userRepositoryImpl) GetUser(id string) (User, error) {

	for _, u := range u.users {
		if u.Id == id {
			return u, nil
		}
	}

	return User{}, ErrorNotFound

}

var ErrorNotFound = errors.New("not found")
