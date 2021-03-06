// Most of this file is taken from github.com/calmh/logger and therefore:
// Copyright (C) 2014,2015 Jakob Borg & mrd0ll4r. All rights reserved. Use of this source code
// is governed by an MIT-style license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// StdlibLogger is an implementation of the Logger interface that uses the stdlibs loggger for logging.
type StdlibLogger struct {
	logger    *log.Logger
	level     LogLevel
	calldepth int
}

// to see if we implement Logger
var _ Logger = NewStdlibLogger()

// NewStdlibLogger returns a new StdlibLogger that logs to Stdout with a time prefix and LevelInfo logging.
// If the environment variable LOGGER_DISCARD is set, the logger will discard everything (useful for benchmarks)
func NewStdlibLogger() *StdlibLogger {
	if os.Getenv("LOGGER_DISCARD") != "" {
		// Hack to completely disable logging, for example when running benchmarks.
		return &StdlibLogger{
			logger:    log.New(ioutil.Discard, "", 0),
			level:     Off,
			calldepth: 2,
		}
	}

	return &StdlibLogger{
		logger:    log.New(os.Stdout, "", log.Ltime),
		level:     LevelInfo,
		calldepth: 2,
	}
}

// SetCalldepthForDefault sets the calldepth to use for file location lookup.
// Changing this is only ever required when using (log.Llongfile or log.Lshortfile) and the default functions!
// If you set this, using the logger directly will be impossible.
func (l *StdlibLogger) SetCalldepthForDefault() {
	l.calldepth = 3
}

// See Logger.SetLevel
func (l *StdlibLogger) SetLevel(level LogLevel) {
	if level < Everything || level > Off {
		panic("Invalid log level")
	}
	l.level = level
}

// See log.SetOutput
func (l *StdlibLogger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}

// See log.SetFlags
func (l *StdlibLogger) SetFlags(flag int) {
	l.logger.SetFlags(flag)
}

// See log.SetPrefix
func (l *StdlibLogger) SetPrefix(prefix string) {
	l.logger.SetPrefix(prefix)
}

// See Logger.Logs
func (l *StdlibLogger) Logs(level LogLevel) bool {
	return level >= l.level
}

// Traceln logs a line with a TRACE prefix.
func (l *StdlibLogger) Traceln(vals ...interface{}) {
	if !l.Logs(LevelTrace) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(l.calldepth, "TRACE: "+s)
}

// Tracef logs a formatted line with a TRACE prefix.
func (l *StdlibLogger) Tracef(format string, vals ...interface{}) {
	if !l.Logs(LevelTrace) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(l.calldepth, "TRACE: "+s)
}

// Debugln logs a line with a DEBUG prefix.
func (l *StdlibLogger) Debugln(vals ...interface{}) {
	if !l.Logs(LevelDebug) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(l.calldepth, "DEBUG: "+s)
}

// Debugf logs a formatted line with a DEBUG prefix.
func (l *StdlibLogger) Debugf(format string, vals ...interface{}) {
	if !l.Logs(LevelDebug) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(l.calldepth, "DEBUG: "+s)
}

// Infoln logs a line with an INFO prefix.
func (l *StdlibLogger) Infoln(vals ...interface{}) {
	if !l.Logs(LevelInfo) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(l.calldepth, "INFO: "+s)
}

// Infof logs a formatted line with an INFO prefix.
func (l *StdlibLogger) Infof(format string, vals ...interface{}) {
	if !l.Logs(LevelInfo) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(l.calldepth, "INFO: "+s)
}

// Warnln logs a line with a WARNING prefix.
func (l *StdlibLogger) Warnln(vals ...interface{}) {
	if !l.Logs(LevelWarn) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(l.calldepth, "WARNING: "+s)
}

// Warnf logs a formatted line with a WARNING prefix.
func (l *StdlibLogger) Warnf(format string, vals ...interface{}) {
	if !l.Logs(LevelWarn) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(l.calldepth, "WARNING: "+s)
}

// Fatalln logs a line with a FATAL prefix and exits the process with exit code 1.
func (l *StdlibLogger) Fatalln(vals ...interface{}) {
	if !l.Logs(LevelFatal) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(l.calldepth, "FATAL: "+s)
	os.Exit(1)
}

// Fatalf logs a formatted line with a FATAL prefix and exits the process with exit code 1.
func (l *StdlibLogger) Fatalf(format string, vals ...interface{}) {
	if !l.Logs(LevelFatal) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(l.calldepth, "FATAL: "+s)
	os.Exit(1)
}
