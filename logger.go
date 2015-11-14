// Package logger contains a level-based logger interface and an implementation of that interface using the stdlib log
package logger

// LogLevel is a level of logging
type LogLevel int

// The log levels
const (
	Everything LogLevel = iota // can be used to log everything
	LevelDebug
	LevelVerbose
	LevelInfo
	LevelOK
	LevelWarn
	LevelFatal
	Off // can be used to log nothig, even if levels are added or removed in the future
)

// Levels returns a slice of all Levels (excluding Off and NumLevels) ordered from fine to severe (Debug to Fatal)
func Levels() []LogLevel {
	return []LogLevel{LevelDebug, LevelVerbose, LevelInfo, LevelOK, LevelWarn, LevelFatal}
}

// Logger is an interface for leveled loggers.
// All logging methods must be thread-safe, no guarantee is provided regarding the other methods.
type Logger interface {
	// SetLevel instructs the logger to log everything above and including the provided level
	SetLevel(LogLevel)
	// Logs returns whether this logger logs the specified level
	Logs(LogLevel) bool

	// Debugln logs a debug-marked line
	Debugln(val ...interface{})
	// Debugf logs a debug-marked formatted line
	Debugf(format string, val ...interface{})

	// Verboseln logs a verbose-marked line
	Verboseln(val ...interface{})
	// Verbosef logs a verbose-marked formatted line
	Verbosef(format string, val ...interface{})

	// Infoln logs an info-marked line
	Infoln(val ...interface{})
	// Infof logs an info-marked formatted line
	Infof(format string, val ...interface{})

	// Okln logs an OK-marked line
	Okln(val ...interface{})
	// Okf logs an OK-marked formatted line
	Okf(format string, val ...interface{})

	// Warnln logs a warn-marked line
	Warnln(val ...interface{})
	// Warnf logs a warn-marked formatted line
	Warnf(format string, val ...interface{})

	// Fatalln logs a fatal-marked line and exits the program with return code 1
	Fatalln(val ...interface{})
	// Fatalf logs a fatal-marked formatted line and exits the program with return code 1
	Fatalf(format string, val ...interface{})
}

// DefaultLogger logs to standard output with a time prefix, see NewStdlibLogger()
var DefaultLogger Logger = NewStdlibLogger()
