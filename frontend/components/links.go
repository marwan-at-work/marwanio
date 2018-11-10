package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"marwan.io/vecty-router"
)

// Links renders the Work/Blog/Resume links
type Links struct {
	vecty.Core
}

// Render renders Footer
func (l *Links) Render() vecty.ComponentOrHTML {
	return elem.Paragraph(
		vecty.Markup(prop.ID("links-container")),
		&Link{
			Name: "GITHUB",
			Link: "https://www.github.com/marwan-at-work",
		},
		router.Link(
			"/blog",
			"BLOG",
			router.LinkOptions{
				Class: "link",
			},
		),
		// &Link{
		// 	Name: "RESUME",
		// 	Link: "/resume",
		// },
		router.Link(
			"/talks",
			"TALKS",
			router.LinkOptions{
				Class: "link",
			},
		),
	)
}
