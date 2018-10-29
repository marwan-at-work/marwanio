package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/marwan-at-work/marwanio/router"
	"github.com/marwan-at-work/sourcemapper"
)

func main() {
	h := gziphandler.GzipHandler(
		sourcemapper.NewHandler(
			getMux(),
		),
	)

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}

func getMux() http.Handler {
	r := mux.NewRouter()
	router.RegisterRoutes(r)
	return r
}
