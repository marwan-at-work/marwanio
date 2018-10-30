package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"marwan.io/marwanio/blog"
)

func blogHandler(w http.ResponseWriter, r *http.Request) {
	bps, err := getBlogPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bps)
}

type blogPosts []blog.Post

// Len returns length of collection
func (bps blogPosts) Len() int {
	return len(bps)
}

// Less returns the earliest time.
func (bps blogPosts) Less(i int, j int) bool {
	first, second := bps[i], bps[j]

	if first.CreatedAt.After(second.CreatedAt) {
		return true
	}

	return false
}

// Swap swaps two elements
func (bps blogPosts) Swap(i int, j int) {
	bps[i], bps[j] = bps[j], bps[i]
}

func getBlogPosts() (bps []blog.Post, err error) {
	files, err := ioutil.ReadDir("./blog/posts")
	if err != nil {
		return
	}

	bps = []blog.Post{}

	for _, file := range files {
		postID := strings.Split(file.Name(), ".")[0]
		bts, err := ioutil.ReadFile("./blog/posts/" + file.Name())
		if err != nil {
			return nil, err
		}
		link := fmt.Sprintf("/blog/%v", postID)

		lines := bytes.Split(bts, []byte("\n"))
		firstLine := string(lines[0])
		title := strings.Replace(
			firstLine,
			"# ",
			"",
			-1,
		)

		secondLine := string(lines[1])
		createdAt, err := time.Parse("2006-01-02", secondLine)
		if err != nil {
			return nil, err
		}

		p := blog.Post{
			ID:        postID,
			Title:     string(title),
			Markdown:  bts,
			Link:      link,
			CreatedAt: createdAt,
		}
		bps = append(bps, p)
	}

	sort.Sort(blogPosts(bps))

	return
}
