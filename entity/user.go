package entity

type User struct {
	Id          uint64 `gorm:"primary_key" json:"id"`
	Name        string `gorm:"size:256" json:"name"`
	PhoneNumber string `gorm:"size:12" json:"phone_number"`
	Password    string `gorm:"size:512" json:"password"`
	Role        Role   `json:"role"`
}
