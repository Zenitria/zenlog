package zenlog

var (
	// defaultLogger is the default logger instance
	defaultLogger = DefaultLogger()
)

// Debug prints a debug log with the given message.
func Debug(msg string) {
	defaultLogger.Debug(msg)
}

// Debugf prints a debug log with the given format message.
func Debugf(format string, a ...interface{}) {
	defaultLogger.Debugf(format, a...)
}

// Success prints a success log with the given message.
func Success(msg string) {
	defaultLogger.Success(msg)
}

// Successf prints a success log with the given format message.
func Successf(format string, a ...interface{}) {
	defaultLogger.Successf(format, a...)
}

// Info prints a info log with the given message.
func Info(msg string) {
	defaultLogger.Info(msg)
}

// Infof prints a info log with the given format message.
func Infof(format string, a ...interface{}) {
	defaultLogger.Infof(format, a...)
}

// Warn prints a warn log with the given message.
func Warn(msg string) {
	defaultLogger.Warn(msg)
}

// Warnf prints a warn log with the given format message.
func Warnf(format string, a ...interface{}) {
	defaultLogger.Warnf(format, a...)
}

// Error prints a error log with the given message.
func Error(msg string) {
	defaultLogger.Error(msg)
}

// Errorf prints a error log with the given format message.
func Errorf(format string, a ...interface{}) {
	defaultLogger.Errorf(format, a...)
}

// Fatal prints a fatal log with the given message.
func Fatal(msg string) {
	defaultLogger.Fatal(msg)
}

// Fatalf prints a fatal log with the given format message.
func Fatalf(format string, a ...interface{}) {
	defaultLogger.Fatalf(format, a...)
}

// SetTimeFormat sets the time format of the logger
func SetTimeFormat(format string) {
	defaultLogger.SetTimeFormat(format)
}

// SetReportCaller sets the report caller of the logger
func SetReportCaller(report bool) {
	defaultLogger.SetReportCaller(report)
}

// SetDebug sets the debug mode of the logger
func SetDebug(debug bool) {
	defaultLogger.SetDebug(debug)
}
