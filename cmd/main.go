package main

import (
	server "github.com/c1tad3l/backend-wedo"
	"github.com/c1tad3l/backend-wedo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalln("Error start server: " + err.Error())
	}
}
