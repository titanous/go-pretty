package prompt

import "strings"

// TerminationChecker defines a function that tells if the user has typed all
// that is needed in the prompt and it is okay to return control to caller.
type TerminationChecker func(s string) bool

// TerminationCheckerSQL is a TerminationChecker that works for SQL prompts.
func TerminationCheckerSQL(s string) bool {
	// SQLs end with a ';'
	if strings.HasSuffix(s, ";") {
		return true
	}
	// SQL command can begin with a '/' (avoid cases with comment /*comment*/)
	if strings.HasPrefix(s, "/") && !strings.HasPrefix(s, "/*") {
		return true
	}
	return false
}
