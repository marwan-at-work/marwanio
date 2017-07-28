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
	body := &components.Body{}
	vecty.RenderBody(body)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
