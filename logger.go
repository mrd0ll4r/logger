// Copyright (C) 2014 Jakob Borg. All rights reserved. Use of this source code
// is governed by an MIT-style license that can be found in the LICENSE file.

// Package logger contains a level-based logger interface and an implementation of that interface using the stdlib log
package logger

// LogLevel is a level of logging
type LogLevel int

const (
	Off LogLevel = iota // can be used to log nothing
	LevelDebug
	LevelVerbose
	LevelInfo
	LevelOK
	LevelWarn
	LevelFatal
	NumLevels // can be used to log everything, even if levels are added or removed in the future
)

// Logger is an interface for leveled loggers.
// All logging methods must be thread-safe, no guarantee is provided regarding the other methods.
type Logger interface {
	// SetLevel instructs the logger to log everything above and including the provided level
	SetLevel(LogLevel)
	// Logs returns whether this logger logs the specified level
	Logs(LogLevel) bool

	// Debugln logs a debug-marked line
	Debugln(...interface{})
	// Debugf logs a debug-marked formatted line
	Debugf(string, ...interface{})

	// Verboseln logs a verbose-marked line
	Verboseln(...interface{})
	// Verbosef logs a verbose-marked formatted line
	Verbosef(string, ...interface{})

	// Infoln logs an info-marked line
	Infoln(...interface{})
	// Infof logs an info-marked formatted line
	Infof(string, ...interface{})

	// Okln logs an OK-marked line
	Okln(...interface{})
	// Okf logs an OK-marked formatted line
	Okf(string, ...interface{})

	// Warnln logs a warn-marked line
	Warnln(...interface{})
	// Warnf logs a warn-marked formatted line
	Warnf(string, ...interface{})

	// Fatalln logs a fatal-marked line and exits the program with return code 1
	Fatalln(...interface{})
	// Fatalf logs a fatal-marked formatted line and exits the program with return code 1
	Fatalf(string, ...interface{})
}

// The default logger logs to standard output with a time prefix, see NewStdlibLogger()
var DefaultLogger Logger = NewStdlibLogger()
