package router

import (
	"net/http"
)

// RegisterRoutes takes a mux and registers all the routes callbacks within this package
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", home)

	mux.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./resume.pdf")
	})

	// This shoudl be part of the /public handler, but Chrome will not work with source maps if
	// the imported path is /public/frontend.js :/
	mux.HandleFunc("/frontend.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/frontend.js")
	})

	mux.HandleFunc("/frontend.js.map", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/frontend.js.map")
	})

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("/api/blog", blogHandler)
}
