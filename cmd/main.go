package main

import (
	server "github.com/c1tad3l/backend-wedo"
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/controllers"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDb()
}
func main() {
	handlers := new(controllers.Handler)
	srv := new(server.Server)
	port := os.Getenv("PORT")
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalln("Error start server: " + err.Error())
	}
}
