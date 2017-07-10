package router

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
}
