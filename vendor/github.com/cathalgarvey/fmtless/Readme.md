# fmtless: All the convenience of `fmt` without the bloat
by Cathal Garvey, ©2016, Released under the GNU AGPLv3 or later

### Why Avoid `fmt`?
The `fmt` library is a super-rich way to present and parse data
in your Go application. I love `fmt`; everyone loves `fmt`!

However, `fmt` is really big, adding a large premium to output
binaries. For straight compilation to static binaries, this isn't
usually a dealbreaker (*cf.* the success of Go overall).

However, in edge-cases, like embedding Go in storage-constrained
devices, or shipping a collection of small apps in Go, or when
[transpiling to JS](https://github.com/gopherjs/gopherjs), the
premium can be really costly. In my own experience, removing `fmt`
from a GopherJS application removed 0.5Mb from the output Javascript,
which can be a huge deal when shipping JS.

Most `fmt` imports seem to use it only for output, or for errors.
So, `fmtless` is a toolkit for rapidly porting software using `fmt`
for output and errors, hopefully only requiring a single-line change:
instead of `import "fmt"`, just do `import "github.com/cathalgarvey/fmtless"`!

`fmtless` also has mirrors of Go's standard libraries with `fmt` replaced
with `fmtless`, to reduce binary size. Try switching over and see if it
makes a difference to your application!

At present the included stdlibs in this repo are lifted directly from
my local install. In the future I plan to have a script pull chosen
libraries directly from the most recent tag/release of Go and converts
them.

### Usage
Right now, it's just replace `fmt` with `github.com/cathalgarvey/fmtless`.
For stdlibs, likewise, it's just "prepend your stdlib imports with
github.com/cathalgarvey/fmtless". Currently supported:

* `encoding/json` (was only using stdlibs for errors)
* `encoding/xml` (was only using stdlibs for errors)
* `net/url` (was only using stdlibs for errors)
