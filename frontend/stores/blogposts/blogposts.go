package blogposts

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/marwan-at-work/marwanio/blog"
)

var posts = []blog.Post{}

// ErrNotFound is used when searching for a blog post is unfruitful.
var ErrNotFound = errors.New("blog post was not found")

// Fetch will populate all blog posts from the server.
func Fetch() error {
	resp, err := http.Get("/api/blog")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&posts)
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
