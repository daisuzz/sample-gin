package infrastructure

type Users []string

type UserRepository interface {
	GetUsers() Users
	GetUser(name string) Users
}

type userRepositoryImpl struct {
	Users
}

func CreateUserRepository() UserRepository {
	return &userRepositoryImpl{
		Users: []string{"ichiro", "jiro"},
	}
}

func (u *userRepositoryImpl) GetUsers() Users {
	return u.Users
}

func (u *userRepositoryImpl) GetUser(name string) Users {
	return u.Users
}
