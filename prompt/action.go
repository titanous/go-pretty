package prompt

// Action defines an activity that is done based on a key sequence.
type Action string

// Supported Actions.
const (
	None  Action = ""      // no action
	Abort Action = "Abort" // abort the prompt completely and return to caller

	// insert-mode actions
	DeleteCharCurrent      Action = "DeleteCharCurrent"      // delete the character at the cursor
	DeleteCharPrevious     Action = "DeleteCharPrevious"     // delete the character before the cursor
	DeleteWordNext         Action = "DeleteWordNext"         // delete the next work
	DeleteWordPrevious     Action = "DeleteWordPrevious"     // delete the previous word
	EraseEverything        Action = "EraseEverything"        // erase the entire prompt
	EraseToBeginningOfLine Action = "EraseToBeginningOfLine" // erase from cursor to the the beginning of current line
	EraseToEndOfLine       Action = "EraseToEndOfLine"       // erase from cursor to the the end of current line
	HistoryDown            Action = "HistoryDown"            // show command executed after current command if any
	HistoryUp              Action = "HistoryUp"              // show previously executed command if any
	MakeWordCapitalCase    Action = "MakeWordCapitalCase"    // make the word at the cursor capitalized
	MakeWordLowerCase      Action = "MakeWordLowerCase"      // make the word at the cursor lower case
	MakeWordUpperCase      Action = "MakeWordUpperCase"      // make the word at the cursor upper case
	MoveDownOneLine        Action = "MoveDownOneLine"        // move the cursor down one line
	MoveLeftOneCharacter   Action = "MoveLeftOneCharacter"   // move the cursor left one character
	MoveRightOneCharacter  Action = "MoveRightOneCharacter"  // move the cursor right one character
	MoveUpOneLine          Action = "MoveUpOneLine"          // move the cursor up one line
	MoveToBeginning        Action = "MoveToBeginning"        // move to the beginning of the entire prompt text
	MoveToBeginningOfLine  Action = "MoveToBeginningOfLine"  // move to the beginning of the current line
	MoveToEnd              Action = "MoveToEnd"              // move to the end of the entire prompt text
	MoveToEndOfLine        Action = "MoveToEndOfLine"        // move to the end of the current line
	MoveToWordNext         Action = "MoveToWordNext"         // move to the beginning of the next word
	MoveToWordPrevious     Action = "MoveToWordPrevious"     // move to the beginning of the previous word
	SwapCharacterNext      Action = "SwapCharacterNext"      // swap the character at the cursor with the next
	SwapCharacterPrevious  Action = "SwapCharacterPrevious"  // swap the character at the cursor with the previous
	SwapWordNext           Action = "SwapWordNext"           // swap the word at the cursor with the next
	SwapWordPrevious       Action = "SwapWordPrevious"       // swap the word at the cursor with the previous
	Terminate              Action = "Terminate"              // trigger the termination checker if any, or call the callback function

	// auto-complete specific actions
	AutoCompleteChooseNext     Action = "AutoCompleteChooseNext"     // choose the next suggestion
	AutoCompleteChoosePrevious Action = "AutoCompleteChoosePrevious" // choose the previous suggestion
	AutoCompleteHide           Action = "AutoCompleteHide"           // hide the auto-complete suggestions pop-up
	AutoCompleteSelect         Action = "AutoCompleteSelect"         // select the current suggestion
)
