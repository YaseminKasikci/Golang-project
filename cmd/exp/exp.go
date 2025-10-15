package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dropboxID := os.Getenv("DROPBOX_APP_ID")
	dropboxSecret := os.Getenv("DROPBOX_APP_SECRET")
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     dropboxID,
		ClientSecret: dropboxSecret,
		Scopes:       []string{"files.metadata.read", "files.content.read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.dropbox.com/oauth2/authorize",
			TokenURL: "https://api.dropboxapi.com/oauth2/token",
		},
	}

	// use PKCE to protect against CSRF attacks
	// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
	verifier := oauth2.GenerateVerifier()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)
	fmt.Printf("VOnce you have a code, paste it and press enter: ")

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}

	enc :=json.NewEncoder(os.Stdout)
	enc.SetIndent(""," ")
	enc.Encode(tok)
	// client := conf.Client(ctx, tok)
	// resp, err := client.Post("https://api.dropboxapi.com/2/files/list_folder", "application/json", strings.NewReader(`{
	// 	"path": ""
	//  }`))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)

}

// -------------------------------------------------------------------------------

// package main

// import (
// 	"io"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	sketchyURL := "http://localhost:3000/galleries/2/images/../images-1/test.png"
// 	resp, err := http.Get(sketchyURL)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	io.Copy(os.Stdout, resp.Body)
// }

// // func main() {
// // 	gs := models.GalleryService{}
// // 	fmt.Println(gs.Images(2))
// // }

// // func main() {
// // 	err := godotenv.Load()
// // 	if err != nil {
// // 		log.Fatal("Error loading .env file")
// // 	}

// // 	host := os.Getenv("SMTP_HOST")
// // 	portStr := os.Getenv("SMTP_PORT")
// // 	port, err := strconv.Atoi(portStr)
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	username := os.Getenv("SMTP_USERNAME")
// // 	password := os.Getenv("SMTP_PASSWORD")

// // 	configSMTP := models.SMTPConfig{
// // 		Host:     host,
// // 		Port:     port,
// // 		Username: username,
// // 		Password: password,
// // 	}
// // 	// spew.Dump(configSMTP)

// // 	es := models.NewEmailService(configSMTP)
// // 	err = es.ForgotPassword(
// // 		"yasemin.kasikci@ringover.com",
// // 		"https://lenslocked.com/reset-pw?token=abc123")
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	fmt.Println("Email sent")
// // }
