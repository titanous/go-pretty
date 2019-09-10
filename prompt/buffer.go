package prompt

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/alecthomas/chroma/quick"
	"github.com/jedib0t/go-pretty/text"
)

// cursorPosition contains the current cursor position in a 2d-wall-of-text; the
// values are 0-indexed to keep it simple to manipulate the wall of text
type cursorPosition struct {
	Line int
	Col  int
}

// buffer helps store the user input, track the cursor position, and help
// manipulate the user input with adding/removing strings from it
type buffer struct {
	AutoCompleter AutoCompleter
	Style         *Style

	clearString            strings.Builder
	done                   bool
	lines                  []string
	linesRendered          string
	mutex                  sync.Mutex
	numRenders             int
	position               cursorPosition
	positionRendered       cursorPosition
	firstLinePrefix        string
	syntaxHighlighterCache map[string]string
}

// newBuffer returns a buffer object with sane defaults
func newBuffer(autoCompleter AutoCompleter, style *Style) *buffer {
	b := &buffer{
		AutoCompleter: autoCompleter,
		Style:         style,

		syntaxHighlighterCache: make(map[string]string),
	}
	b.Clear()
	return b
}

// Clear clears the buffer
func (b *buffer) Clear() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.lines = []string{""}
	b.linesRendered = fmt.Sprint(time.Now().Format(time.RFC3339Nano))
	b.position = cursorPosition{Line: 0, Col: 0}
}

// DeleteBackward deletes n runes backwards
func (b *buffer) DeleteBackward(n int, locked ...bool) {
	if len(locked) == 0 {
		b.mutex.Lock()
		defer b.mutex.Unlock()
	}

	// if asked to delete till beginning, just set N to the max value possible
	if n == -1 {
		n = len(strings.Join(b.lines, "\n"))
	}

	// delete backward rune by rune
	for ; n > 0; n-- {
		if b.position.Col == 0 {
			if b.position.Line > 0 {
				prevLine, line := b.getLine(b.position.Line-1), b.getLine(b.position.Line)
				var lines []string
				lines = append(lines, b.lines[:b.position.Line-1]...)
				lines = append(lines, prevLine+line)
				if b.position.Line < len(b.lines)-1 {
					lines = append(lines, b.lines[b.position.Line+1:]...)
				}

				b.lines = lines
				b.position.Line--
				b.position.Col = len(prevLine)
			}
		} else {
			line := b.getCurrentLine()
			b.lines[b.position.Line] = line[:b.position.Col-1] + line[b.position.Col:]
			b.position.Col--
		}
	}
}

// DeleteForward deletes n runes forwards
func (b *buffer) DeleteForward(n int, locked ...bool) {
	if len(locked) == 0 {
		b.mutex.Lock()
		defer b.mutex.Unlock()
	}

	// if asked to delete till end, just set N to the max value possible
	if n == -1 {
		n = len(strings.Join(b.lines, "\n"))
	}

	// delete forward rune by rune
	for ; n > 0; n-- {
		line := b.getCurrentLine()
		if b.position.Col == len(line) {
			if b.position.Line == len(b.lines)-1 {
				return
			}
			line += b.getLine(b.position.Line + 1)

			var lines []string
			lines = append(lines, b.lines[:b.position.Line]...)
			lines = append(lines, line)
			if b.position.Line < len(b.lines)-2 {
				lines = append(lines, b.lines[b.position.Line+2:]...)
			}

			b.lines = lines
		} else if b.position.Col > 0 {
			b.lines[b.position.Line] = line[:b.position.Col] + line[b.position.Col+1:]
		} else {
			b.lines[b.position.Line] = line[b.position.Col+1:]
		}
	}
}

// DeleteWordBackward deletes the previous word
func (b *buffer) DeleteWordBackward() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// already on the first column? delete one char backwards...
	if b.position.Col == 0 {
		b.DeleteBackward(1, true)
		return
	}

	// delete till beginning of previous word
	foundWord := false
	line := b.getCurrentLine()
	for idx := b.position.Col - 1; idx >= 0; idx-- {
		isPartOfWord := b.isPartOfWord(line[idx])
		if !isPartOfWord && foundWord {
			b.lines[b.position.Line] = line[:idx] + line[b.position.Col:]
			b.position.Col = idx
			return
		}
		if isPartOfWord {
			foundWord = true
		}
	}
	b.lines[b.position.Line] = line[b.position.Col:]
	b.position.Col = 0
}

// DeleteWordForward deletes the next word
func (b *buffer) DeleteWordForward() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// already on the last column? delete one char forwards...
	line := b.getCurrentLine()
	if b.position.Col == len(line) {
		b.DeleteForward(1, true)
		return
	}

	// delete till beginning of previous word
	foundWord, foundNonWord := false, false
	for idx := b.position.Col; idx < len(line); idx++ {
		isPartOfWord := b.isPartOfWord(line[idx])
		if !isPartOfWord {
			foundNonWord = true
		}
		if isPartOfWord && foundWord && foundNonWord {
			b.lines[b.position.Line] = line[:b.position.Col] + line[idx:]
			return
		}
		if isPartOfWord {
			foundWord = true
		}
	}
	b.lines[b.position.Line] = line[:b.position.Col]
}

// HasChanges returns true if Render() will return something else on the next
// call to it.
func (b *buffer) HasChanges() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.linesRendered != fmt.Sprintf("%v", b.lines) ||
		b.positionRendered != b.position ||
		b.Style.Timestamp.Enabled
}

// Insert inserts the string at the current cursor position
func (b *buffer) Insert(r rune) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if r == '\n' {
		line := b.getCurrentLine()

		var lines []string
		lines = append(lines, b.lines[:b.position.Line]...)
		lines = append(lines, line[:b.position.Col])
		if b.position.Col < len(line) { // cursor somewhere before the end
			lines = append(lines, line[b.position.Col:])
		} else {
			lines = append(lines, "")
		}
		lines = append(lines, b.lines[b.position.Line+1:]...)

		b.lines = lines
		b.position.Line++
		b.position.Col = 0
	} else {
		rStr := fmt.Sprintf("%c", r)
		if b.Style.Tab != "\t" && rStr == "\t" {
			rStr = b.Style.Tab
		}

		line := b.getCurrentLine()
		line = line[:b.position.Col] + rStr + line[b.position.Col:]

		b.lines[b.position.Line] = line
		b.position.Col += len(rStr)
	}
}

func (b *buffer) IsDone() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.done
}

// Length returns the current input length
func (b *buffer) Length() int {
	return len(b.String())
}

// MakeWordCapitalCase converts the current word to Capital case
func (b *buffer) MakeWordCapitalCase() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	line := b.getCurrentLine()
	word, idxWordStart, idxWordEnd := b.getCurrentWord(line)
	if word == "" || idxWordStart == -1 || idxWordEnd == -1 {
		return
	}

	word = strings.ToUpper(fmt.Sprintf("%c", word[0])) + word[1:]
	b.lines[b.position.Line] = line[:idxWordStart] + word + line[idxWordEnd:]
	b.MoveWordRight(true)
}

// MakeWordCapitalCase converts the current word to Lower case
func (b *buffer) MakeWordLowerCase() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	line := b.getCurrentLine()
	word, idxWordStart, idxWordEnd := b.getCurrentWord(line)
	if word == "" || idxWordStart == -1 || idxWordEnd == -1 {
		return
	}

	b.lines[b.position.Line] = line[:idxWordStart] + strings.ToLower(word) + line[idxWordEnd:]
	b.MoveWordRight(true)
}

// MakeWordCapitalCase converts the current word to Upper case
func (b *buffer) MakeWordUpperCase() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	line := b.getCurrentLine()
	word, idxWordStart, idxWordEnd := b.getCurrentWord(line)
	if word == "" || idxWordStart == -1 || idxWordEnd == -1 {
		return
	}

	b.lines[b.position.Line] = line[:idxWordStart] + strings.ToUpper(word) + line[idxWordEnd:]
	b.MoveWordRight(true)
}

// MarkAsDone signifies that the user input is done
func (b *buffer) MarkAsDone() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.done = true
}

// MoveDown attempts to move the cursor to the same position in the next line
func (b *buffer) MoveDown(n int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.position.Line += n
	if b.position.Line >= len(b.lines) {
		b.position.Line = len(b.lines) - 1
	}
	line := b.getCurrentLine()
	if b.position.Col > len(line) {
		b.position.Col = len(line)
	}
}

// MoveLeft moves the cursor left n runes
func (b *buffer) MoveLeft(n int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// move to the very beginning
	if n == -1 {
		b.position = cursorPosition{Line: 0, Col: 0}
		return
	}

	// move left until n becomes 0, or beginning of buffer is reached
	for ; n > 0; n-- {
		b.position.Col--
		if b.position.Col < 0 {
			b.position.Line--
			if b.position.Line < 0 {
				b.position.Line = 0
				b.position.Col = 0
				break
			}
			b.position.Col = len(b.getCurrentLine())
		}
	}
}

// MoveLineBegin moves the cursor right to the beginning of the current line
func (b *buffer) MoveLineBegin() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.position.Col = 0
}

// MoveLineEnd moves the cursor right to the end of the current line
func (b *buffer) MoveLineEnd() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.position.Col = len(b.getCurrentLine())
}

// MoveRight moves the cursor right n runes
func (b *buffer) MoveRight(n int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// move to the very end
	if n == -1 {
		b.position.Line = len(b.lines) - 1
		b.position.Col = len(b.getCurrentLine())
		return
	}

	// move right until n becomes 0, or end of buffer is reached
	for ; n > 0; n-- {
		line := b.getCurrentLine()
		b.position.Col++
		if b.position.Col > len(line) {
			if b.position.Line == len(b.lines)-1 {
				b.position.Col = len(line)
				break
			}
			b.position.Line++
			b.position.Col = 0
		}
	}
}

// MoveUp attempts to move the cursor to the same position in the previous line
func (b *buffer) MoveUp(n int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.position.Line -= n
	if b.position.Line < 0 {
		b.position.Line = 0
	}
	line := b.getCurrentLine()
	if b.position.Col > len(line) {
		b.position.Col = len(line)
	}
}

// MoveWordLeft moves the cursor left to the previous word
func (b *buffer) MoveWordLeft(locked ...bool) {
	if len(locked) == 0 {
		b.mutex.Lock()
		defer b.mutex.Unlock()
	}

	// if cursor is on the first column, move it to the previous line
	if b.position.Col == 0 {
		if b.position.Line == 0 {
			return
		}
		b.position.Line--
		b.position.Col = len(b.getCurrentLine())
	}

	// move line by line until previous word is found
	foundWord := false
	idxStartingLine := b.position.Line
	for lineIdx := b.position.Line; lineIdx >= 0; lineIdx-- {
		b.position.Line = lineIdx
		if lineIdx != idxStartingLine {
			b.position.Col = len(b.getCurrentLine())
		}

		line := b.lines[lineIdx]
		for colIdx := b.position.Col - 1; colIdx >= 0; colIdx-- {
			b.position.Col = colIdx
			isPartOfWord := b.isPartOfWord(line[colIdx])
			if foundWord && (!isPartOfWord || colIdx == 0) {
				if !isPartOfWord {
					b.position.Col++
				}
				return
			}
			if isPartOfWord {
				foundWord = true
			}
		}
	}
}

// MoveWordRight moves the cursor right to the next word
func (b *buffer) MoveWordRight(locked ...bool) {
	if len(locked) == 0 {
		b.mutex.Lock()
		defer b.mutex.Unlock()
	}

	// if cursor is on the last column, move to the next line
	foundBreak := false
	idxStartingLine := b.position.Line
	line := b.getCurrentLine()
	if b.position.Col == len(line) {
		// if already on the last line, there is nothing to do
		if b.position.Line == len(b.lines)-1 {
			return
		}
		b.position.Line++
		b.position.Col = 0
		foundBreak = true
	}

	// go line by line until next word is found
	for lineIdx := b.position.Line; lineIdx < len(b.lines); lineIdx++ {
		b.position.Line = lineIdx
		if lineIdx != idxStartingLine {
			b.position.Col = 0
		}

		line := b.lines[lineIdx]
		for colIdx := b.position.Col; colIdx < len(line); colIdx++ {
			b.position.Col = colIdx
			isPartOfWord := b.isPartOfWord(line[b.position.Col])
			if isPartOfWord && foundBreak {
				return
			}
			if !isPartOfWord {
				foundBreak = true
			}
		}
		b.position.Col = len(line)
		foundBreak = true
	}
}

// Render returns the current buffer contents as a string that can be printed.
// The string also contains control sequences to clears any earlier content that
// was returned on a previous call.
func (b *buffer) Render() string {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// update state
	b.positionRendered = b.position
	b.linesRendered = fmt.Sprintf("%v", b.lines)
	b.numRenders++

	// format the input
	inLines := strings.Split(b.stylizeUserInput(), "\n")

	// build the string from the buffer contents
	out := &strings.Builder{}
	// hide the cursor to remove cursor movement from the screen
	out.WriteString(text.CursorHide.Sprint())
	// clear any previous renders from screen
	if b.clearString.Len() > 0 {
		out.WriteString(b.clearString.String())
		b.clearString.Reset()
	}
	// render the user input
	b.renderInput(out, inLines)
	// render auto-completion suggestions
	if b.AutoCompleter != nil {
		b.renderSuggestions(out, inLines)
	}
	// move the cursor to the appropriate location
	b.renderCursor(out, inLines)
	// show the cursor
	out.WriteString(text.CursorShow.Sprint())
	return out.String()
}

func (b *buffer) Set(str string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.Style.Tab != "\t" {
		str = strings.ReplaceAll(str, "\t", b.Style.Tab)
	}
	b.lines = strings.Split(str, "\n")
	b.linesRendered = time.Now().Format(time.RFC3339Nano)
	b.position = cursorPosition{
		Line: len(b.lines) - 1,
		Col:  len(b.lines[len(b.lines)-1]),
	}
}

// String returns the current input from the user.
func (b *buffer) String() string {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return strings.Join(b.lines, "\n")
}

func (b *buffer) SwapCharacterNext() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

}

func (b *buffer) SwapCharacterPrevious() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

}

func (b *buffer) SwapWordNext() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

}

func (b *buffer) SwapWordPrevious() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

}

func (b *buffer) getCurrentLine() string {
	return b.getLine(b.position.Line)
}

func (b *buffer) getCurrentWord(line string) (string, int, int) {
	if len(line) == 0 || b.position.Col >= len(line) || !b.isPartOfWord(line[b.position.Col]) {
		return "", -1, -1
	}

	idxWordStart, idxWordEnd := -1, -1
	for idx := b.position.Col; idx >= 0; idx-- {
		if !b.isPartOfWord(line[idx]) {
			break
		}
		idxWordStart = idx
	}
	for idx := b.position.Col; idx < len(line); idx++ {
		if !b.isPartOfWord(line[idx]) {
			break
		}
		idxWordEnd = idx
	}

	return line[idxWordStart : idxWordEnd+1], idxWordStart, idxWordEnd + 1
}

func (b *buffer) getLine(n int) string {
	return b.lines[n]
}

func (b *buffer) isPartOfWord(r uint8) bool {
	return !(r == ' ' || r == '\n' || r == '\t' || r == ';' || r == '.' || r == ',')
}

func (b *buffer) renderCursor(out *strings.Builder, lines []string) {
	// move the cursor to the very start of the prompt
	out.WriteString(text.CursorUp.Sprintn(len(lines)))

	// move down the cursor to the current line
	if b.position.Line > 0 {
		out.WriteString(text.CursorDown.Sprintn(b.position.Line))
	}

	// move the cursor "right"
	numMovesRight := text.RuneCount(b.firstLinePrefix)
	if b.position.Line > 0 {
		numMovesRight = text.RuneCount(b.Style.NewlineIndent)
	}
	numMovesRight += b.position.Col
	out.WriteString(text.CursorRight.Sprintn(numMovesRight))

	// reset the clear string by moving the cursor to the very bottom of the
	// prompt
	clearString := b.clearString.String()
	b.clearString.Reset()
	b.clearString.WriteString(strings.Repeat("\n", len(lines)-b.position.Line))
	b.clearString.WriteString(clearString)
}

func (b *buffer) renderInput(out *strings.Builder, lines []string) {
	b.firstLinePrefix = b.Style.Timestamp.Generate() + b.Style.Prefix
	for idx, line := range lines {
		if idx == 0 {
			out.WriteString(b.firstLinePrefix)
		} else {
			out.WriteString(b.Style.NewlineIndent)
		}
		out.WriteString(line)
		out.WriteRune('\n')

		// form the string to clear out this line
		b.clearString.WriteString(text.CursorUp.Sprint())
		b.clearString.WriteString(text.EraseLine.Sprint())
	}
}

func (b *buffer) renderSuggestions(out *strings.Builder, lines []string) {

}

func (b *buffer) stylizeUserInput() string {
	in := strings.Join(b.lines, "\n")
	if in == "" {
		return ""
	} else if out, ok := b.syntaxHighlighterCache[in]; ok {
		return out
	}

	outBuilder := &strings.Builder{}
	sh := b.Style.SyntaxHighlighter
	if sh.Enabled {
		if err := quick.Highlight(outBuilder, in, sh.Language, sh.Formatter, sh.Style); err != nil {
			outBuilder.Reset()
			outBuilder.WriteString(in)
		}
	} else {
		outBuilder.WriteString(b.Style.Colors.Sprint(in))
	}
	out := outBuilder.String()
	b.syntaxHighlighterCache[in] = out
	return out
}
