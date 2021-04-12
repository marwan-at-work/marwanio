package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// NameAndTitleView is the main text in the center of the page.
type NameAndTitleView struct {
	vecty.Core
}

// Render renders the name in the middle of the screen.
func (h *NameAndTitleView) Render() vecty.ComponentOrHTML {
	return elem.Heading1(
		vecty.Markup(prop.ID("name-and-title")),
		vecty.Text("Marwan Sulaiman - Software Developer"),
	)
}
