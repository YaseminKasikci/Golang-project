package main

import (
	"fmt"
	"github/yaseminkasikci/lenslocked/controllers"
	"github/yaseminkasikci/lenslocked/views"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	tpl, err := views.Parse(filepath.Join("templates/", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	// -----------------------------

	tpl, err = views.Parse(filepath.Join("templates/", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	// ------------------------------

	tpl, err = views.Parse(filepath.Join("templates/", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	// r.Use(middleware.Logger)

	// r.Route("/", func(r chi.Router) {

	// 	r.Get("/", homeHandler)
	// 	r.Get("/{id}", addUrlPath)
	// })

	// r.Get("/contact", contactHandler)
	// r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
