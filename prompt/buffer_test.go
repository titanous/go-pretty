package prompt

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getNewBuffer(t *testing.T) *buffer {
	b := newBuffer(nil, &StyleDefault)
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	return b
}

func TestBuffer_Clear(t *testing.T) {
	b := buffer{
		lines:            []string{"abc", "def", "ghi"},
		position:         cursorPosition{Line: 2, Col: 3},
		positionRendered: cursorPosition{Line: 2, Col: 3},
	}

	b.Clear()
	assert.Equal(t, []string{""}, b.lines)
	assert.True(t, strings.HasPrefix(b.linesRendered, fmt.Sprint(time.Now().Year())))
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	assert.Equal(t, cursorPosition{Line: 2, Col: 3}, b.positionRendered)
}

func TestBuffer_DeleteBackward(t *testing.T) {
	b := getNewBuffer(t)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 1, Col: 1}
	b.DeleteBackward(1)
	assert.Equal(t, []string{"abc", "ef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
	b.DeleteBackward(1)
	assert.Equal(t, []string{"abcef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)
	b.DeleteBackward(1)
	assert.Equal(t, []string{"abef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteBackward(1)
	assert.Equal(t, []string{"aef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 1}, b.position)
	b.DeleteBackward(1)
	assert.Equal(t, []string{"ef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteBackward(1)
	assert.Equal(t, []string{"ef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.position = cursorPosition{Line: 0, Col: 2}
	b.DeleteBackward(1)
	assert.Equal(t, []string{"e"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 1}, b.position)
	b.DeleteBackward(1)
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteBackward(1)
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def", "ghi"}
	b.position = cursorPosition{Line: 1, Col: 1}
	b.DeleteBackward(-1)
	assert.Equal(t, []string{"ef", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def", "ghi"}
	b.position = cursorPosition{Line: 2, Col: 3}
	b.DeleteBackward(-1)
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
}

func TestBuffer_DeleteForward(t *testing.T) {
	b := getNewBuffer(t)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 0, Col: 2}
	b.DeleteForward(1)
	assert.Equal(t, []string{"ab", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteForward(1)
	assert.Equal(t, []string{"abdef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteForward(1)
	assert.Equal(t, []string{"abef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteForward(1)
	assert.Equal(t, []string{"abf"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteForward(1)
	assert.Equal(t, []string{"ab"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteForward(1)
	assert.Equal(t, []string{"ab"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.DeleteForward(1)
	assert.Equal(t, []string{"b"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteForward(1)
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteForward(1)
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def", "ghi"}
	b.position = cursorPosition{Line: 0, Col: 3}
	b.DeleteForward(1)
	assert.Equal(t, []string{"abcdef", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 0, Col: 3}
	b.DeleteForward(-1)
	assert.Equal(t, []string{"abc"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)

	b.lines = []string{"abc", "def", "ghi"}
	b.position = cursorPosition{Line: 0, Col: 2}
	b.DeleteForward(-1)
	assert.Equal(t, []string{"ab"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
}

func TestBuffer_DeleteWordBackward(t *testing.T) {
	b := getNewBuffer(t)

	b.lines = []string{"abc def"}
	b.position = cursorPosition{Line: 0, Col: 4}
	b.DeleteWordBackward()
	assert.Equal(t, []string{"def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{"def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc def"}
	b.position = cursorPosition{Line: 0, Col: 3}
	b.DeleteWordBackward()
	assert.Equal(t, []string{" def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{" def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc def ghi"}
	b.position = cursorPosition{Line: 0, Col: 11}
	b.DeleteWordBackward()
	assert.Equal(t, []string{"abc def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 7}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{"abc"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 1, Col: 3}
	b.DeleteWordBackward()
	assert.Equal(t, []string{"abc", ""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{"abc"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)

	b.lines = []string{"abc", "def", "ghi"}
	b.position = cursorPosition{Line: 2, Col: 3}
	b.DeleteWordBackward()
	assert.Equal(t, []string{"abc", "def", ""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 0}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{"abc", ""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{"abc"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteWordBackward()
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
}

func TestBuffer_DeleteWordForward(t *testing.T) {
	b := getNewBuffer(t)

	b.lines = []string{"abc def"}
	b.position = cursorPosition{Line: 0, Col: 4}
	b.DeleteWordForward()
	assert.Equal(t, []string{"abc "}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 4}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{"abc "}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 4}, b.position)

	b.lines = []string{"abc def"}
	b.position = cursorPosition{Line: 0, Col: 2}
	b.DeleteWordForward()
	assert.Equal(t, []string{"abdef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{"ab"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{"ab"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 0, Col: 2}
	b.DeleteWordForward()
	assert.Equal(t, []string{"ab", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{"abdef"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)

	b.lines = []string{"abc", "def", "ghi"}
	b.position = cursorPosition{Line: 0, Col: 0}
	b.DeleteWordForward()
	assert.Equal(t, []string{"", "def", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{"def", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{"", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{"ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
	b.DeleteWordForward()
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
}

func TestBuffer_HasChanges(t *testing.T) {
	b := getNewBuffer(t)

	b.linesRendered = "[]"
	assert.False(t, b.HasChanges())

	b.lines = []string{"abc", "def"}
	assert.True(t, b.HasChanges())

	b.linesRendered = "[abc def]"
	assert.False(t, b.HasChanges())

	b.position = cursorPosition{Line: 1, Col: 1}
	assert.True(t, b.HasChanges())

	b.positionRendered = b.position
	assert.False(t, b.HasChanges())
}

func TestBuffer_Insert(t *testing.T) {
	b := getNewBuffer(t)

	b.Insert('\n')
	assert.Equal(t, []string{"", ""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)

	b.lines = []string{""}
	b.position = cursorPosition{Line: 0, Col: 0}
	b.Insert('a')
	assert.Equal(t, []string{"a"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 1}, b.position)

	b.Insert('b')
	assert.Equal(t, []string{"ab"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)

	b.Insert('c')
	assert.Equal(t, []string{"abc"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)

	b.Insert('\n')
	assert.Equal(t, []string{"abc", ""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)

	b.Insert('\t')
	assert.Equal(t, []string{"abc", "    "}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 4}, b.position)

	b.Insert('\n')
	assert.Equal(t, []string{"abc", "    ", ""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 0}, b.position)

	b.Insert('d')
	assert.Equal(t, []string{"abc", "    ", "d"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 1}, b.position)

	b.Insert('e')
	assert.Equal(t, []string{"abc", "    ", "de"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 2}, b.position)

	b.Insert('f')
	assert.Equal(t, []string{"abc", "    ", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 3}, b.position)

	b.position = cursorPosition{Line: 1, Col: 2}
	b.Insert('\n')
	assert.Equal(t, []string{"abc", "  ", "  ", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 0}, b.position)

	b.Insert('1')
	assert.Equal(t, []string{"abc", "  ", "1  ", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 1}, b.position)

	b.position = cursorPosition{Line: 2, Col: 3}
	b.Insert('2')
	assert.Equal(t, []string{"abc", "  ", "1  2", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 4}, b.position)

	b.Insert('\n')
	assert.Equal(t, []string{"abc", "  ", "1  2", "", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 3, Col: 0}, b.position)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.Insert('?')
	assert.Equal(t, []string{"?abc", "  ", "1  2", "", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 1}, b.position)
}

func TestBuffer_IsDone(t *testing.T) {
	b := getNewBuffer(t)
	assert.False(t, b.IsDone())

	b.done = true
	assert.True(t, b.IsDone())
}

func TestBuffer_Length(t *testing.T) {
	b := getNewBuffer(t)
	assert.Equal(t, 0, b.Length())

	b.lines = []string{"abc"}
	assert.Equal(t, 3, b.Length())

	b.lines = []string{"abc", ""}
	assert.Equal(t, 4, b.Length())

	b.lines = []string{"abc", "def"}
	assert.Equal(t, 7, b.Length())
}

func TestBuffer_MakeWordCapitalCase(t *testing.T) {
	b := getNewBuffer(t)
	b.lines = []string{"abc", "def", "ghi"}
	b.position = cursorPosition{Line: 0, Col: 0}

	b.MakeWordCapitalCase()
	assert.Equal(t, []string{"Abc", "def", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
	b.position.Col = 1
	b.MakeWordCapitalCase()
	assert.Equal(t, []string{"Abc", "Def", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 0}, b.position)
	b.position.Col = 2
	b.MakeWordCapitalCase()
	assert.Equal(t, []string{"Abc", "Def", "Ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 3}, b.position)
}

func TestBuffer_MakeWordLowerCase(t *testing.T) {
	b := getNewBuffer(t)
	b.lines = []string{"ABC", "DEF", "GHI"}
	b.position = cursorPosition{Line: 0, Col: 0}

	b.MakeWordLowerCase()
	assert.Equal(t, []string{"abc", "DEF", "GHI"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
	b.position.Col = 1
	b.MakeWordLowerCase()
	assert.Equal(t, []string{"abc", "def", "GHI"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 0}, b.position)
	b.position.Col = 2
	b.MakeWordLowerCase()
	assert.Equal(t, []string{"abc", "def", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 3}, b.position)
}

func TestBuffer_MakeWordUpperCase(t *testing.T) {
	b := getNewBuffer(t)
	b.lines = []string{"abc", "def", "ghi"}
	b.position = cursorPosition{Line: 0, Col: 0}

	b.MakeWordUpperCase()
	assert.Equal(t, []string{"ABC", "def", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
	b.position.Col = 1
	b.MakeWordUpperCase()
	assert.Equal(t, []string{"ABC", "DEF", "ghi"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 0}, b.position)
	b.position.Col = 2
	b.MakeWordUpperCase()
	assert.Equal(t, []string{"ABC", "DEF", "GHI"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 2, Col: 3}, b.position)
}

func TestBuffer_MarkAsDone(t *testing.T) {
	b := getNewBuffer(t)
	assert.False(t, b.done)

	b.MarkAsDone()
	assert.True(t, b.done)
}

func TestBuffer_MoveDown(t *testing.T) {
	b := getNewBuffer(t)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveDown(1)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc123", "def"}
	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveDown(1)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)

	b.position = cursorPosition{Line: 0, Col: 2}
	b.MoveDown(1)
	assert.Equal(t, cursorPosition{Line: 1, Col: 2}, b.position)
	b.MoveDown(1)
	assert.Equal(t, cursorPosition{Line: 1, Col: 2}, b.position)

	b.position = cursorPosition{Line: 0, Col: 5}
	b.MoveDown(1)
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)
	b.MoveDown(1)
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)
}

func TestBuffer_MoveLeft(t *testing.T) {
	b := getNewBuffer(t)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveLeft(1)
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 1, Col: 3}
	b.MoveLeft(-1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.position = cursorPosition{Line: 1, Col: 3}
	b.MoveLeft(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 2}, b.position)

	b.MoveLeft(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 1}, b.position)

	b.MoveLeft(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)

	b.MoveLeft(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)

	b.MoveLeft(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)

	b.MoveLeft(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 1}, b.position)

	b.MoveLeft(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.MoveLeft(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
}

func TestBuffer_MoveLineBegin(t *testing.T) {
	b := getNewBuffer(t)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveLineBegin()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 1, Col: 3}
	b.MoveLineBegin()
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)

	b.position = cursorPosition{Line: 1, Col: 2}
	b.MoveLineBegin()
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)

	b.position = cursorPosition{Line: 0, Col: 3}
	b.MoveLineBegin()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
}

func TestBuffer_MoveLineEnd(t *testing.T) {
	b := getNewBuffer(t)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveLineEnd()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 1, Col: 3}
	b.MoveLineEnd()
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)

	b.position = cursorPosition{Line: 1, Col: 2}
	b.MoveLineEnd()
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)

	b.position = cursorPosition{Line: 0, Col: 3}
	b.MoveLineEnd()
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)

	b.position = cursorPosition{Line: 0, Col: 2}
	b.MoveLineEnd()
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)
}

func TestBuffer_MoveRight(t *testing.T) {
	b := getNewBuffer(t)

	b.MoveRight(1)
	assert.Equal(t, []string{""}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def"}
	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveRight(-1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)

	b.position = cursorPosition{Line: 1, Col: 3}
	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 1}, b.position)

	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)

	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)

	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)

	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 1}, b.position)

	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 2}, b.position)

	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)

	b.MoveRight(1)
	assert.Equal(t, []string{"abc", "def"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)
}

func TestBuffer_MoveUp(t *testing.T) {
	b := getNewBuffer(t)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveUp(1)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc", "def123"}
	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveUp(1)
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.position = cursorPosition{Line: 1, Col: 2}
	b.MoveUp(1)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)
	b.MoveUp(1)
	assert.Equal(t, cursorPosition{Line: 0, Col: 2}, b.position)

	b.position = cursorPosition{Line: 1, Col: 5}
	b.MoveUp(1)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)
	b.MoveUp(1)
	assert.Equal(t, cursorPosition{Line: 0, Col: 3}, b.position)
}

func TestBuffer_MoveWordLeft(t *testing.T) {
	b := getNewBuffer(t)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc 123", "def"}
	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.position = cursorPosition{Line: 0, Col: 1}
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.position = cursorPosition{Line: 0, Col: 2}
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.position = cursorPosition{Line: 0, Col: 3}
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.position = cursorPosition{Line: 0, Col: 4}
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.position = cursorPosition{Line: 1, Col: 3}
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 0, Col: 4}, b.position)
	b.MoveWordLeft()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)
}

func TestBuffer_MoveWordRight(t *testing.T) {
	b := getNewBuffer(t)

	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveWordRight()
	assert.Equal(t, cursorPosition{Line: 0, Col: 0}, b.position)

	b.lines = []string{"abc 123 ", "def"}
	b.position = cursorPosition{Line: 0, Col: 0}
	b.MoveWordRight()
	assert.Equal(t, cursorPosition{Line: 0, Col: 4}, b.position)
	b.MoveWordRight()
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
	b.MoveWordRight()
	assert.Equal(t, cursorPosition{Line: 1, Col: 3}, b.position)

	b.position = cursorPosition{Line: 0, Col: 1}
	b.MoveWordRight()
	assert.Equal(t, cursorPosition{Line: 0, Col: 4}, b.position)

	b.position = cursorPosition{Line: 0, Col: 2}
	b.MoveWordRight()
	assert.Equal(t, cursorPosition{Line: 0, Col: 4}, b.position)

	b.position = cursorPosition{Line: 0, Col: 3}
	b.MoveWordRight()
	assert.Equal(t, cursorPosition{Line: 0, Col: 4}, b.position)

	b.position = cursorPosition{Line: 0, Col: 8}
	b.MoveWordRight()
	assert.Equal(t, cursorPosition{Line: 1, Col: 0}, b.position)
}

func TestBuffer_Render(t *testing.T) {

}

func TestBuffer_Set(t *testing.T) {
	b := getNewBuffer(t)

	b.Set("echo $VARIABLE\necho $VARIABLE2\t#testing")
	assert.Equal(t, []string{"echo $VARIABLE", "echo $VARIABLE2    #testing"}, b.lines)
	assert.Equal(t, cursorPosition{Line: 1, Col: 27}, b.position)
}

func TestBuffer_String(t *testing.T) {
	b := getNewBuffer(t)
	assert.Equal(t, "", b.String())

	b.lines = []string{"abc"}
	assert.Equal(t, "abc", b.String())

	b.lines = []string{"abc", ""}
	assert.Equal(t, "abc\n", b.String())

	b.lines = []string{"abc", "def"}
	assert.Equal(t, "abc\ndef", b.String())
}

func TestBuffer_SwapCharacterNext(t *testing.T) {

}

func TestBuffer_SwapCharacterPrevious(t *testing.T) {

}

func TestBuffer_SwapWordNext(t *testing.T) {

}

func TestBuffer_SwapWordPrevious(t *testing.T) {

}
