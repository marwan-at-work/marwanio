package main

import (
	"html/template"
	"net/http"
	"strings"
)

var vanity = template.Must(template.New("vanity").Parse(vanityTemplate))

func notFoundVanity(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")[0]

	w.Header().Set("Cache-Control", "public")
	vanity.Execute(w, path)
}

var vanityTemplate = `<html>
  <head>
    <meta name="go-import" content="marwan.io/{{ . }} git https://github.com/marwan-at-work/{{ . }}">
  </head>
  <body>
    Install: go get -u marwan.io/{{ . }} <br>
    <a href="http://godoc.org/marwan.io/{{ . }}">Documentation</a><br>
    <a href="https://github.com/marwan-at-work/{{ . }}">Source</a>
  </body>
</html>
`
