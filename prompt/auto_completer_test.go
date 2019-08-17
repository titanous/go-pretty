package prompt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoCompleteGoLangKeywords(t *testing.T) {
	ac := AutoCompleteGoLangKeywords()

	matches := ac("foo", "foo", 0)
	assert.Empty(t, matches)

	matches = ac("ca", "ca", 0)
	assert.Empty(t, matches)
	matches = ac("cas", "cas", 0)
	assert.NotEmpty(t, matches)
	assert.Len(t, matches, 1)
	assert.Equal(t, matches[0].Value, "case")
	matches = ac("case", "case", 0)
	assert.Empty(t, matches)
}

func TestAutoCompletePythonKeywords(t *testing.T) {
	ac := AutoCompletePythonKeywords()

	matches := ac("foo", "foo", 0)
	assert.Empty(t, matches)

	matches = ac("ba", "ba", 0)
	assert.Empty(t, matches)
	matches = ac("bas", "bas", 0)
	assert.NotEmpty(t, matches)
	assert.Len(t, matches, 1)
	assert.Equal(t, matches[0].Value, "basestring")

	matches = ac("Bas", "Bas", 0)
	assert.NotEmpty(t, matches)
	assert.Len(t, matches, 1)
	assert.Equal(t, matches[0].Value, "BaseException")

	matches = ac("ass", "ass", 0)
	assert.NotEmpty(t, matches)
	assert.Len(t, matches, 1)
	assert.Equal(t, matches[0].Value, "assert")
	matches = ac("assert", "assert", 0)
	assert.Empty(t, matches)
}

func TestAutoCompleteSQLKeywords(t *testing.T) {
	ac := AutoCompleteSQLKeywords()

	matches := ac("foo", "foo", 0)
	assert.Empty(t, matches)

	matches = ac("se", "sel", 0)
	assert.NotEmpty(t, matches)
	assert.Len(t, matches, 2)
	assert.Equal(t, matches[0].Value, "SELECT")
	assert.Equal(t, matches[1].Value, "SELF")
	matches = ac("sele", "sele", 0)
	assert.NotEmpty(t, matches)
	assert.Len(t, matches, 1)
	assert.Equal(t, matches[0].Value, "SELECT")
	matches = ac("select", "select", 0)
	assert.Empty(t, matches)

	matches = ac("SELECT * fr", "fr", 9)
	assert.Empty(t, matches)
	matches = ac("SELECT * fro", "fro", 9)
	assert.NotEmpty(t, matches)
	assert.Len(t, matches, 1)
	assert.Equal(t, matches[0].Value, "FROM")
	matches = ac("SELECT * from", "from", 9)
	assert.Empty(t, matches)
}

func TestAutoCompleteWords(t *testing.T) {
	t.Run("case insensitive", func(t *testing.T) {
		possibleWords := []string{"foo", "BAZ", "bar"}
		ac := AutoCompleteWords(possibleWords, 2, true)

		matches := ac("A Big Croc", "Croc", 6)
		assert.Empty(t, matches)

		matches = ac("fo", "f", 0)
		assert.Empty(t, matches)
		matches = ac("fo", "fo", 0)
		assert.NotEmpty(t, matches)
		assert.Len(t, matches, 1)
		assert.Equal(t, matches[0].Value, "foo")
		matches = ac("foo", "foo", 0)
		assert.Empty(t, matches)

		matches = ac("foo BA", "BA", 4)
		assert.NotEmpty(t, matches)
		assert.Len(t, matches, 2)
		assert.Equal(t, matches[0].Value, "BAZ")
		assert.Equal(t, matches[1].Value, "bar")
	})

	t.Run("case sensitive", func(t *testing.T) {
		possibleWords := []string{"foo", "BAZ", "bar"}
		ac := AutoCompleteWords(possibleWords, 2, false)

		matches := ac("A Big Croc", "Croc", 6)
		assert.Empty(t, matches)

		matches = ac("fo", "f", 0)
		assert.Empty(t, matches)
		matches = ac("fo", "fo", 0)
		assert.NotEmpty(t, matches)
		assert.Len(t, matches, 1)
		assert.Equal(t, matches[0].Value, "foo")
		matches = ac("foo", "foo", 0)
		assert.Empty(t, matches)

		matches = ac("foo ba", "ba", 4)
		assert.NotEmpty(t, matches)
		assert.Len(t, matches, 1)
		assert.Equal(t, matches[0].Value, "bar")

		matches = ac("foo BA", "BA", 4)
		assert.NotEmpty(t, matches)
		assert.Len(t, matches, 1)
		assert.Equal(t, matches[0].Value, "BAZ")
	})
}
