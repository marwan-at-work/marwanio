package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Link renders a footer link with a dispatcher for onClick
type Link struct {
	vecty.Core
	Name string
	Link string
}

// Render renders Link
func (fl *Link) Render() *vecty.HTML {
	if fl.Link != "" {
		return elem.Anchor(
			prop.Class("footer-link"),
			vecty.Text(fl.Name),
			prop.Href(fl.Link),
		)
	}

	return elem.Anchor(
		prop.Class("link"),
		vecty.Text(fl.Name),
	)
}
