package user

type UserLoginReq struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
