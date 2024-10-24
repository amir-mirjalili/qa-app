package user

type CreateUserResponse struct {
	Id          int32
	Name        string
	PhoneNumber string
	Password    string
	Role        int32
}
