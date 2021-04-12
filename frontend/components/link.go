package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Link renders a footer link with a dispatcher for onClick
type Link struct {
	vecty.Core
	Name string
	Link string
}

// Render renders Link
func (fl *Link) Render() vecty.ComponentOrHTML {
	if fl.Link != "" {
		return elem.Anchor(
			vecty.Markup(
				vecty.Class("link", "footer-link"),
				prop.Href(fl.Link),
			),
			vecty.Text(fl.Name),
		)
	}

	return elem.Anchor(
		vecty.Markup(vecty.Class("link")),
		vecty.Text(fl.Name),
	)
}
