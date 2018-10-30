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
	"github.com/marwan-at-work/marwanio/router"
	"github.com/marwan-at-work/marwanio/security"
	"github.com/marwan-at-work/sourcemapper"
)

func main() {
	githubToken := getToken()
	h := gziphandler.GzipHandler(
		sourcemapper.NewHandler(
			getMux(githubToken),
		),
	)

	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", h))
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
