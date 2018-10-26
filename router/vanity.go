package router

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var vt = template.Must(template.New("vanity").Parse(vt2))

func vanityHandler(w http.ResponseWriter, r *http.Request) {
	pkg := mux.Vars(r)["pkg"]
	vt.Execute(w, pkg)
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
