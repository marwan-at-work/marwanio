package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	goMode := os.Getenv("GO_MODE")
	validateGoMode(goMode)
	srv := getServer(goMode)

	fmt.Println("listening on port", srv.Addr)
	if goMode == development || goMode == gae {
		log.Fatal(srv.ListenAndServe())
	}

	log.Fatal(srv.ListenAndServeTLS("", ""))
}

func validateGoMode(goMode string) {
	if goMode != development && goMode != production && goMode != "gae" {
		log.Fatalf("incorrect GO_MODE %v - must be production, development, or gae.\n", goMode)
	}
}
