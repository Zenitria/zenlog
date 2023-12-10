package zenlog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type Logger struct {
	TimeFormat   string
	ReportCaller bool
	DebugMode    bool
}

// DefaultLogger returns a new default logger
func DefaultLogger() *Logger {
	return &Logger{
		TimeFormat:   "2006-01-02 15:04:05",
		ReportCaller: false,
		DebugMode:    false,
	}
}

// SetTimeFormat sets the time format
func (l *Logger) SetTimeFormat(format string) {
	l.TimeFormat = format
}

// SetReportCaller sets whether to report the caller
func (l *Logger) SetReportCaller(report bool) {
	l.ReportCaller = report
}

// SetDebug sets the debug mode of the logger
func (l *Logger) SetDebug(debug bool) {
	l.DebugMode = debug
}

// Log prints a log with the given level and message
func (l *Logger) Log(level Level, msg string) {
	now := time.Now().Format(l.TimeFormat)

	if l.ReportCaller {
		_, path, line, _ := runtime.Caller(2)
		_, module := filepath.Split(filepath.Dir(path))
		_, file := filepath.Split(path)

		reportStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#999999"))

		report := reportStyle.Render(fmt.Sprintf("<%s/%s:%d>", module, file, line))

		fmt.Printf("%s %s %s %s\n", now, level.Style().Render(level.String()), report, msg)
	} else {
		fmt.Printf("%s %s %s\n", now, level.Style().Render(level.String()), msg)
	}
}

// Logf prints a log with the given level and format message.
func (l *Logger) Logf(level Level, format string, a ...interface{}) {
	l.Log(level, fmt.Sprintf(format, a...))
}

// Debug prints a debug log with the given message.
func (l *Logger) Debug(msg string) {
	if l.DebugMode {
		l.Log(DebugLevel, msg)
	}
}

// Debugf prints a debug log with the given format message.
func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.DebugMode {
		l.Logf(DebugLevel, format, a...)
	}
}

// Success prints a success log with the given message.
func (l *Logger) Success(msg string) {
	l.Log(SuccessLevel, msg)
}

// Successf prints a success log with the given format message.
func (l *Logger) Successf(format string, a ...interface{}) {
	l.Logf(SuccessLevel, format, a...)
}

// Info prints a info log with the given message.
func (l *Logger) Info(msg string) {
	l.Log(InfoLevel, msg)
}

// Infof prints a info log with the given format message.
func (l *Logger) Infof(format string, a ...interface{}) {
	l.Logf(InfoLevel, format, a...)
}

// Warn prints a warn log with the given message.
func (l *Logger) Warn(msg string) {
	l.Log(WarnLevel, msg)
}

// Warnf prints a warn log with the given format message.
func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Logf(WarnLevel, format, a...)
}

// Error prints a error log with the given message.
func (l *Logger) Error(msg string) {
	l.Log(ErrorLevel, msg)
}

// Errorf prints a error log with the given format message.
func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Logf(ErrorLevel, format, a...)
}

// Fatal prints a fatal log with the given message.
func (l *Logger) Fatal(msg string) {
	l.Log(FatalLevel, msg)
	os.Exit(1)
}

// Fatalf prints a fatal log with the given format message.
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Logf(FatalLevel, format, a...)
	os.Exit(1)
}
