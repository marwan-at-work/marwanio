package components

import (
	fmt "github.com/cathalgarvey/fmtless"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Talks renders index of talks I've given
type Talks struct {
	vecty.Core
}

// Render renders Talks component
func (tlks *Talks) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("blog-container")),
		elem.Heading1(
			vecty.Markup(vecty.Class("blog-header")),
			vecty.Text("Talks"),
			getTalk(
				"Handling Go Errors (At The New York Times)",
				"https://slides.com/marwansameer/handling-go-errors#/",
				"2018, Sep 28",
				"GopherCon Brazil",
				"https://2018.gopherconbr.org/en/",
			),
			getTalk(
				"Migrating The Go Community",
				"https://talks.godoc.org/github.com/marwan-at-work/presentations/gophercon/talk.slide#1",
				"2018, Aug 27",
				"GopherCon",
				"https://gophercon.com",
			),
			getTalk(
				"The Go Download Protocol",
				"https://talks.godoc.org/github.com/marwan-at-work/presentations/googlemeetup/talk.slide#1",
				"2018, Aug 23",
				"GolangNYC",
				"https://www.meetup.com/golanguagenewyork/events/253273146/",
			),
			getTalk(
				"Build Your Own Go CI Server",
				"/public/bowery-golang.pdf",
				"2017, Aug 3",
				"Bowery Golang",
				"https://www.meetup.com/Bowery-Go/events/241363507",
			),
		),
	)
}

func getTalk(title, link, date, eventTitle, eventLink string) *vecty.HTML {
	return elem.Paragraph(
		vecty.Markup(vecty.Class("post-title")),
		elem.Anchor(
			vecty.Markup(
				vecty.Class("post-title-text"),
				prop.Href(link),
			),
			vecty.Text(title),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class("post-title-date"),
				vecty.UnsafeHTML(
					fmt.Sprintf(
						"%v - <a href=\"%v\">%v</a>",
						date,
						eventLink,
						eventTitle,
					),
				),
			),
		),
	)
}
