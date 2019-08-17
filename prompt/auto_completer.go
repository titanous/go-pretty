package prompt

import (
	"sort"
	"strings"
)

// Suggestion is what is returned by the auto-completer.
type Suggestion struct {
	Value string
	Hint  string
}

// AutoCompleter defines a function that takes the entire user input, the word
// the user is specifically on, and the location of the cursor on the entire
// sentence. It is expected to return zero or more strings that match what the
// user may type.
type AutoCompleter func(sentence string, word string, location uint) []Suggestion

// AutoCompleteGoLangKeywords is a simple auto-completer that helps
// auto-complete most of the known GoLang keywords.
func AutoCompleteGoLangKeywords() AutoCompleter {
	keywords := []string{
		"append", "bool", "break", "byte", "cap", "case", "chan", "close", "complex", "complex128", "complex64", "const", "continue", "copy", "default", "defer", "delete", "else", "error", "fallthrough", "float", "float32", "float64", "for", "func", "go", "goto", "if", "import", "int", "int16", "int32", "int64", "int8", "interface", "len", "make", "map", "new", "package", "panic", "print", "println", "range", "real", "recover", "return", "rune", "select", "string", "struct", "switch", "type", "uint", "uint16", "uint32", "uint64", "uint8", "uintptr", "var",
	}
	return AutoCompleteWords(keywords, 3, false)
}

// AutoCompletePythonKeywords is a simple auto-completer that helps
// auto-complete most of the known Python keywords.
func AutoCompletePythonKeywords() AutoCompleter {
	keywords := []string{
		"ArithmeticError", "AssertionError", "AttributeError", "BaseException", "DeprecationWarning", "EOFError", "EnvironmentError", "Exception", "FloatingPointError", "FutureWarning", "GeneratorExit", "IOError", "ImportError", "ImportWarning", "IndentationError", "IndexError", "KeyError", "KeyboardInterrupt", "LookupError", "MemoryError", "NameError", "NotImplemented", "NotImplementedError", "OSError", "OverflowError", "OverflowWarning", "PendingDeprecationWarning", "ReferenceError", "RuntimeError", "RuntimeWarning", "StandardError", "StopIteration", "SyntaxError", "SyntaxWarning", "SystemError", "SystemExit", "TabError", "TypeError", "UnboundLocalError", "UnicodeDecodeError", "UnicodeEncodeError", "UnicodeError", "UnicodeTranslateError", "UnicodeWarning", "UserWarning", "VMSError", "ValueError", "Warning", "WindowsError", "ZeroDivisionError", "__import__", "abs", "all", "any", "apply", "as", "assert", "basestring", "bin", "bool", "break", "buffer", "bytearray", "bytes", "callable", "chr", "classmethod", "cmp", "coerce", "compile", "complex", "continue", "del", "delattr", "dict", "dir", "divmod", "elif", "else", "enumerate", "eval", "except", "exec", "execfile", "exit", "file", "filter", "finally", "float", "for", "from", "frozenset", "getattr", "global", "globals", "hasattr", "hash", "hex", "id", "if", "input", "int", "intern", "isinstance", "issubclass", "iter", "lambda", "len", "list", "locals", "long", "map", "max", "min", "next", "object", "oct", "open", "ord", "pass", "pow", "print", "property", "raise", "range", "raw_input", "reduce", "reload", "repr", "return", "reversed", "round", "set", "setattr", "slice", "sorted", "staticmethod", "str", "sum", "super", "try", "tuple", "type", "unichr", "unicode", "vars", "while", "with", "xrange", "yield", "zip",
	}
	return AutoCompleteWords(keywords, 3, false)
}

// AutoCompleteSQLKeywords is a simple auto-completer that helps
// auto-complete most of the known SQL keywords.
func AutoCompleteSQLKeywords() AutoCompleter {
	keywords := []string{
		"ABORT", "ABS", "ABSOLUTE", "ACCESS", "ADA", "ADD", "ADMIN", "AFTER", "AGGREGATE", "ALIAS", "ALL", "ALLOCATE", "ALTER", "ANALYSE", "ANALYZE", "AND", "ANY", "ARE", "AS", "ASC", "ASENSITIVE", "ASSERTION", "ASSIGNMENT", "ASYMMETRIC", "AT", "ATOMIC", "AUTHORIZATION", "AVG", "BACKWARD", "BEFORE", "BEGIN", "BETWEEN", "BITVAR", "BIT_LENGTH", "BOTH", "BREADTH", "BY", "C", "CACHE", "CALL", "CALLED", "CARDINALITY", "CASCADE", "CASCADED", "CASE", "CAST", "CATALOG", "CATALOG_NAME", "CHAIN", "CHARACTERISTICS", "CHARACTER_LENGTH", "CHARACTER_SET_CATALOG", "CHARACTER_SET_NAME", "CHARACTER_SET_SCHEMA", "CHAR_LENGTH", "CHECK", "CHECKED", "CHECKPOINT", "CLASS", "CLASS_ORIGIN", "CLOB", "CLOSE", "CLUSTER", "COALSECE", "COBOL", "COLLATE", "COLLATION", "COLLATION_CATALOG", "COLLATION_NAME", "COLLATION_SCHEMA", "COLUMN", "COLUMN_NAME", "COMMAND_FUNCTION", "COMMAND_FUNCTION_CODE", "COMMENT", "COMMIT", "COMMITTED", "COMPLETION", "CONDITION_NUMBER", "CONNECT", "CONNECTION", "CONNECTION_NAME", "CONSTRAINT", "CONSTRAINTS", "CONSTRAINT_CATALOG", "CONSTRAINT_NAME", "CONSTRAINT_SCHEMA", "CONSTRUCTOR", "CONTAINS", "CONTINUE", "CONVERSION", "CONVERT", "COPY", "CORRESPONTING", "COUNT", "CREATE", "CREATEDB", "CREATEUSER", "CROSS", "CUBE", "CURRENT", "CURRENT_DATE", "CURRENT_PATH", "CURRENT_ROLE", "CURRENT_TIME", "CURRENT_TIMESTAMP", "CURRENT_USER", "CURSOR", "CURSOR_NAME", "CYCLE", "DATA", "DATABASE", "DATETIME_INTERVAL_CODE", "DATETIME_INTERVAL_PRECISION", "DAY", "DEALLOCATE", "DECLARE", "DEFAULT", "DEFAULTS", "DEFERRABLE", "DEFERRED", "DEFINED", "DEFINER", "DELETE", "DELIMITER", "DELIMITERS", "DEREF", "DESC", "DESCRIBE", "DESCRIPTOR", "DESTROY", "DESTRUCTOR", "DETERMINISTIC", "DIAGNOSTICS", "DICTIONARY", "DISCONNECT", "DISPATCH", "DISTINCT", "DO", "DOMAIN", "DROP", "DYNAMIC", "DYNAMIC_FUNCTION", "DYNAMIC_FUNCTION_CODE", "EACH", "ELSE", "ELSIF", "ENCODING", "ENCRYPTED", "END", "END-EXEC", "EQUALS", "ESCAPE", "EVERY", "EXCEPTION", "EXCEPT", "EXCLUDING", "EXCLUSIVE", "EXEC", "EXECUTE", "EXISTING", "EXISTS", "EXPLAIN", "EXTERNAL", "EXTRACT", "FALSE", "FETCH", "FINAL", "FIRST", "FOR", "FORCE", "FOREIGN", "FORTRAN", "FORWARD", "FOUND", "FREE", "FREEZE", "FROM", "FULL", "FUNCTION", "G", "GENERAL", "GENERATED", "GET", "GLOBAL", "GO", "GOTO", "GRANT", "GRANTED", "GROUP", "GROUPING", "HANDLER", "HAVING", "HIERARCHY", "HOLD", "HOST", "IDENTITY", "IF", "IGNORE", "ILIKE", "IMMEDIATE", "IMMUTABLE", "IMPLEMENTATION", "IMPLICIT", "IN", "INCLUDING", "INCREMENT", "INDEX", "INDITCATOR", "INFIX", "INHERITS", "INITIALIZE", "INITIALLY", "INNER", "INOUT", "INPUT", "INSENSITIVE", "INSERT", "INSTANTIABLE", "INSTEAD", "INTERSECT", "INTO", "INVOKER", "IS", "ISNULL", "ISOLATION", "ITERATE", "JOIN", "KEY", "KEY_MEMBER", "KEY_TYPE", "LANCOMPILER", "LANGUAGE", "LARGE", "LAST", "LATERAL", "LEADING", "LEFT", "LENGTH", "LESS", "LEVEL", "LIKE", "LIMIT", "LISTEN", "LOAD", "LOCAL", "LOCALTIME", "LOCALTIMESTAMP", "LOCATION", "LOCATOR", "LOCK", "LOWER", "MAP", "MATCH", "MAX", "MAXVALUE", "MESSAGE_LENGTH", "MESSAGE_OCTET_LENGTH", "MESSAGE_TEXT", "METHOD", "MIN", "MINUTE", "MINVALUE", "MOD", "MODE", "MODIFIES", "MODIFY", "MONTH", "MORE", "MOVE", "MUMPS", "NAMES", "NATIONAL", "NATURAL", "NCHAR", "NCLOB", "NEW", "NEXT", "NO", "NOCREATEDB", "NOCREATEUSER", "NONE", "NOT", "NOTHING", "NOTIFY", "NOTNULL", "NULL", "NULLABLE", "NULLIF", "OBJECT", "OCTET_LENGTH", "OF", "OFF", "OFFSET", "OIDS", "OLD", "ON", "ONLY", "OPEN", "OPERATION", "OPERATOR", "OPTION", "OPTIONS", "OR", "ORDER", "ORDINALITY", "OUT", "OUTER", "OUTPUT", "OVERLAPS", "OVERLAY", "OVERRIDING", "OWNER", "PAD", "PARAMETER", "PARAMETERS", "PARAMETER_MODE", "PARAMATER_NAME", "PARAMATER_ORDINAL_POSITION", "PARAMETER_SPECIFIC_CATALOG", "PARAMETER_SPECIFIC_NAME", "PARAMATER_SPECIFIC_SCHEMA", "PARTIAL", "PASCAL", "PENDANT", "PLACING", "PLI", "POSITION", "POSTFIX", "PRECISION", "PREFIX", "PREORDER", "PREPARE", "PRESERVE", "PRIMARY", "PRIOR", "PRIVILEGES", "PROCEDURAL", "PROCEDURE", "PUBLIC", "READ", "READS", "RECHECK", "RECURSIVE", "REF", "REFERENCES", "REFERENCING", "REINDEX", "RELATIVE", "RENAME", "REPEATABLE", "REPLACE", "RESET", "RESTART", "RESTRICT", "RESULT", "RETURN", "RETURNED_LENGTH", "RETURNED_OCTET_LENGTH", "RETURNED_SQLSTATE", "RETURNS", "REVOKE", "RIGHT", "ROLE", "ROLLBACK", "ROLLUP", "ROUTINE", "ROUTINE_CATALOG", "ROUTINE_NAME", "ROUTINE_SCHEMA", "ROW", "ROWS", "ROW_COUNT", "RULE", "SAVE_POINT", "SCALE", "SCHEMA", "SCHEMA_NAME", "SCOPE", "SCROLL", "SEARCH", "SECOND", "SECURITY", "SELECT", "SELF", "SENSITIVE", "SERIALIZABLE", "SERVER_NAME", "SESSION", "SESSION_USER", "SET", "SETOF", "SETS", "SHARE", "SHOW", "SIMILAR", "SIMPLE", "SIZE", "SOME", "SOURCE", "SPACE", "SPECIFIC", "SPECIFICTYPE", "SPECIFIC_NAME", "SQL", "SQLCODE", "SQLERROR", "SQLEXCEPTION", "SQLSTATE", "SQLWARNINIG", "STABLE", "START", "STATE", "STATEMENT", "STATIC", "STATISTICS", "STDIN", "STDOUT", "STORAGE", "STRICT", "STRUCTURE", "STYPE", "SUBCLASS_ORIGIN", "SUBLIST", "SUBSTRING", "SUM", "SYMMETRIC", "SYSID", "SYSTEM", "SYSTEM_USER", "TABLE", "TABLE_NAME", " TEMP", "TEMPLATE", "TEMPORARY", "TERMINATE", "THAN", "THEN", "TIMESTAMP", "TIMEZONE_HOUR", "TIMEZONE_MINUTE", "TO", "TOAST", "TRAILING", "TRANSATION", "TRANSACTIONS_COMMITTED", "TRANSACTIONS_ROLLED_BACK", "TRANSATION_ACTIVE", "TRANSFORM", "TRANSFORMS", "TRANSLATE", "TRANSLATION", "TREAT", "TRIGGER", "TRIGGER_CATALOG", "TRIGGER_NAME", "TRIGGER_SCHEMA", "TRIM", "TRUE", "TRUNCATE", "TRUSTED", "TYPE", "UNCOMMITTED", "UNDER", "UNENCRYPTED", "UNION", "UNIQUE", "UNKNOWN", "UNLISTEN", "UNNAMED", "UNNEST", "UNTIL", "UPDATE", "UPPER", "USAGE", "USER", "USER_DEFINED_TYPE_CATALOG", "USER_DEFINED_TYPE_NAME", "USER_DEFINED_TYPE_SCHEMA", "USING", "VACUUM", "VALID", "VALIDATOR", "VALUES", "VARIABLE", "VERBOSE", "VERSION", "VIEW", "VOLATILE", "WHEN", "WHENEVER", "WHERE", "WITH", "WITHOUT", "WORK", "WRITE", "YEAR", "ZONE",
	}
	return AutoCompleteWords(keywords, 3, true)
}

// AutoCompleteWords is an auto-completer that finds all words that starts
// with the user's word in a pre-defined list of words.
func AutoCompleteWords(possibleMatches []string, minChars int, caseInsensitive bool) AutoCompleter {
	// sort before-hand and avoid sorting while searching in the loop below
	sort.Strings(possibleMatches)

	return func(sentence string, word string, location uint) []Suggestion {
		var matches []Suggestion
		if len(word) >= minChars {
			if caseInsensitive {
				word = strings.ToLower(word)
			}
			for _, possibleMatch := range possibleMatches {
				if caseInsensitive {
					if strings.HasPrefix(strings.ToLower(possibleMatch), word) && len(possibleMatch) > len(word) {
						matches = append(matches, Suggestion{Value: possibleMatch})
					}
				} else {
					if strings.HasPrefix(possibleMatch, word) && len(possibleMatch) > len(word) {
						matches = append(matches, Suggestion{Value: possibleMatch})
					}
				}
			}
		}
		return matches
	}
}
