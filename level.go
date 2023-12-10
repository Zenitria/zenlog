package zenlog

import "github.com/charmbracelet/lipgloss"

type Level int

const (
	// DebugLevel is the debug level
	DebugLevel Level = 0x000
	// InfoLevel is the info level
	InfoLevel Level = 0x100
	// SuccessLevel is the success level
	SuccessLevel Level = 0x200
	// WarnLevel is the warn level
	WarnLevel Level = 0x300
	// ErrorLevel is the error level
	ErrorLevel Level = 0x400
	// FatalLevel is the fatal level
	FatalLevel Level = 0x500
)

// String returns the string representation of the level
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBU"
	case InfoLevel:
		return "INFO"
	case SuccessLevel:
		return "SUCC"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERRO"
	case FatalLevel:
		return "FATA"
	default:
		return "UNKN"
	}
}

// Color returns the color of the level
func (l Level) Color() lipgloss.Color {
	switch l {
	case DebugLevel:
		return lipgloss.Color("#FF4EAA")
	case InfoLevel:
		return lipgloss.Color("#0DB1FC")
	case SuccessLevel:
		return lipgloss.Color("#09E85E")
	case WarnLevel:
		return lipgloss.Color("#FFF712")
	case ErrorLevel:
		return lipgloss.Color("#ED424B")
	case FatalLevel:
		return lipgloss.Color("#BF47FF")
	default:
		return lipgloss.Color("#FFFFFF")
	}
}

// Style returns the style of the level
func (l Level) Style() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(l.Color()).Foreground(lipgloss.Color("#000000")).Bold(true).Padding(0, 1)
}
