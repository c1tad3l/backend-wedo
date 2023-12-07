package main

import (
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/models/users"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&users.UserEstimates{})
	initializers.DB.AutoMigrate(&users.User{})
	initializers.DB.AutoMigrate(&users.UserParents{})

}
