package main

import (
	server "github.com/c1tad3l/backend-wedo"
	"github.com/c1tad3l/backend-wedo/pkg/config"
	"github.com/c1tad3l/backend-wedo/pkg/controllers"
	"github.com/c1tad3l/backend-wedo/pkg/initializers"
	"log"
)

func init() {
	initializers.ConnectDb()
}
func main() {
	handlers := new(controllers.Handler)
	srv := new(server.Server)

	if err := srv.Run(info.Port, handlers.InitRoutes()); err != nil {
		log.Fatalln("Error start server: " + err.Error())
	}
}
