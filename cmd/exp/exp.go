package main

import (
	"fmt"
	"github/yaseminkasikci/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "1a2dab91c58643"
	password = "1e1d4a2d0fe1e3"
)

func main() {

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err := es.ForgotPassword(
		"yasemin.kasikci@ringover.com",
		"https://lenslocked.com/reset-pw?token=abc123")
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent")
}

// msg := mail.NewMessage()
// msg.SetHeader("To", to)
// msg.SetHeader("From", from)
// msg.SetHeader("Subject", subject)
// msg.SetBody("text/pain", plaintext)
// msg.AddAlternative("text/html", html)
// msg.WriteTo(os.Stdout)

// dialer := mail.NewDialer(host, port, username, password)
// err := dialer.DialAndSend(msg)
// if err != nil {
// 	panic(err)
// }
// fmt.Println("Message sent")

// defer sender.Close()
// sender.Send(from,  []string{to}, msg )
// sender.Send(from,  []string{to}, msg )
// }
