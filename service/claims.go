package service

import (
	"github.com/golang-jwt/jwt/v4"
	"qa-app/entity"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint64      `json:"user_id"`
	Role   entity.Role `json:"role"`
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
