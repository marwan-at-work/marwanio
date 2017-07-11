package fmt

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpec(t *testing.T) {
	s, o := getSpec([]byte("%d"))
	assert.True(t, o)
	assert.Equal(t, s, "%d")

	s, o = getSpec([]byte("%df"))
	assert.True(t, o)
	assert.Equal(t, s, "%d")

	s, o = getSpec([]byte("%f"))
	assert.True(t, o)
	assert.Equal(t, s, "%f")

	s, o = getSpec([]byte("%n"))
	assert.False(t, o)
	assert.Equal(t, s, "")

	s, o = getSpec([]byte("ff"))
	assert.False(t, o)
	assert.Equal(t, s, "")

	s, o = getSpec([]byte("%+f"))
	assert.True(t, o)
	assert.Equal(t, s, "%+f")
}

func TestSplitSpecs(t *testing.T) {
	// splitFmtSpecs(fmts string) []sprintMatch {
	specs := splitFmtSpecs("This %s that %d these %f those %g")
	expected := []sprintMatch{
		sprintMatch{"This ", "%s"},
		sprintMatch{" that ", "%d"},
		sprintMatch{" these ", "%f"},
		sprintMatch{" those ", "%g"},
	}
	assert.EqualValues(t, expected, specs)
}

func TestSprintf(t *testing.T) {
	filled := Sprintf("This: '%s' is stringier than: %q ", "\"string\"", "\"string\"")
	assert.Equal(t, `This: '"string"' is stringier than: "\"string\"" `, filled)

	filled = Sprintf("This: '%s' is stringier than: %q", "\"string\"", "\"string\"")
	assert.Equal(t, `This: '"string"' is stringier than: "\"string\""`, filled)

	filled = Sprintf("There are %d ways to kill someone who rounds pi to %f", 3, 3.1)
	assert.Equal(t, "There are 3 ways to kill someone who rounds pi to 3.1", filled)

	filled = Sprintf("%U + %U != %U", []rune("a")[0], []rune("í")[0], []rune("쎭")[0])
	assert.Equal(t, "U+0061 + U+00ED != U+C3AD", filled)

	filled = Sprintf("%v == %s", Errorf("error %v", 1), fmt.Errorf("error %d", 2))
	assert.Equal(t, "error 1 == error 2", filled)

	filled = Sprintf("%X", []byte{1, 2, 3, 4})
	assert.Equal(t, fmt.Sprintf("%X", []byte{1, 2, 3, 4}), filled)

	filled = Sprintf("%X", []byte{1, 2, 4, 8, 16, 32, 64, 128, 255})
	assert.Equal(t, fmt.Sprintf("%X", []byte{1, 2, 4, 8, 16, 32, 64, 128, 255}), filled)

	filled = Sprintf("%x", []byte{1, 2, 4, 8, 16, 32, 64, 128, 255})
	assert.Equal(t, fmt.Sprintf("%x", []byte{1, 2, 4, 8, 16, 32, 64, 128, 255}), filled)

	// Similar to a problematic error in json decode.go
	errfilled := Errorf("failed to unmarshal %q into %v", "1", reflect.ValueOf("").Type())
	assert.Equal(t, errors.New("failed to unmarshal \"1\" into string"), errfilled)
}

func TestSprintCodepoint(t *testing.T) {
	// fmtUEscape(i.(rune))
	uesc := fmtUEscape('\x12')
	assert.Equal(t, "U+0012", uesc)

	uesc = fmtUEscape(18)
	assert.Equal(t, "U+0012", uesc)

	uesc = fmtUEscape([]rune("í")[0])
	assert.Equal(t, "U+00ED", uesc)

	uesc = fmtUEscape([]rune("쎭")[0])
	assert.Equal(t, "U+C3AD", uesc)
}
