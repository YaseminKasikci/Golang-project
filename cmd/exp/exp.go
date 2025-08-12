package main

import (
	"os"

	"github.com/go-mail/mail/v2"
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
}
