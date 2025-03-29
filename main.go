package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/decedis/lenslocked/controllers"
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
	tpl := views.Must(views.Parse(filepath.Join("templates", "home.go.html")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.go.html")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.go.html")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
