package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my bitchin' site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1> <p>Contact me at <a href=#>sbrcak@live.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ</h1> <p>This is the FAQ...</p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

func main() {
	//http.Handler - interface with the ServeHTTP method
	//http.HandlerFunc - a function that accetpts same args ase ServeHTTP method. Also implements http.Handler.

	var router http.HandlerFunc = pathHandler
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
