package blogposts

import (
	"errors"

	"github.com/cathalgarvey/fmtless/encoding/json"
	"marwan.io/marwanio/blog"
	fetch "marwan.io/wasm-fetch"
)

var posts = []blog.Post{}

// ErrNotFound is used when searching for a blog post is unfruitful.
var ErrNotFound = errors.New("blog post was not found")

// Fetch will populate all blog posts from the server.
func Fetch() error {
	resp, err := fetch.Fetch("/api/blog", &fetch.Opts{Method: "GET"})
	if err != nil {
		return err
	}
	bts := resp.Body
	return json.Unmarshal(bts, &posts)
}

// GetAll returns a copy of all blog posts
func GetAll() []blog.Post {
	return posts
}

// GetByID returns a blog post or ErrNotFound
func GetByID(id string) (blog.Post, error) {
	for _, p := range posts {
		if p.ID == id {
			return p, nil
		}
	}

	return blog.Post{}, ErrNotFound
}
