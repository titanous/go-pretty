package prompt

import (
	"fmt"
	"sort"
	"strings"
)

// KeyMap can be used to customize or define the behavior of the Prompt for each
// special Key sequences that is entered by the User.
type KeyMap struct {
	Insert       InsertKeyMap
	AutoComplete AutoCompleteKeyMap

	errors []error
}

// InsertKeyMap is the KeyMap used in Insert mode.
type InsertKeyMap struct {
	Abort                  KeySequences
	DeleteCharCurrent      KeySequences
	DeleteCharPrevious     KeySequences
	DeleteWordNext         KeySequences
	DeleteWordPrevious     KeySequences
	EraseEverything        KeySequences
	EraseToBeginningOfLine KeySequences
	EraseToEndOfLine       KeySequences
	HistoryDown            KeySequences
	HistoryUp              KeySequences
	MakeWordCapitalCase    KeySequences
	MakeWordLowerCase      KeySequences
	MakeWordUpperCase      KeySequences
	MoveDownOneLine        KeySequences
	MoveLeftOneCharacter   KeySequences
	MoveRightOneCharacter  KeySequences
	MoveToBeginning        KeySequences
	MoveToBeginningOfLine  KeySequences
	MoveToEnd              KeySequences
	MoveToEndOfLine        KeySequences
	MoveToWordNext         KeySequences
	MoveToWordPrevious     KeySequences
	MoveUpOneLine          KeySequences
	SwapCharacterNext      KeySequences
	SwapCharacterPrevious  KeySequences
	SwapWordNext           KeySequences
	SwapWordPrevious       KeySequences
	Terminate              KeySequences
}

// AutoCompleteKeyMap is the KeyMap used in AutoComplete mode.
type AutoCompleteKeyMap struct {
	ChooseNext     KeySequences
	ChoosePrevious KeySequences
	Hide           KeySequences
	Select         KeySequences
}

// keyMapReversed is an internal representation of the KeyMap for easy
// programmatic access when acting on key sequences.
type keyMapReversed struct {
	Insert       map[KeySequence]Action
	AutoComplete map[KeySequence]Action
}

var (
	// DefaultKeyMap that defines sane key sequences for each supported action.
	// Quite a few of these are the default short-cuts in BASH.
	DefaultKeyMap = KeyMap{
		Insert: InsertKeyMap{
			Abort:                  KeySequences{CtrlC, CtrlD},
			DeleteCharCurrent:      KeySequences{Delete},
			DeleteCharPrevious:     KeySequences{Backspace, CtrlH},
			DeleteWordNext:         KeySequences{AltD},
			DeleteWordPrevious:     KeySequences{CtrlW},
			EraseEverything:        KeySequences{AltW},
			EraseToBeginningOfLine: KeySequences{CtrlU},
			EraseToEndOfLine:       KeySequences{CtrlK},
			HistoryDown:            KeySequences{CtrlArrowDown},
			HistoryUp:              KeySequences{CtrlArrowUp},
			MakeWordCapitalCase:    KeySequences{AltC},
			MakeWordLowerCase:      KeySequences{AltL},
			MakeWordUpperCase:      KeySequences{AltU},
			MoveDownOneLine:        KeySequences{ArrowDown},
			MoveLeftOneCharacter:   KeySequences{ArrowLeft},
			MoveRightOneCharacter:  KeySequences{ArrowRight},
			MoveToBeginning:        KeySequences{CtrlHome},
			MoveToBeginningOfLine:  KeySequences{Home},
			MoveToEnd:              KeySequences{CtrlEnd},
			MoveToEndOfLine:        KeySequences{End},
			MoveToWordNext:         KeySequences{CtrlArrowRight, AltF},
			MoveToWordPrevious:     KeySequences{CtrlArrowLeft, AltB},
			MoveUpOneLine:          KeySequences{ArrowUp},
			SwapCharacterNext:      KeySequences{CtrlN},
			SwapCharacterPrevious:  KeySequences{CtrlT},
			SwapWordNext:           KeySequences{AltN},
			SwapWordPrevious:       KeySequences{AltT},
			Terminate:              KeySequences{Enter, CtrlM},
		},
		AutoComplete: AutoCompleteKeyMap{
			ChooseNext:     KeySequences{ArrowDown},
			ChoosePrevious: KeySequences{ArrowUp},
			Hide:           KeySequences{Escape},
			Select:         KeySequences{Tab, CtrlI},
		},
	}
)

func (k *KeyMap) reverse() (*keyMapReversed, error) {
	rsp := &keyMapReversed{
		Insert:       make(map[KeySequence]Action),
		AutoComplete: make(map[KeySequence]Action),
	}

	k.errors = make([]error, 0)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.Abort, Abort)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.DeleteCharCurrent, DeleteCharCurrent)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.DeleteCharPrevious, DeleteCharPrevious)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.DeleteWordNext, DeleteWordNext)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.DeleteWordPrevious, DeleteWordPrevious)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.EraseEverything, EraseEverything)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.EraseToBeginningOfLine, EraseToBeginningOfLine)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.EraseToEndOfLine, EraseToEndOfLine)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.HistoryDown, HistoryDown)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.HistoryUp, HistoryUp)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MakeWordCapitalCase, MakeWordCapitalCase)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MakeWordLowerCase, MakeWordLowerCase)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MakeWordUpperCase, MakeWordUpperCase)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveDownOneLine, MoveDownOneLine)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveLeftOneCharacter, MoveLeftOneCharacter)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveRightOneCharacter, MoveRightOneCharacter)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveToBeginning, MoveToBeginning)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveToBeginningOfLine, MoveToBeginningOfLine)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveToEnd, MoveToEnd)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveToEndOfLine, MoveToEndOfLine)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveToWordNext, MoveToWordNext)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveToWordPrevious, MoveToWordPrevious)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.MoveUpOneLine, MoveUpOneLine)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.SwapCharacterNext, SwapCharacterNext)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.SwapCharacterPrevious, SwapCharacterPrevious)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.SwapWordNext, SwapWordNext)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.SwapWordPrevious, SwapWordPrevious)
	k.reverseAddKeySequences(rsp.Insert, k.Insert.Terminate, Terminate)
	k.reverseAddKeySequences(rsp.AutoComplete, k.AutoComplete.ChooseNext, AutoCompleteChooseNext)
	k.reverseAddKeySequences(rsp.AutoComplete, k.AutoComplete.ChoosePrevious, AutoCompleteChoosePrevious)
	k.reverseAddKeySequences(rsp.AutoComplete, k.AutoComplete.Hide, AutoCompleteHide)
	k.reverseAddKeySequences(rsp.AutoComplete, k.AutoComplete.Select, AutoCompleteSelect)
	if len(k.errors) > 0 {
		errStrs := make([]string, len(k.errors))
		for idx, err := range k.errors {
			errStrs[idx] = fmt.Sprintf("- %v", err.Error())
		}
		sort.Strings(errStrs)
		return nil, fmt.Errorf("key-map has errors:\n%v", strings.Join(errStrs, "\n"))
	}

	return rsp, nil
}

func (k *KeyMap) reverseAddKeySequences(m map[KeySequence]Action, keySequences KeySequences, action Action) {
	for _, keySequence := range keySequences {
		if existingAction, ok := m[keySequence]; ok {
			k.errors = append(k.errors, fmt.Errorf(
				"more than one action defined for %v: [%v, %v]",
				keySequence, existingAction, action,
			))
		}
		m[keySequence] = action
	}
}
