package router

import (
	"html/template"
	"net/http"
)

var vt = template.Must(template.New("vanity").Parse(vt2))

func vanityHandler(w http.ResponseWriter, r *http.Request) {
	pkg, ok := getPkg(r.URL.Path)
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte("Package not found\n"))
	}
	vt.Execute(w, pkg)
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
