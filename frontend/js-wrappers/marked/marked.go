// Package marked wraps around the marked node module whih provides markdown parsing and highlighting ufnctionalities.
package marked

import (
	"syscall/js"
)

func init() {
	js.Global().Get("marked").Call("setOptions", map[string]interface{}{
		"highlight": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return js.Global().Get("hljs").Call("highlightAuto", args[0].String()).Get("value").String()
		}),
	})
}

// Marked wraps the marked library. It takes markdown as input and returns HTML.
func Marked(bts []byte) string {
	return js.Global().Call("marked", string(bts)).String()
}
