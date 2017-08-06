package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/marwan-at-work/vecty-router"
)

// Body renders the <body>  tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() *vecty.HTML {
	return elem.Body(
		router.NewRoute("/", &MainView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/blog", &BlogView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/blog/{id}", &PostView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/talks", &Talks{}, router.NewRouteOpts{ExactMatch: true}),
	)
}
