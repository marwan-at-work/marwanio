package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/marwan-at-work/marwanio/router"
	"golang.org/x/crypto/acme/autocert"
)

const (
	production  = "production"
	development = "development"
)

func getServer(goMode string) *http.Server {
	port := resolvePort(goMode)
	mux := getMux()
	srv := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  2 * time.Minute,
		Handler:      mux,
	}

	if goMode == production {
		addTLS(srv)
		go runRedirectServer()
	}

	return srv
}

func runRedirectServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		newURI := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, newURI, http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

func addTLS(srv *http.Server) {
	hostPolicy := func(ctx context.Context, host string) error {
		fmt.Println("host policy:", host)
		allowedHost := "www.marwan.io"
		if host != allowedHost {
			return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
		}

		return nil
	}

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: hostPolicy,
		Cache:      autocert.DirCache("/miocerts"),
	}

	srv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
}

func getMux() *http.ServeMux {
	mux := &http.ServeMux{}
	router.RegisterRoutes(mux)

	return mux
}

// figures out whether to serve on 3000 for development, 80 for no certs, 443 for let's encrypt
func resolvePort(goMode string) string {
	if goMode == development {
		return ":3000"
	}

	return ":443"
}
