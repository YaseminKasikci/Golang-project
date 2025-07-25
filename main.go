package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=uft-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "Their was an error parsing template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Execute template: %v", err)
		http.Error(w, "Their was an error Execute template", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=uft-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">jon@calhun.io</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contente-type", "text/html;charset=utf-8")
	fmt.Fprint(w, `<H1> FAQ PAGE  </H1> 
		<ul>
			<li>
				<b>IS there a free version</b>
				  Yes! We offer a free trial for 30 days on any paid plans.
			</li>
		</ul>
		<ul>
			<li>
			<b>What are your support hours?</b>
    			We have support staff answering emails 24/7, though response
    			times may be a bit slower on weekends.
			</li>
		</ul>
		<ul>
			<li>
				  <b>How do I contact support?</b>
				  Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
			</li>
		</ul>
		`)
}

func addUrlPath(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	w.Write([]byte(fmt.Sprintf("hi there %v, you just added the following path", userID)))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {

		r.Get("/", homeHandler)
		r.Get("/{id}", addUrlPath)
	})
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
