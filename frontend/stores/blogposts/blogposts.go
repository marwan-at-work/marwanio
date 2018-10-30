package blogposts

import (
	"errors"

	"github.com/cathalgarvey/fmtless/encoding/json"
	"honnef.co/go/js/xhr"
	"marwan.io/marwanio/blog"
)

var posts = []blog.Post{}

// ErrNotFound is used when searching for a blog post is unfruitful.
var ErrNotFound = errors.New("blog post was not found")

// Fetch will populate all blog posts from the server.
func Fetch() error {
	bts, err := xhr.Send("GET", "/api/blog", nil)
	if err != nil {
		return err
	}

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
