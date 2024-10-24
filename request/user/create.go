package user

import "qa-app/entity"

type CreateUserRequest struct {
	Name        string      `json:"name"`
	PhoneNumber string      `json:"phone_number"`
	Password    string      `json:"password"`
	Role        entity.Role `json:"role"`
}
