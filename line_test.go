package goline_test

import (
	"goline"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFromMultilineString(t *testing.T) {
	text := "aa\nbb\ncc\n"

	expected := []goline.Line{
		"aa",
		"bb",
		"cc",
		"",
	}

	result := goline.FromMultilineString(text)

	assert.Equal(t, expected, result)
}

func TestParseFromWindowsMultilineString(t *testing.T) {
	text := "aa\r\nbb\r\ncc\r\n"

	expected := []goline.Line{
		"aa",
		"bb",
		"cc",
		"",
	}

	result := goline.FromMultilineString(text)

	assert.Equal(t, expected, result)
}

func TestParseFromMixedMultilineString(t *testing.T) {
	text := "aa\r\nbb\ncc\r\n"

	expected := []goline.Line{
		"aa",
		"bb",
		"cc",
		"",
	}

	result := goline.FromMultilineString(text)

	assert.Equal(t, expected, result)
}

func TestTrim(t *testing.T) {
	l := goline.Line("  hello      ")

	assert.Equal(t, goline.Line("hello"), l.Trim())
}

func TestCapture(t *testing.T) {
	l := goline.Line("# Title")

	cap, ok := l.Capture("# (.*)", 1)

	assert.True(t, ok)
	assert.Equal(t, "Title", cap)
}

func TestOOBCapture(t *testing.T) {
	l := goline.Line("# Title")

	cap, ok := l.Capture("# (.*)", 2)

	assert.False(t, ok)
	assert.Equal(t, "", cap)
}

func TestUnmatchedCapture(t *testing.T) {
	l := goline.Line("# Title")

	cap, ok := l.Capture("## (.*)", 1)

	assert.False(t, ok)
	assert.Equal(t, "", cap)
}

func TestLookLike(t *testing.T) {
	l := goline.Line("# Title")

	ok := l.LookLike("# (.*)")

	assert.True(t, ok)
}

func TestUnmatchedLookLike(t *testing.T) {
	l := goline.Line("# Title")

	ok := l.LookLike("^## (.*)")

	assert.False(t, ok)
}
