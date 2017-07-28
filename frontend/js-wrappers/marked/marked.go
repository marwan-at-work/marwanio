// Package marked wraps around the marked node module whih provides markdown parsing and highlighting ufnctionalities.
package marked

import (
	"github.com/gopherjs/gopherjs/js"
)

func init() {
	js.Global.Get("marked").Call("setOptions", map[string]func(string) string{
		"highlight": func(code string) string {
			return js.Global.Get("hljs").Call("highlightAuto", code).Get("value").String()
		},
	})
}

// Marked wraps the marked library. It takes markdown as input and returns HTML.
func Marked(bts []byte) string {
	return js.Global.Call("marked", string(bts)).String()
}
