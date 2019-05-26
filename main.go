package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/hashicorp/hcl"
	"github.com/marwan-at-work/sourcemapper"
	"marwan.io/marwanio/router"
	"marwan.io/marwanio/security"
)

func main() {
	githubToken := getToken()
	h := gziphandler.GzipHandler(
		sourcemapper.NewHandler(
			getMux(githubToken),
		),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, h))
}

func getToken() string {
	bts, err := ioutil.ReadFile("./config.hcl")
	if os.IsNotExist(err) {
		fmt.Println("vanity imports disabled")
		return ""
	} else if err != nil {
		log.Fatal(err)
	}
	var cfg security.GCPConfig
	err = hcl.Unmarshal(bts, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	tok, err := security.GithubToken(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return tok
}

func getMux(tok string) http.Handler {
	r := mux.NewRouter()
	router.RegisterRoutes(r, tok)
	return r
}
