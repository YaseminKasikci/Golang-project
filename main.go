package main

import (
	"fmt"
	"github/yaseminkasikci/lenslocked/controllers"
	"github/yaseminkasikci/lenslocked/models"
	"github/yaseminkasikci/lenslocked/templates"
	"github/yaseminkasikci/lenslocked/views"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	// "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"contact.gohtml", "tailwind.gohtml",
	))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS,
		"faq.gohtml", "tailwind.gohtml",
	))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	// logRequest := func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		log.Println("Method:", r.Method)
	// 		log.Println("Origin:", r.Header.Get("Origin"))
	// 		log.Println("Referer:", r.Header.Get("Referer"))
	// 		log.Println("Token from cookie:", r.Header.Get("Cookie"))
	// 		log.Println("Token from header:", r.Header.Get("X-CSRF-Token"))
	// 		next.ServeHTTP(w, r)
	// 	})
	// }

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false), // OK pour dev
	)

	err = http.ListenAndServe(":3000", csrfMw(r))
	// err = http.ListenAndServe(":3000", logRequest(csrfMw(r)))
	if err != nil {
		log.Fatal(err)
	}
}
