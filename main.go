package main

import (
	"fmt"
	"net/http"

	"github.com/decedis/lenslocked/controllers"
	"github.com/decedis/lenslocked/templates"
	"github.com/decedis/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // <--<< Logger should come before Recoverer
	r.Use(middleware.Recoverer)

	//http.Handler - interface with the ServeHTTP method
	//http.HandlerFunc - a function that accetpts same args ase ServeHTTP method. Also implements http.Handler.
	homeTpl := views.Must(views.ParseFS(templates.FS, "home.go.html", "tailwind.go.html"))
	r.Get("/", controllers.StaticHandler(homeTpl))

	contactTpl := views.Must(views.ParseFS(templates.FS, "contact.go.html", "tailwind.go.html"))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	faqTpl := views.Must(views.ParseFS(templates.FS, "faq.go.html", "tailwind.go.html"))
	r.Get("/faq", controllers.FAQ(faqTpl))

	servicesTpl := views.Must(views.ParseFS(templates.FS, "services.go.html", "tailwind.go.html"))
	r.Get("/services", controllers.StaticHandler(servicesTpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
