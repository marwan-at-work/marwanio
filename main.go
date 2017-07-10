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
	if goMode == development {
		log.Fatal(srv.ListenAndServe())
	}

	log.Fatal(srv.ListenAndServeTLS("", ""))
}

func validateGoMode(goMode string) {
	if goMode != "development" && goMode != "production" {
		log.Fatalf("incorrect GO_MODE %v - must be production or development\n", goMode)
	}
}
