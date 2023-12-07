package main

import (
	server "github.com/c1tad3l/backend-wedo"
	"github.com/c1tad3l/backend-wedo/pkg/config"
	"github.com/c1tad3l/backend-wedo/pkg/controllers"
	"log"
	"net/smtp"
)

func main() {
	info, _ := config.LoadConfig()
	auth := smtp.PlainAuth("", info.EmailSenderName, info.EmailSenderAddress, info.EmailSenderPassword)
	err := smtp.SendMail("smtp.yandex.ru:465", auth, info.EmailSenderAddress, []string{"dt3csgo@mail.ru"}, []byte("Текст письма."))
	if err != nil {
		log.Fatal(err)
	}
	//info, _ := config.LoadConfig()
	//fmt.Printf("%s %s %s", info.EmailSenderName, info.EmailSenderAddress, info.EmailSenderPassword)
	//sender := mail.NewGmailSender(info.EmailSenderName, info.EmailSenderAddress, info.EmailSenderPassword)
	//
	//subject := "A test email"
	//content := `
	//<h1>Hello world</h1>
	//`
	//to := []string{"zloymolodoy88@gmail.com"}
	//
	//fmt.Print(sender.SendEmail(subject, content, to, nil, nil, nil))

	handlers := new(controllers.Handler)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalln("Error start server: " + err.Error())
	}
}
