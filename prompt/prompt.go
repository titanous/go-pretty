package prompt

import (
	"fmt"
	"os"
	"time"
)

// Defaults
const (
	DefaultRefreshRate = 60 // times a second
)

type prompt struct {
	// autoCompleter helps with auto-completion
	autoCompleter AutoCompleter
	// debugLogFile contains a pointer to a File where all debug logs are printed
	debugLogFile *os.File
	// debug controls debug logs to the debugLogFile
	debug bool
	// doNotUseTermbox forces the use of "github.com/eiannone/keyboard" instead
	// of "github.com/nsf/termbox-go" library
	doNotUseTermbox bool
	// done is the channel used to send the termination signal for the renderers
	done chan bool
	// escapeSeq stores the current Escape sequence being handled (ex.: \x1b[1;5D)
	escapeSeq string
	// maxHeight defines the maximum height of the terminal
	maxHeight int
	// maxWidth defines the maximum width of the terminal
	maxWidth int
	// history contains all previous user inputs
	history []string
	// historyIdx contains the current history location
	historyIdx int
	// keyMap stores the mapping of supported key sequences for each action
	keyMap *KeyMap
	// keySequenceActionMap is the simplified form of keyMap
	keyMapReversed *keyMapReversed
	// refreshRate is the number of times to redraw the prompt a second
	refreshRate int
	// style contains all the options for the prompt, and more
	style *Style
	// title contains the string to be rendered before the first prompt
	title string
	// terminationChecker contains the function that returns true if the
	// user-prompt is over and the command is ready to the returned to caller
	terminationChecker TerminationChecker
}

// Debug turns on debugging mode and logs to a "prompt.log" file
func (p *prompt) Debug() {
	p.debug = true
}

// DoNotUseTermbox tries to render the prompt without using Termbox.
func (p *prompt) DoNotUseTermbox() {
	p.doNotUseTermbox = true
}

// KeyMap returns the KeyMap currently in use.
func (p *prompt) KeyMap() *KeyMap {
	if p.keyMap == nil {
		tempKeyMap := DefaultKeyMap
		p.keyMap = &tempKeyMap
	}
	return p.keyMap
}

// SetAutoCompleter sets up the auto-completer to be called when the user is
// typing something.
func (p *prompt) SetAutoCompleter(autoCompleter AutoCompleter) {
	p.autoCompleter = autoCompleter
}

// SetHistory sets up the user input history to a previous known state.
func (p *prompt) SetHistory(history []string) {
	p.history = history
	p.historyIdx = len(p.history)
}

// SetKeyMap sets the KeyMap to use when dealing with user input.
func (p *prompt) SetKeyMap(keyMap KeyMap) {
	p.keyMap = &keyMap
}

// SetRefreshRate sets the refresh rate (number of redraws per second). This
// rate is attempted but cannot be guaranteed due to the non-deterministic
// nature of the Auto-Completion logic and TerminationChecker logic.
//
// Acceptable values are between 1 and 120.
func (p *prompt) SetRefreshRate(refreshRate int) {
	if refreshRate > 1 && refreshRate <= 120 {
		p.refreshRate = refreshRate
	}
}

// SetStyle overrides the DefaultStyle with the provided one.
func (p *prompt) SetStyle(style Style) {
	p.style = &style
}

// SetTitle sets the title for the prompt
func (p *prompt) SetTitle(title string) {
	p.title = title
}

// SetTerminationChecker sets up the function to be called every time user
// presses "Enter" to determine if the prompt is over and the user is ready for
// the command to be executed.
//
// This is necessary if you want a multi-line prompt and don't want to
// prematurely terminate the input on user pressing "Enter".
//
// If not provided, a simple "Enter" will terminate Render().
func (p *prompt) SetTerminationChecker(terminationChecker TerminationChecker) {
	p.terminationChecker = terminationChecker
}

// Style returns the current style.
func (p *prompt) Style() *Style {
	if p.style == nil {
		tempStyle := StyleDefault
		p.style = &tempStyle
	}
	return p.style
}

func (p *prompt) appendToHistory(input string) {
	if input != "" && p.history[len(p.history)-1] != input {
		p.history = append(p.history, input)
		p.historyIdx = len(p.history)
		//fmt.Printf("\n++ %#v: %#v", p.historyIdx, p.history[p.historyIdx-1])
	}
}

func (p *prompt) cleanup() {
	if p.debugLogFile != nil {
		_ = p.debugLogFile.Close()
	}
}

func (p *prompt) debugLog(msg string, a ...interface{}) {
	if p.debugLogFile != nil {
		msg = fmt.Sprintf(msg, a...)
		_, _ = p.debugLogFile.WriteString(fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), msg))
		_ = p.debugLogFile.Sync()
	}
}

func (p *prompt) getHistoryItem(delta int) string {
	p.historyIdx += delta
	if p.historyIdx < 0 {
		p.historyIdx = 0
	} else if p.historyIdx > len(p.history) {
		p.historyIdx = len(p.history)
	}

	if p.historyIdx >= 0 && p.historyIdx < len(p.history) {
		p.debugLog("getHistoryItem: idx=%v, value=%#v", p.historyIdx, p.history[p.historyIdx])
		return p.history[p.historyIdx]
	}
	return ""
}

func (p *prompt) initForRender() (err error) {
	// pick defaults
	p.KeyMap()
	p.Style()

	p.escapeSeq = ""
	p.keyMapReversed, err = p.keyMap.reverse()
	if err != nil {
		return err
	}
	if p.refreshRate == 0 {
		p.refreshRate = DefaultRefreshRate
	}
	p.historyIdx = len(p.history)

	// debugging
	if p.debug {
		p.debugLogFile, _ = os.Create("prompt.log")
	}

	return nil
}
