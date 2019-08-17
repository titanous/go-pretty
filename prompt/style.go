package prompt

import (
	"time"

	"github.com/jedib0t/go-pretty/text"
)

// Style declares how to render the InputContext and provides very fine-grained
// control on how the InputContext gets rendered on the Console.
type Style struct {
	Name              string
	AutoComplete      AutoCompleteOptions
	Colors            text.Colors
	NewlineIndent     string
	Prefix            string
	SyntaxHighlighter SyntaxHighlighterOptions
	Tab               string
	Timestamp         TimestampOptions
}

var (
	// StyleDefault defines sensible bare-minimum options for the InputContext
	StyleDefault = Style{
		Name:              "StyleDefault",
		AutoComplete:      AutoCompleteOptionsDefault,
		Colors:            text.Colors{},
		NewlineIndent:     "... ",
		Prefix:            "> ",
		SyntaxHighlighter: SyntaxHighlighterOptionsDefault,
		Tab:               "    ",
		Timestamp:         TimestampOptionsDefault,
	}

	// StyleColored defines a colored prompt
	StyleColored = Style{
		Name:              "StyleDefault",
		AutoComplete:      AutoCompleteOptionsDefault,
		Colors:            text.Colors{text.FgHiYellow},
		NewlineIndent:     text.FgHiBlack.Sprint("... "),
		Prefix:            text.Colors{text.Bold, text.FgGreen}.Sprint(">>> "),
		SyntaxHighlighter: SyntaxHighlighterOptionsDefault,
		Tab:               "    ",
		Timestamp:         TimestampOptionsDefault,
	}
)

// AutoCompleteOptions defines the auto-completion options.
type AutoCompleteOptions struct {
	MaxSuggestions     int
	Hint               text.Colors
	HintSelected       text.Colors
	Suggestion         text.Colors
	SuggestionSelected text.Colors
}

var (
	// AutoCompleteOptionsDefault defines sensible auto-completion style.
	AutoCompleteOptionsDefault = AutoCompleteOptionsSimple

	// AutoCompleteOptionsSimple defines a simple but readable style.
	AutoCompleteOptionsSimple = AutoCompleteOptions{
		MaxSuggestions:     5,
		Hint:               text.Colors{text.BgBlue, text.FgWhite, text.Italic},
		HintSelected:       text.Colors{text.BgHiBlue, text.FgHiWhite, text.Bold, text.Italic},
		Suggestion:         text.Colors{text.BgBlue, text.FgWhite},
		SuggestionSelected: text.Colors{text.BgHiBlue, text.FgHiWhite, text.Bold},
	}
)

// SyntaxHighlighterOptions defines the syntax highlighting options. This is
// internally passed on to the awesome library Chroma. To find out all supported
// formatters, languages and styles, please refer to:
//   https://github.com/alecthomas/chroma
//
// When the Language value is empty, Chroma tries to auto-determine the language
// and may not exactly be what you want. When the Style value is empty, Chroma
// uses the fallback style (swapoff).
type SyntaxHighlighterOptions struct {
	Enabled   bool
	Formatter string
	Language  string
	Style     string
}

var (
	// SyntaxHighlighterOptionsDefault defines sensible syntax highlighting
	// options - basically none.
	SyntaxHighlighterOptionsDefault = SyntaxHighlighterOptionsOff

	// SyntaxHighlighterOptionsOff turns off syntax highlighting.
	SyntaxHighlighterOptionsOff = SyntaxHighlighterOptions{
		Enabled:   false,
		Formatter: "",
		Language:  "",
		Style:     "",
	}

	// SyntaxHighlighterOptionsGolang enables syntax highlighting and expects
	// Golang content
	SyntaxHighlighterOptionsGolang = SyntaxHighlighterOptions{
		Enabled:   true,
		Formatter: "terminal256",
		Language:  "go",
		Style:     "vim",
	}

	// SyntaxHighlighterOptionsPython enables syntax highlighting and expects
	// Python content
	SyntaxHighlighterOptionsPython = SyntaxHighlighterOptions{
		Enabled:   true,
		Formatter: "terminal256",
		Language:  "python",
		Style:     "vim",
	}
	// SyntaxHighlighterOptionsSQL enables syntax highlighting and expects
	// SQL content
	SyntaxHighlighterOptionsSQL = SyntaxHighlighterOptions{
		Enabled:   true,
		Formatter: "terminal256",
		Language:  "sql",
		Style:     "vim",
	}
)

// TimestampOptions controls how the Timestamp is rendered.
type TimestampOptions struct {
	Colors       text.Colors
	Enabled      bool
	Layout       string
	Location     *time.Location
	PaddingLeft  string // this will not be colorized with the Colors above
	PaddingRight string // this will not be colorized with the Colors above
	Prefix       string
	Suffix       string
}

var (
	// TimestampOptionsDefault defines sensible Timestamp options - none.
	TimestampOptionsDefault = TimestampOptionsOff

	// TimestampOptionsOff disables the timestamp on the prompt.
	TimestampOptionsOff = TimestampOptions{
		Enabled: false,
	}

	// TimestampOptionsSimple defines sensible Timestamp options.
	TimestampOptionsSimple = TimestampOptions{
		Colors:       text.Colors{text.FgHiBlack},
		Enabled:      true,
		Layout:       "2006-01-02 15:04:05 MST",
		Location:     time.Local,
		PaddingLeft:  " ",
		PaddingRight: "",
		Prefix:       "[",
		Suffix:       "]",
	}
)

// Generate generates the timestamp with the provided options
func (t *TimestampOptions) Generate() string {
	if !t.Enabled {
		return ""
	}
	// default to local location if no location specified
	if t.Location == nil {
		t.Location = time.Local
	}

	return t.Colors.Sprintf("%s%s%s",
		t.Prefix, time.Now().In(t.Location).Format(t.Layout), t.Suffix,
	) + t.PaddingLeft
}
