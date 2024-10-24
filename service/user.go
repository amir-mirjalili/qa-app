package service

import (
	"qa-app/entity"
	"qa-app/repository"
	"qa-app/request/user"
)

type UserService interface {
	Create(user user.CreateUserRequest)
	FindAll() []entity.User
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (u UserServiceImpl) Create(user user.CreateUserRequest) {
	userEntity := entity.User{
		Name:     user.Name,
		Password: user.Password,
		Role:     user.Role,
	}
	u.UserRepository.Save(userEntity)
}

func (u UserServiceImpl) FindAll() []entity.User {
	return u.UserRepository.FindAll()
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}
