package sourcemapper

import (
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
		return
	}

	io.Copy(w, f)
	f.Close()
}
