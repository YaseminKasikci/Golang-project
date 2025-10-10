package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	sketchyURL := "http://localhost:3000/galleries/2/images/../images-1/test.png"
	resp, err := http.Get(sketchyURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

// func main() {
// 	gs := models.GalleryService{}
// 	fmt.Println(gs.Images(2))
// }

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	host := os.Getenv("SMTP_HOST")
// 	portStr := os.Getenv("SMTP_PORT")
// 	port, err := strconv.Atoi(portStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	username := os.Getenv("SMTP_USERNAME")
// 	password := os.Getenv("SMTP_PASSWORD")

// 	configSMTP := models.SMTPConfig{
// 		Host:     host,
// 		Port:     port,
// 		Username: username,
// 		Password: password,
// 	}
// 	// spew.Dump(configSMTP)

// 	es := models.NewEmailService(configSMTP)
// 	err = es.ForgotPassword(
// 		"yasemin.kasikci@ringover.com",
// 		"https://lenslocked.com/reset-pw?token=abc123")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Email sent")
// }
