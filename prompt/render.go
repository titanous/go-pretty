package prompt

import (
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/nsf/termbox-go"
)

// Constants
const (
	// EscapeSequenceDelay defines the time to wait after receiving an Esc code
	// "/x1b" before treating it as an Escape key-hit instead of waiting for a
	// special sequence like "/x1b[1;5A"
	EscapeSequenceDelay = time.Millisecond * 10
)

// Errors returned by Render
var (
	ErrAbort = errors.New("abort")
)

// Render renders the prompt, and waits for user-input. It highlights syntax if
// asked to, shows auto-complete options when available.
//
// Every time the user enters a complete input, this will be sent to the
// Callback function.
//
// Code terminates when the Callback returns an error (which will be blindly
// returned by Render to the original caller).
func (p *prompt) Render(callback Callback) error {
	// init for rendering
	if err := p.initForRender(); err != nil {
		return err
	}
	defer p.cleanup()

	// init the sub-system to handle key strokes
	if !p.doNotUseTermbox {
		err := termbox.Init()
		if err != nil {
			return fmt.Errorf("failed to initialize termbox: %v", err)
		}
		defer termbox.Close()
	} else {
		err := keyboard.Open()
		if err != nil {
			return fmt.Errorf("failed to initialize keyboard: %v", err)
		}
		defer keyboard.Close()
	}

	// show title
	if p.title != "" {
		fmt.Println(p.title)
	}

	// prompt the user until callback returns error
	for {
		// poll for user input and append to history
		p.debugLog("--------------------------------------------------------------------------------")
		input, err := p.readUserInput()
		p.debugLog("input: %#v", input)
		p.debugLog("err: %#v", err)
		if input != "" {
			p.appendToHistory(input)
		}
		fmt.Println()

		// call callback and ensure it is ok to continue
		err = callback(input, err)
		if err != nil {
			p.debugLog("callback.err: %#v", err.Error())
			return err
		}
	}
}

func (p *prompt) readUserInput() (string, error) {
	buffer := newBuffer(p.autoCompleter, p.style)
	chErrors := make(chan error, 1)

	// poll for events continuously
	go p.pollKeyboard(buffer, chErrors)

	// redraw Line times a second
	tick := time.Tick(time.Second / time.Duration(p.refreshRate))
	for {
		select {
		case err := <-chErrors:
			return "", err
		case <-tick:
			if buffer.HasChanges() {
				fmt.Print(buffer.Render())
			}
			if buffer.IsDone() {
				return buffer.String(), nil
			}
		}
	}
}

func (p *prompt) pollKeyboard(buffer *buffer, chErrors chan error) {
	defer func() {
		if r := recover(); r != nil {
			chErrors <- fmt.Errorf("panic: %v: %v", r, string(debug.Stack()))
		}
	}()

	var ch rune
	var key termbox.Key
	var err error
	for {
		if !p.doNotUseTermbox {
			event := termbox.PollEvent()
			if event.Err != nil {
				err = event.Err
			} else {
				ch, key = event.Ch, event.Key
			}
		} else {
			if kbCh, kbKey, kbErr := keyboard.GetKey(); kbErr != nil {
				err = kbErr
			} else {
				ch, key = kbCh, termbox.Key(kbKey)
			}
		}
		if err != nil {
			chErrors <- err
			return
		}

		p.debugLog(">> %#v, %v ++ ch='%c', key=%#v, escapeSeq=%#v", buffer.String(), buffer.position, ch, key, p.escapeSeq)
		if ch == 0 || p.escapeSeq != "" {
			p.handleKeySequence(buffer, chErrors, ch, key)
		} else if ch != 0 {
			buffer.Insert(ch)
		}
		p.debugLog("   %#v, %v -- ch='%c', key=%#v, escapeSeq=%#v", buffer.String(), buffer.position, ch, key, p.escapeSeq)

		if buffer.IsDone() {
			return
		}
	}
}

func (p *prompt) handleKeySequence(buffer *buffer, chErrors chan error, ch rune, termboxKey termbox.Key) {
	var key KeySequence
	if p.escapeSeq != "" {
		p.escapeSeq += fmt.Sprintf("%c", ch)
		key = termboxEscKeySequenceMap[p.escapeSeq]
	} else {
		key = termboxKeyKeySequenceMap[termboxKey]
	}

	action := p.keyMapReversed.Insert[key]
	if key != "" && action != "" {
		p.debugLog("  ~ key=%v, action=%v", key, action)
	}
	switch action {
	case None:
		switch key {
		case Escape:
			p.escapeSeq = "\x1b"
			// Escape key press can result in just a \x1b with nothing following
			// it; a hacky way of handling it is to wait for a few milliseconds
			// and if no trailing characters are found, reset the escapeSeq and
			// treat it as an Escape KeySequence
			go func() {
				time.Sleep(EscapeSequenceDelay)
				if p.escapeSeq == "\x1b" {
					p.escapeSeq = ""
				}
			}()
		case Space:
			buffer.Insert(' ')
		case Tab:
			buffer.Insert('\t')
		}
	case Abort:
		chErrors <- ErrAbort
		return
	case DeleteCharCurrent:
		buffer.DeleteForward(1)
	case DeleteCharPrevious:
		buffer.DeleteBackward(1)
	case DeleteWordNext:
		buffer.DeleteWordForward()
	case DeleteWordPrevious:
		buffer.DeleteWordBackward()
	case EraseEverything:
		buffer.Clear()
	case EraseToBeginningOfLine:
		buffer.DeleteBackward(-1)
	case EraseToEndOfLine:
		buffer.DeleteForward(-1)
	case HistoryDown:
		buffer.Set(p.getHistoryItem(1))
	case HistoryUp:
		buffer.Set(p.getHistoryItem(-1))
	case MakeWordCapitalCase:
		buffer.MakeWordCapitalCase()
	case MakeWordLowerCase:
		buffer.MakeWordLowerCase()
	case MakeWordUpperCase:
		buffer.MakeWordUpperCase()
	case MoveDownOneLine:
		buffer.MoveDown(1)
	case MoveLeftOneCharacter:
		buffer.MoveLeft(1)
	case MoveRightOneCharacter:
		buffer.MoveRight(1)
	case MoveToBeginning:
		buffer.MoveLeft(-1)
	case MoveToBeginningOfLine:
		buffer.MoveLineBegin()
	case MoveToEnd:
		buffer.MoveRight(-1)
	case MoveToEndOfLine:
		buffer.MoveLineEnd()
	case MoveToWordNext:
		buffer.MoveWordRight()
	case MoveToWordPrevious:
		buffer.MoveWordLeft()
	case MoveUpOneLine:
		buffer.MoveUp(1)
	case Terminate:
		if p.terminationChecker == nil || p.terminationChecker(buffer.String()) {
			buffer.MarkAsDone()
			return
		}
		buffer.Insert('\n')
	}

	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		p.escapeSeq = ""
	}
}
