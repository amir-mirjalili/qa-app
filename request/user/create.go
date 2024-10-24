package user

import "qa-app/entity"

type CreateUserRequest struct {
	Name        string
	PhoneNumber string
	Password    string
	Role        entity.Role
}
