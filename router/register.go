package router

import "net/http"

// RegisterRoutes takes a mux and registers all the routes callbacks within this package
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", home)

	mux.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./resume.pdf")
	})

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("/api/blog", blogHandler)
}
