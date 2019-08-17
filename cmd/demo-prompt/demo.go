package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/prompt"
	"github.com/jedib0t/go-pretty/text"
	"github.com/nsf/termbox-go"
)

var (
	flagDebug         = flag.Bool("debug", false, "Write debug logs to prompt.log?")
	flagTestNoTermBox = flag.Bool("no-termbox", false, "Do not use TermBox?")
)

func testTermBoxPolling() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	count := 0
	fmt.Printf("Polling for events... (press Ctrl+D to exit)\n")
	for {
		event := termbox.PollEvent()
		count++
		fmt.Printf(">> %#v [ch: %c]\n", event, event.Ch)
		if event.Key == termbox.KeyCtrlD {
			break
		}
	}
	termbox.Close()

	fmt.Printf("Polling for events... done! [found %d events]\n", count)

}

func main() {
	flag.Parse()

	colorFaint := text.FgHiBlack
	colorTitle := text.Colors{text.FgHiCyan, text.Bold}
	count := 0
	errStop := errors.New("demo.stop")
	inputHandler := prompt.Callback(func(userInput string, err error) error {
		// handle errors and user exit attempts
		if err != nil {
			if err == prompt.ErrAbort {
				return errStop
			}
			return err
		} else if userInput == "/quit" {
			return errStop
		}

		// deal with user input: to keep it a simple demo, just print it out
		count++
		fmt.Println(colorFaint.Sprintf("userInput: %#v", userInput))
		return nil
	})
	timeStart := time.Now()
	title := "" +
		colorTitle.Sprint("*****************************************************************") + "\n" +
		colorTitle.Sprint("*** SQL prompt powered by github.com/jedib0t/go-pretty/prompt ***") + "\n" +
		colorTitle.Sprint("*****************************************************************") + "\n" +
		colorFaint.Sprint("(press Ctrl+D or type '/quit' and press Enter to terminate)")

	p := prompt.NewWriter()
	if *flagDebug {
		p.Debug()
	}
	if *flagTestNoTermBox {
		p.DoNotUseTermbox()
	}
	p.SetAutoCompleter(prompt.AutoCompleteSQLKeywords())
	p.SetHistory([]string{
		"SELECT * FROM employees WHERE id = 1;",
		"SELECT * FROM employees WHERE\n\tid BETWEEN (1, 101) AND\n\temployees.salary > 13000;",
	})
	p.SetStyle(prompt.StyleColored)
	p.SetTerminationChecker(prompt.TerminationCheckerSQL)
	p.SetTitle(title)
	p.Style().SyntaxHighlighter = prompt.SyntaxHighlighterOptionsSQL
	defer func() {
		fmt.Printf(colorFaint.Sprintf(
			"Bye! (evaluated %d command(s) in %s)\n", count, time.Since(timeStart).Round(time.Second),
		))
	}()

	if err := p.Render(inputHandler); err != nil {
		if err != errStop {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(-1)
		}
	}
}
