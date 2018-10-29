package router

import (
	"html/template"
	"net/http"
	"strings"
)

var vanity = template.Must(template.New("vanity").Parse(vanityTemplate))

func notFoundOrVanity(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")[0]

	if !pkgExists(path) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Cache-Control", "public")
	vanity.Execute(w, path)
}

func pkgExists(pkg string) bool {
	resp, err := http.Get("https://api.github.com/repos/marwan-at-work/" + pkg)
	if err != nil {
		return false
	}
	resp.Body.Close()
	// if rate limited, go with ok. TODO: use auth token or keep in mem map of my repos.
	return resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusForbidden
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
