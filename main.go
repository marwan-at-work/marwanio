package main

import (
	"bytes"
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yuin/goldmark"
	"marwan.io/muxw"
	"marwan.io/serverctx"
)

//go:embed public
var public embed.FS

var dev = flag.Bool("dev", false, "run in developer mode")

func main() {
	flag.Parse()

	r := muxw.NewRouter()

	var public fs.FS = public
	if *dev {
		public = os.DirFS("public")
	} else {
		var err error
		public, err = fs.Sub(public, "public")
		if err != nil {
			log.Fatal(err)
		}
	}

	r.Get("/", serveFile(public, "index.html"))
	r.Get("/talks", serveFile(public, "talks.html"))
	r.Get("/blog", serveFile(public, "blogs.html"))
	r.Get("/blog/{post}", serveBlogPost(public))
	r.Mount("/public", http.StripPrefix("/public", http.FileServerFS(public)))
	r.SetNotFoundHandler(notFoundVanity)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := serverctx.Run(ctx, &http.Server{
		Addr:    ":3091",
		Handler: r,
	}, 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

func serveFile(public fs.FS, file string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, public, file)
	}
}

func serveBlogPost(public fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := fs.ReadFile(public, "blog/"+r.PathValue("post")+".md")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		var buf bytes.Buffer
		err = goldmark.Convert(file, &buf)
		if err != nil {
			fmt.Println(err)
		}
		post, err := fs.ReadFile(public, "blogpost.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		resp := fmt.Sprintf(string(post), buf.String())
		fmt.Fprintln(w, resp)
	}
}
