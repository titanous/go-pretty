package prompt

import "github.com/nsf/termbox-go"

// KeySequence defines a key-sequence that the user presses.
type KeySequence string

// KeySequences are a slice of KeySequence(s).
type KeySequences []KeySequence

// Support Keys
const (
	AltA           KeySequence = "Alt+A"
	AltB           KeySequence = "Alt+B"
	AltC           KeySequence = "Alt+C"
	AltD           KeySequence = "Alt+D"
	AltE           KeySequence = "Alt+E"
	AltF           KeySequence = "Alt+F"
	AltG           KeySequence = "Alt+G"
	AltH           KeySequence = "Alt+H"
	AltI           KeySequence = "Alt+I"
	AltJ           KeySequence = "Alt+J"
	AltK           KeySequence = "Alt+K"
	AltL           KeySequence = "Alt+L"
	AltM           KeySequence = "Alt+M"
	AltN           KeySequence = "Alt+N"
	AltO           KeySequence = "Alt+O"
	AltP           KeySequence = "Alt+P"
	AltQ           KeySequence = "Alt+Q"
	AltR           KeySequence = "Alt+R"
	AltS           KeySequence = "Alt+S"
	AltT           KeySequence = "Alt+T"
	AltU           KeySequence = "Alt+U"
	AltV           KeySequence = "Alt+V"
	AltW           KeySequence = "Alt+W"
	AltX           KeySequence = "Alt+X"
	AltY           KeySequence = "Alt+Y"
	AltZ           KeySequence = "Alt+Z"
	ArrowDown      KeySequence = "ArrowDown"
	ArrowLeft      KeySequence = "ArrowLeft"
	ArrowRight     KeySequence = "ArrowRight"
	ArrowUp        KeySequence = "ArrowUp"
	Backspace      KeySequence = "Backspace"
	CtrlA          KeySequence = "Ctrl+A"
	CtrlB          KeySequence = "Ctrl+B"
	CtrlC          KeySequence = "Ctrl+C"
	CtrlD          KeySequence = "Ctrl+D"
	CtrlE          KeySequence = "Ctrl+E"
	CtrlF          KeySequence = "Ctrl+F"
	CtrlG          KeySequence = "Ctrl+G"
	CtrlH          KeySequence = "Ctrl+H"
	CtrlI          KeySequence = "Ctrl+I"
	CtrlJ          KeySequence = "Ctrl+J"
	CtrlK          KeySequence = "Ctrl+K"
	CtrlL          KeySequence = "Ctrl+L"
	CtrlM          KeySequence = "Ctrl+M"
	CtrlN          KeySequence = "Ctrl+N"
	CtrlO          KeySequence = "Ctrl+O"
	CtrlP          KeySequence = "Ctrl+P"
	CtrlQ          KeySequence = "Ctrl+Q"
	CtrlR          KeySequence = "Ctrl+R"
	CtrlS          KeySequence = "Ctrl+S"
	CtrlT          KeySequence = "Ctrl+T"
	CtrlU          KeySequence = "Ctrl+U"
	CtrlV          KeySequence = "Ctrl+V"
	CtrlW          KeySequence = "Ctrl+W"
	CtrlX          KeySequence = "Ctrl+X"
	CtrlY          KeySequence = "Ctrl+Y"
	CtrlZ          KeySequence = "Ctrl+Z"
	CtrlArrowDown  KeySequence = "Ctrl+ArrowDown"
	CtrlArrowLeft  KeySequence = "Ctrl+ArrowLeft"
	CtrlArrowRight KeySequence = "Ctrl+ArrowRight"
	CtrlArrowUp    KeySequence = "Ctrl+ArrowUp"
	CtrlEnd        KeySequence = "Ctrl+End"
	CtrlHome       KeySequence = "Ctrl+Home"
	Delete         KeySequence = "Delete"
	End            KeySequence = "End"
	Enter          KeySequence = "Enter"
	Home           KeySequence = "Home"
	PageDown       KeySequence = "PageDown"
	PageUp         KeySequence = "PageUp"
	Space          KeySequence = "Space"
	Tab            KeySequence = "Tab"

	// following key sequences are not supported externally and are here only
	// for internal handling
	escape KeySequence = "escape"
)

var (
	// termboxKeyKeySequenceMap maps all supported termbox.Key types to an
	// equivalent KeySequence object
	termboxKeyKeySequenceMap = map[termbox.Key]KeySequence{
		termbox.KeyArrowDown:  ArrowDown,
		termbox.KeyArrowLeft:  ArrowLeft,
		termbox.KeyArrowRight: ArrowRight,
		termbox.KeyArrowUp:    ArrowUp,
		termbox.KeyBackspace:  Backspace, // same as termbox.KeyCtrlH
		termbox.KeyBackspace2: Backspace,
		termbox.KeyCtrlA:      CtrlA,
		termbox.KeyCtrlB:      CtrlB,
		termbox.KeyCtrlC:      CtrlC,
		termbox.KeyCtrlD:      CtrlD,
		termbox.KeyCtrlE:      CtrlE,
		termbox.KeyCtrlF:      CtrlF,
		termbox.KeyCtrlG:      CtrlG,
		termbox.KeyCtrlJ:      CtrlJ,
		termbox.KeyCtrlK:      CtrlK,
		termbox.KeyCtrlL:      CtrlL,
		termbox.KeyCtrlN:      CtrlN,
		termbox.KeyCtrlO:      CtrlO,
		termbox.KeyCtrlP:      CtrlP,
		termbox.KeyCtrlQ:      CtrlQ,
		termbox.KeyCtrlR:      CtrlR,
		termbox.KeyCtrlS:      CtrlS,
		termbox.KeyCtrlT:      CtrlT,
		termbox.KeyCtrlU:      CtrlU,
		termbox.KeyCtrlV:      CtrlV,
		termbox.KeyCtrlW:      CtrlW,
		termbox.KeyCtrlX:      CtrlX,
		termbox.KeyCtrlY:      CtrlY,
		termbox.KeyCtrlZ:      CtrlZ,
		termbox.KeyDelete:     Delete,
		termbox.KeyEnter:      Enter, // same as termbox.KeyCtrlM
		termbox.KeyEnd:        End,
		termbox.KeyEsc:        escape,
		termbox.KeyHome:       Home,
		termbox.KeyPgdn:       PageDown,
		termbox.KeyPgup:       PageUp,
		termbox.KeySpace:      Space,
		termbox.KeyTab:        Tab, // same as termbox.KeyCtrlI
	}

	// termboxEscKeySequenceMap maps all supported termbox detected escape
	// Sequences to an equivalent KeySequence object. This may not work on all
	// OSes/terminals and needs further research/work.
	termboxEscKeySequenceMap = map[string]KeySequence{
		"\x1ba":     AltA,
		"\x1bb":     AltB,
		"\x1bc":     AltC,
		"\x1bd":     AltD,
		"\x1be":     AltE,
		"\x1bf":     AltF,
		"\x1bg":     AltG,
		"\x1bh":     AltH,
		"\x1bi":     AltI,
		"\x1bj":     AltJ,
		"\x1bk":     AltK,
		"\x1bl":     AltL,
		"\x1bm":     AltM,
		"\x1bn":     AltN,
		"\x1bo":     AltO,
		"\x1bp":     AltP,
		"\x1bq":     AltQ,
		"\x1br":     AltR,
		"\x1bs":     AltS,
		"\x1bt":     AltT,
		"\x1bu":     AltU,
		"\x1bv":     AltV,
		"\x1bw":     AltW,
		"\x1bx":     AltX,
		"\x1by":     AltY,
		"\x1bz":     AltZ,
		"\x1b[1;5A": CtrlArrowUp,
		"\x1b[1;5B": CtrlArrowDown,
		"\x1b[1;5C": CtrlArrowRight,
		"\x1b[1;5D": CtrlArrowLeft,
		"\x1b[1;5F": CtrlEnd,
		"\x1b[1;5H": CtrlHome,
	}
)
