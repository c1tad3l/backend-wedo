package main

import (
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
)

func init() {
	initializers.ConnectDb()
}

func main() {

	err := initializers.DB.AutoMigrate(&users.User{})
	if err != nil {
		return
	}
	err = initializers.DB.AutoMigrate(&users.Email{})
	if err != nil {
		return
	}
}
