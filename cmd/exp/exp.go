package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
)

const(
	host = "sandbox.smtp.mailtrap.io"
	port = 587
	username = "1a2dab91c58643"
	password = "1e1d4a2d0fe1e3"
)

func main() {
	from := "test@lenslocked.com"
	to := "yasemin.kasikci@ringover.com"
	subject := "this is a test mail"
	plaintext :="this is the body of the email"
	html := `<h1>Hello there buddy</h1><p>this is the email</p><p>Hope you enjoy it</p>`
	
	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/pain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
	
	// defer sender.Close()
	// sender.Send(from,  []string{to}, msg )
	// sender.Send(from,  []string{to}, msg )
}
