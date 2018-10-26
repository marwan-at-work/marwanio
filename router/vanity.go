package router

import (
	"html/template"
	"net/http"
)

var vt = template.Must(template.New("vanity").Parse(vt2))

func vanityHandler(w http.ResponseWriter, r *http.Request) {
	pkg, ok := getPkg(r.URL.Path)
	if !ok || !pkgExists(pkg) {
		w.WriteHeader(404)
		w.Write([]byte("Package not found\n"))
	}
	vt.Execute(w, pkg)
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

func getPkg(s string) (string, bool) {
	l := len("/pkg/")
	if len(s) <= l {
		return "", false
	}

	return s[l:], true
}

var vt2 = `<html>
  <head>
    <meta name="go-import" content="marwan.io/pkg/{{ . }} git https://github.com/marwan-at-work/{{ . }}">
  </head>
  <body>
    Install: go get -u marwan.io/pkg/{{ . }} <br>
    <a href="http://godoc.org/marwan.io/pkg/{{ . }}">Documentation</a><br>
    <a href="https://github.com/marwan-at-work/{{ . }}">Source</a>
  </body>
</html>
`
