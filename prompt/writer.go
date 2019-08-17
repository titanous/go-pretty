package prompt

// Callback defines the function that will be called for every user input.
type Callback func(userInput string, err error) error

// Writer declares the interfaces that can be used to setup and render a prompt.
type Writer interface {
	Debug()
	DoNotUseTermbox()
	KeyMap() *KeyMap
	Render(callback Callback) error
	SetAutoCompleter(autoCompleter AutoCompleter)
	SetHistory(history []string)
	SetKeyMap(keyMap KeyMap)
	SetRefreshRate(refreshRate int)
	SetStyle(style Style)
	SetTitle(title string)
	SetTerminationChecker(terminationChecker TerminationChecker)
	Style() *Style
}

// NewWriter initializes and returns a Writer.
func NewWriter() Writer {
	return &prompt{}
}
