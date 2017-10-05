/*
Package sourcemapper exposes an `http.Handler` that serves any `.go`
files that exist in your GOPATH or GOROOT.  The purpose of this package
is to implement `.go` source maps for GopherJS by using your own custom
server without having to run `gopherjs serve`.

Example

	package main

	import (
		"net/http"

		"path/to/sourcemapper"
	)

	func main() {
		http.Handle("/one", myHandler)
		http.Handle("/two", myOtherHandler)

		http.ListenAndServe(port, sourcemapper.NewHandler(http.DefaultServeMux))
	}

If you pass nil to sourcemapper.NewHandler, http.DefaultServeMux is used by default.
However, you can use this to pass your own router.
*/
package sourcemapper

import (
	"fmt"
	"go/build"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type handler struct {
	givenHandler http.Handler
}

// NewHandler wraps the http.Handler to see if it gets any reguests for .go files. If true, then
// your handler will not be called. If not, it will forward the request to your handler.
// If .go is at the end of a URL string, then it will try to find that file for you
// in your system and serve it back.
func NewHandler(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	return &handler{h}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if !strings.HasSuffix(url, ".go") {
		h.givenHandler.ServeHTTP(w, r)
		return
	}

	importPath, fileName := filepath.Split(url)

	if strings.HasPrefix(importPath, "/") {
		importPath = importPath[1:]
	}

	pkg, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fullPath := filepath.Join(pkg.Dir, fileName)
	f, err := os.Open(fullPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not open %v: %v", fileName, err)
		return
	}

	io.Copy(w, f)
	f.Close()
}
