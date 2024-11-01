package user

import "qa-app/entity"

type CreateTokenReq struct {
	Id   uint64      `json:"id"`
	Role entity.Role `json:"role"`
}
