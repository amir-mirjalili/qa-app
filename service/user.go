package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"qa-app/entity"
	hashGenerator "qa-app/pkg/md5hash"
	"qa-app/pkg/richerror"
	"qa-app/repository"
	"qa-app/request/user"
	"qa-app/response"
	"time"
)

type UserService interface {
	Create(user user.CreateUserRequest)
	FindAll() []response.UserResponse
	CreateAccessToken(user user.CreateTokenReq) (string, error)
	CreateRefreshToken(user user.CreateTokenReq) (string, error)
	Login(user user.UserLoginReq) (res response.UserLoginResponse, err error)
	GetUserByPhoneNumber(phoneNumber string) (res response.UserResponse, err error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (u UserServiceImpl) GetUserByPhoneNumber(phoneNumber string) (res response.UserResponse, err error) {
	findByPhone, _ := u.UserRepository.FindByPhone(phoneNumber)
	return response.UserResponse{
		Id:          findByPhone.Id,
		Name:        findByPhone.Name,
		PhoneNumber: findByPhone.PhoneNumber,
		Password:    findByPhone.Password,
		Role:        findByPhone.Role,
	}, nil
}

func (u UserServiceImpl) Login(req user.UserLoginReq) (res response.UserLoginResponse, err error) {
	const op = "user.Login"

	item, err := u.UserRepository.FindByPhone(req.PhoneNumber)
	if err != nil {
		return response.UserLoginResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}
	if item.Password != hashGenerator.HashGenerator(req.Password) {
		return response.UserLoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	createAccess := user.CreateTokenReq{
		Id:   item.Id,
		Role: item.Role,
	}
	accessToken, err := u.CreateAccessToken(createAccess)
	if err != nil {
		return response.UserLoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	refreshToken, err := u.CreateRefreshToken(createAccess)
	if err != nil {
		return response.UserLoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return response.UserLoginResponse{
		User: response.UserResponse{
			Id:          item.Id,
			Name:        item.Name,
			PhoneNumber: item.PhoneNumber,
			Password:    item.Password,
			Role:        item.Role,
		},
		Tokens: response.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil

}

func (u UserServiceImpl) CreateAccessToken(user user.CreateTokenReq) (string, error) {
	return u.createToken(user.Id, user.Role)
}

func (u UserServiceImpl) CreateRefreshToken(user user.CreateTokenReq) (string, error) {
	return u.createToken(user.Id, user.Role)
}

func (u UserServiceImpl) Create(user user.CreateUserRequest) {
	userEntity := entity.User{
		Name:        user.Name,
		Password:    hashGenerator.HashGenerator(user.Password),
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
	}
	u.UserRepository.Save(userEntity)
}

func (u UserServiceImpl) FindAll() []response.UserResponse {
	return u.UserRepository.FindAll()
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (u UserServiceImpl) createToken(userID uint64, role entity.Role) (string, error) {
	// create a signer for rsa 256

	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "test",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3000 * time.Second)),
		},
		UserID: userID,
		Role:   role,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte("abcdefg"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
