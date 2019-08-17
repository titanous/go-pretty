package text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleCursor_Sprint() {
	fmt.Printf("CursorDown : %#v\n", CursorDown.Sprint())
	fmt.Printf("CursorLeft : %#v\n", CursorLeft.Sprint())
	fmt.Printf("CursorHide : %#v\n", CursorHide.Sprint())
	fmt.Printf("CursorRight: %#v\n", CursorRight.Sprint())
	fmt.Printf("CursorShow : %#v\n", CursorShow.Sprint())
	fmt.Printf("CursorUp   : %#v\n", CursorUp.Sprint())
	fmt.Printf("EraseLine  : %#v\n", EraseLine.Sprint())

	// Output: CursorDown : "\x1b[B"
	// CursorLeft : "\x1b[D"
	// CursorHide : "\x1b[?25l"
	// CursorRight: "\x1b[C"
	// CursorShow : "\x1b[?25h"
	// CursorUp   : "\x1b[A"
	// EraseLine  : "\x1b[K"
}

func TestCursor_Sprint(t *testing.T) {
	assert.Equal(t, "\x1b[B", CursorDown.Sprint())
	assert.Equal(t, "\x1b[D", CursorLeft.Sprint())
	assert.Equal(t, "\x1b[?25l", CursorHide.Sprint())
	assert.Equal(t, "\x1b[C", CursorRight.Sprint())
	assert.Equal(t, "\x1b[?25h", CursorShow.Sprint())
	assert.Equal(t, "\x1b[A", CursorUp.Sprint())
	assert.Equal(t, "\x1b[K", EraseLine.Sprint())
}

func ExampleCursor_Sprintn() {
	fmt.Printf("CursorDown : %#v\n", CursorDown.Sprintn(5))
	fmt.Printf("CursorLeft : %#v\n", CursorLeft.Sprintn(5))
	fmt.Printf("CursorHide : %#v\n", CursorHide.Sprintn(5))
	fmt.Printf("CursorRight: %#v\n", CursorRight.Sprintn(5))
	fmt.Printf("CursorShow : %#v\n", CursorShow.Sprintn(5))
	fmt.Printf("CursorUp   : %#v\n", CursorUp.Sprintn(5))
	fmt.Printf("EraseLine  : %#v\n", EraseLine.Sprintn(5))

	// Output: CursorDown : "\x1b[5B"
	// CursorLeft : "\x1b[5D"
	// CursorHide : "\x1b[?25l"
	// CursorRight: "\x1b[5C"
	// CursorShow : "\x1b[?25h"
	// CursorUp   : "\x1b[5A"
	// EraseLine  : "\x1b[K"
}

func TestCursor_Sprintn(t *testing.T) {
	assert.Equal(t, "\x1b[5B", CursorDown.Sprintn(5))
	assert.Equal(t, "\x1b[5D", CursorLeft.Sprintn(5))
	assert.Equal(t, "\x1b[?25l", CursorHide.Sprintn(5))
	assert.Equal(t, "\x1b[5C", CursorRight.Sprintn(5))
	assert.Equal(t, "\x1b[?25h", CursorShow.Sprintn(5))
	assert.Equal(t, "\x1b[5A", CursorUp.Sprintn(5))
	assert.Equal(t, "\x1b[K", EraseLine.Sprintn(5))
}
