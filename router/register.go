package router

import "net/http"

// RegisterRoutes takes a mux and registers all the routes callbacks within this package
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", home)

	mux.HandleFunc("/public/css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/main.css")
	})

	mux.HandleFunc("/public/frontend.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/frontend.js")
	})

	mux.HandleFunc("/public/frontend.js.map", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/frontend.js.map")
	})

	mux.HandleFunc("/public/normalize.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/normalize.css")
	})

	mux.HandleFunc("/public/highlight.pack.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/highlight.pack.js")
	})

	mux.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./resume.pdf")
	})

	mux.HandleFunc("/api/blog", blogHandler)
}
