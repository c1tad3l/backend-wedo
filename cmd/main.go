package main

import (
	"fmt"
	server "github.com/c1tad3l/backend-wedo"
	"github.com/c1tad3l/backend-wedo/initializers"
	"github.com/c1tad3l/backend-wedo/pkg/controllers"
	"log"
	"net/smtp"
	"strings"
)

func init() {
	initializers.ConnectDb()
}
func SendToMail(user, password, host, subject, body, mailtype, replyToAddress string, to, cc, bcc []string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	cc_address := strings.Join(cc, ";")
	bcc_address := strings.Join(bcc, ";")
	to_address := strings.Join(to, ";")
	msg := []byte("To: " + to_address + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + cc_address + "\r\nBcc: " + bcc_address + "\r\n" + content_type + "\r\n\r\n" + body)

	send_to := MergeSlice(to, cc)
	send_to = MergeSlice(send_to, bcc)
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
func main() {
	user := "kairexblin@gmail.com"
	password := "qhjpxrfnuylclngl"
	host := "smtp.gmail.com:25"
	to := []string{"fellowgram@gmail.com", "test2***@example.net"}
	cc := []string{"test3***@example.net", "test4***@example.net"}
	bcc := []string{"test5***@example.net", "test6***@example.net"}

	subject := "test Golang to sendmail"
	mailtype := "html"
	replyToAddress := "test7***@example.net"

	body := `
	<html>
	<body>
	<h3>
		"Test send to email"
	</h3>
	</body>
	</html>
		`
	fmt.Println("отправка email")
	err := SendToMail(user, password, host, subject, body, mailtype, replyToAddress, to, cc, bcc)
	if err != nil {
		fmt.Println("ошибка!")
		fmt.Println(err)
	} else {
		fmt.Println("успех!")
	}
	handlers := new(controllers.Handler)
	srv := new(server.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalln("Error start server: " + err.Error())
	}
}
func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}
