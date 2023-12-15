package sender

import (
	"github.com/c1tad3l/backend-wedo/pkg/config"
	"net/smtp"
	"strings"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:25"
)

//to := []string{"zloymolodoy88@gmail.com"}
//cc := []string{}
//bcc := []string{}
//
//subject := "test Golang to sendmail"
//mailtype := "html"
//replyToAddress := ""
//
//body := `
//	<html>
//	<body>
//	<h3>
//		Test send to email
//	</h3>
//	</body>
//	</html>`

func SendToMail(subject, body, mailtype, replyToAddress string, to, cc, bcc []string) error {
	env, _ := config.LoadConfig()

	hp := strings.Split(smtpServerAddress, ":")
	auth := smtp.PlainAuth("", env.EmailSenderAddress, env.EmailSenderPassword, hp[0])
	var contentType string

	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	ccAddress := strings.Join(cc, ";")
	bccAddress := strings.Join(bcc, ";")
	toAddress := strings.Join(to, ";")
	msg := []byte("To: " + toAddress + "\r\nFrom: " + env.EmailSenderAddress + "\r\nSubject: " + subject + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + ccAddress + "\r\nBcc: " + bccAddress + "\r\n" + contentType + "\r\n\r\n" + body)

	sendTo := MergeSlice(to, cc)
	sendTo = MergeSlice(sendTo, bcc)
	err := smtp.SendMail(smtpServerAddress, auth, env.EmailSenderAddress, sendTo, msg)
	return err
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}
