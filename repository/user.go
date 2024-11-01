package repository

import (
	"errors"
	"gorm.io/gorm"
	"qa-app/entity"
	"qa-app/response"
)

type UserRepository interface {
	Save(user entity.User)
	Update(user entity.User)
	Delete(userId int)
	FindById(userId int) (user entity.User, err error)
	FindAll() []response.UserResponse
	FindByPhone(phone string) (user entity.User, err error)
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u UserRepositoryImpl) FindByPhone(phone string) (user entity.User, err error) {
	result := u.Db.Where("phone_number= ?", phone).First(&user)
	if result.Error != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (u UserRepositoryImpl) Save(user entity.User) {
	result := u.Db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u UserRepositoryImpl) Update(user entity.User) {
	result := u.Db.Updates(user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u UserRepositoryImpl) Delete(userId int) {
	result := u.Db.Where("id = ?", userId).Delete(&entity.User{})
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u UserRepositoryImpl) FindById(userId int) (user entity.User, err error) {
	result := u.Db.Find(&user, userId)
	if result == nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (u UserRepositoryImpl) FindAll() []response.UserResponse {
	var users []response.UserResponse
	u.Db.Find(&users)
	return users
}
