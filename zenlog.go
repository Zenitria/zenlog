package zenlog

var (
	// defaultLogger is the default logger instance
	defaultLogger = DefaultLogger()
)

// Debug prints a debug log with the given message.
func Debug(msg string) {
	defaultLogger.Log(DebugLevel, msg)
}

// Debugf prints a debug log with the given format message.
func Debugf(format string, a ...interface{}) {
	defaultLogger.Logf(DebugLevel, format, a...)
}

// Success prints a success log with the given message.
func Success(msg string) {
	defaultLogger.Log(SuccessLevel, msg)
}

// Successf prints a success log with the given format message.
func Successf(format string, a ...interface{}) {
	defaultLogger.Logf(SuccessLevel, format, a...)
}

// Info prints a info log with the given message.
func Info(msg string) {
	defaultLogger.Log(InfoLevel, msg)
}

// Infof prints a info log with the given format message.
func Infof(format string, a ...interface{}) {
	defaultLogger.Logf(InfoLevel, format, a...)
}

// Warn prints a warn log with the given message.
func Warn(msg string) {
	defaultLogger.Log(WarnLevel, msg)
}

// Warnf prints a warn log with the given format message.
func Warnf(format string, a ...interface{}) {
	defaultLogger.Logf(WarnLevel, format, a...)
}

// Error prints a error log with the given message.
func Error(msg string) {
	defaultLogger.Log(ErrorLevel, msg)
}

// Errorf prints a error log with the given format message.
func Errorf(format string, a ...interface{}) {
	defaultLogger.Logf(ErrorLevel, format, a...)
}

// Fatal prints a fatal log with the given message.
func Fatal(msg string) {
	defaultLogger.Log(FatalLevel, msg)
}

// Fatalf prints a fatal log with the given format message.
func Fatalf(format string, a ...interface{}) {
	defaultLogger.Logf(FatalLevel, format, a...)
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
