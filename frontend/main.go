package main

import (
	"github.com/gopherjs/vecty"
	"github.com/marwan-at-work/marwanio/frontend/components"
	"github.com/marwan-at-work/marwanio/frontend/stores/blogposts"
)

func main() {
	must(blogposts.Fetch())
	vecty.SetTitle("Marwan - Software Engineer")
	vecty.AddStylesheet("/public/css")
	vecty.RenderBody(&components.Body{})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
