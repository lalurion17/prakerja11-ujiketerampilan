package configs

import "main.go/models/user"

func initMigrate() {
	DB.AutoMigrate(&user.User{})
}
