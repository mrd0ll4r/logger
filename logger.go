// Copyright (C) 2014-2015 Jakob Borg & mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

// Package logger contains a level-based logger interface and an implementation of that interface using the stdlib log
package logger

// LogLevel is a level of logging
type LogLevel int

// The log levels
const (
	Everything LogLevel = iota // can be used to log everything
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelFatal
	Off // can be used to log nothig, even if levels are added or removed in the future
)

// Levels returns a slice of all Levels (excluding Off and Everything) ordered from fine to severe (Debug to Fatal)
func Levels() []LogLevel {
	return []LogLevel{LevelTrace, LevelDebug, LevelInfo, LevelWarn, LevelFatal}
}

// Logger is an interface for leveled loggers.
// All logging methods must be thread-safe, no guarantee is provided regarding the other methods.
type Logger interface {
	// SetLevel instructs the logger to log everything above and including the provided level
	SetLevel(LogLevel)
	// Logs returns whether this logger logs the specified level
	Logs(LogLevel) bool

	// Traceln logs a trace-marked line
	Traceln(val ...interface{})
	// Tracef logs a trace-marked formatted line
	Tracef(format string, val ...interface{})

	// Debugln logs a debug-marked line
	Debugln(val ...interface{})
	// Debugf logs a debug-marked formatted line
	Debugf(format string, val ...interface{})

	// Infoln logs an info-marked line
	Infoln(val ...interface{})
	// Infof logs an info-marked formatted line
	Infof(format string, val ...interface{})

	// Warnln logs a warn-marked line
	Warnln(val ...interface{})
	// Warnf logs a warn-marked formatted line
	Warnf(format string, val ...interface{})

	// Fatalln logs a fatal-marked line and exits the program with return code 1
	Fatalln(val ...interface{})
	// Fatalf logs a fatal-marked formatted line and exits the program with return code 1
	Fatalf(format string, val ...interface{})
}

// defaultLogger logs to standard output with a time prefix, see NewStdlibLogger()
// If you want to utilize the default functions (i.e. logger.Infoln("...")), with the StdlibLogger and log.Llongfile or
// log.Lshortfile, you'll have to change the calldepth on the StdlibLogger, see SetCalldepthForDefault()
var defaultLogger Logger = NewStdlibLogger()

// SetDefaultLogger sets the default logger
func SetDefaultLogger(l Logger) {
	defaultLogger = l
}
