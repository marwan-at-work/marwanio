package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes takes a mux and registers all the routes callbacks within this package
func RegisterRoutes(m *mux.Router, tok string) {
	if tok != "" {
		go runVanityUpdater(tok)
	}

	m.HandleFunc("/", home)

	// This shoudl be part of the /public handler, but Chrome will not work with source maps if
	// the imported path is /public/frontend.js :/
	m.HandleFunc("/frontend.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/frontend.js")
	})

	m.HandleFunc("/frontend.js.map", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/frontend.js.map")
	})

	m.HandleFunc("/favicon.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/favicon.png")
	})

	m.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	m.HandleFunc("/api/blog", blogHandler)

	m.NotFoundHandler = http.HandlerFunc(notFoundOrVanity)
}
