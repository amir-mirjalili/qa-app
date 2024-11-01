package response

import "qa-app/entity"

type UserResponse struct {
	Id          uint64 `json:"id"`
	Name        string
	PhoneNumber string
	Password    string
	Role        entity.Role
}
