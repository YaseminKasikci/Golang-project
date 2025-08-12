package main

import (
	"github/yaseminkasikci/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "1a2dab91c58643"
	password = "1e1d4a2d0fe1e3"
)

func main() {
	email := models.Email{
		From:      "test@lenslocked.com",
		To:        "yasemin.kasikci@ringover.com",
		Subject:   "this is a test mail",
		Plaintext: "this is the body of the email",
		HTML:      `<h1>Hello there buddy</h1><p>this is the email</p><p>Hope you enjoy it</p>`,
	}
	es := models.NewEmailService(models.STMPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err := es.Send(email)
	if err != nil {
		panic(err)
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
}
