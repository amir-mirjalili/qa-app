package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"qa-app/entity"
)

func New() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:12345678@tcp(127.0.0.1:3306)/qa-game?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		return
	}

}
