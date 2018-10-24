package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marwan-at-work/sourcemapper"
	"github.com/marwan-at-work/marwanio/router"
)

func main() {
	h := sourcemapper.NewHandler(getMux())

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}

func getMux() *http.ServeMux {
	var mux http.ServeMux
	router.RegisterRoutes(&mux)

	return &mux
}
