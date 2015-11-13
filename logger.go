// Copyright (C) 2014 Jakob Borg. All rights reserved. Use of this source code
// is governed by an MIT-style license that can be found in the LICENSE file.

// Package logger implements a standardized logger with callback functionality
package logger

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

type Logger interface {
	SetLevel(LogLevel)
	Level() LogLevel

	Debugln(...interface{})
	Debugf(string, ...interface{})

	Verboseln(...interface{})
	Verbosef(string, ...interface{})

	Infoln(...interface{})
	Infof(string, ...interface{})

	Okln(...interface{})
	Okf(string, ...interface{})

	Warnln(...interface{})
	Warnf(string, ...interface{})

	Fatalln(...interface{})
	Fatalf(string, ...interface{})
}

// The default logger logs to standard output with a time prefix.
var DefaultLogger Logger = NewStdlibLogger()
