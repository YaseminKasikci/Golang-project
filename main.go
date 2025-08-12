package main

import (
	"fmt"
	"github/yaseminkasikci/lenslocked/controllers"
	"github/yaseminkasikci/lenslocked/migrations"
	"github/yaseminkasikci/lenslocked/models"
	"github/yaseminkasikci/lenslocked/templates"
	"github/yaseminkasikci/lenslocked/views"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	// "github.com/go-chi/chi/v5/middleware"
)

func main() {

	//database
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Change the following line of code
	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// setup services
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// SETup middleware
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false), // OK pour dev
		csrf.TrustedOrigins([]string{"localhost:3000"}),
	)

	// setup controllers
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}

	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	// setup our router and routes
	r := chi.NewRouter()
	r.Use(csrfMw)
	r.Use(umw.SetUser)

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

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	//r.Get("/users/me", usersC.CurrentUser)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// start the server
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
