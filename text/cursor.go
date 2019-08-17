package text

import (
	"fmt"
)

// Cursor helps move the cursor on the console in multiple directions.
type Cursor string

const (
	// CursorDown helps move the Cursor Down X lines
	CursorDown Cursor = "B"

	// CursorHide turns the Cursor invisible
	CursorHide Cursor = "?25l"

	// CursorLeft helps move the Cursor Left X characters
	CursorLeft Cursor = "D"

	// CursorRight helps move the Cursor Right X characters
	CursorRight Cursor = "C"

	// CursorShow turns the Cursor visible
	CursorShow Cursor = "?25h"

	// CursorUp helps move the Cursor Up X lines
	CursorUp Cursor = "A"

	// EraseLine helps erase all characters to the Right of the Cursor in the
	// current line
	EraseLine Cursor = "K"
)

// Sprint prints the Escape Sequence to move the Cursor once.
func (c Cursor) Sprint() string {
	return fmt.Sprintf("%s%s", EscapeStart, c)
}

// Sprintn prints the Escape Sequence to move the Cursor "n" times.
func (c Cursor) Sprintn(n int) string {
	switch c {
	case EraseLine, CursorHide, CursorShow:
		return c.Sprint()
	}
	return fmt.Sprintf("%s%d%s", EscapeStart, n, c)
}
