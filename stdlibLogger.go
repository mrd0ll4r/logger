package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// StdlibLogger is an implementation of the Logger interface that uses the stdlibs loggger for logging.
type StdlibLogger struct {
	logger *log.Logger
	level  LogLevel
}

// to see if we implement Logger
var _ Logger = NewStdlibLogger()

// NewStdlibLogger returns a new StdlibLogger.
// If the environment variable LOGGER_DISCARD is set, the logger will discard everything (useful for benchmarks)
func NewStdlibLogger() *StdlibLogger {
	if os.Getenv("LOGGER_DISCARD") != "" {
		// Hack to completely disable logging, for example when running benchmarks.
		return &StdlibLogger{
			logger: log.New(ioutil.Discard, "", 0),
			level:  Off,
		}
	}

	return &StdlibLogger{
		logger: log.New(os.Stdout, "", log.Ltime),
		level:  LevelOK,
	}
}

func (l *StdlibLogger) SetLevel(level LogLevel) {
	if level < 0 || level > NumLevels {
		panic("Invalid log level")
	}
	l.level = level
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
	return l.check(level)
}

func (l *StdlibLogger) check(level LogLevel) bool {
	return level >= l.level
}

// Debugln logs a line with a DEBUG prefix.
func (l *StdlibLogger) Debugln(vals ...interface{}) {
	if !l.check(LevelDebug) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(2, "DEBUG: "+s)
}

// Debugf logs a formatted line with a DEBUG prefix.
func (l *StdlibLogger) Debugf(format string, vals ...interface{}) {
	if !l.check(LevelDebug) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(2, "DEBUG: "+s)
}

// Infoln logs a line with a VERBOSE prefix.
func (l *StdlibLogger) Verboseln(vals ...interface{}) {
	if !l.check(LevelVerbose) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(2, "VERBOSE: "+s)
}

// Infof logs a formatted line with a VERBOSE prefix.
func (l *StdlibLogger) Verbosef(format string, vals ...interface{}) {
	if !l.check(LevelVerbose) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(2, "VERBOSE: "+s)
}

// Infoln logs a line with an INFO prefix.
func (l *StdlibLogger) Infoln(vals ...interface{}) {
	if !l.check(LevelInfo) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(2, "INFO: "+s)
}

// Infof logs a formatted line with an INFO prefix.
func (l *StdlibLogger) Infof(format string, vals ...interface{}) {
	if !l.check(LevelInfo) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(2, "INFO: "+s)
}

// Okln logs a line with an OK prefix.
func (l *StdlibLogger) Okln(vals ...interface{}) {
	if !l.check(LevelOK) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(2, "OK: "+s)
}

// Okf logs a formatted line with an OK prefix.
func (l *StdlibLogger) Okf(format string, vals ...interface{}) {
	if !l.check(LevelOK) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(2, "OK: "+s)
}

// Warnln logs a line with a WARNING prefix.
func (l *StdlibLogger) Warnln(vals ...interface{}) {
	if !l.check(LevelWarn) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(2, "WARNING: "+s)
}

// Warnf logs a formatted line with a WARNING prefix.
func (l *StdlibLogger) Warnf(format string, vals ...interface{}) {
	if !l.check(LevelWarn) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(2, "WARNING: "+s)
}

// Fatalln logs a line with a FATAL prefix and exits the process with exit
// code 1.
func (l *StdlibLogger) Fatalln(vals ...interface{}) {
	if !l.check(LevelFatal) {
		return
	}
	s := fmt.Sprintln(vals...)
	l.logger.Output(2, "FATAL: "+s)
	os.Exit(1)
}

// Fatalf logs a formatted line with a FATAL prefix and exits the process with
// exit code 1.
func (l *StdlibLogger) Fatalf(format string, vals ...interface{}) {
	if !l.check(LevelFatal) {
		return
	}
	s := fmt.Sprintf(format, vals...)
	l.logger.Output(2, "FATAL: "+s)
	os.Exit(1)
}
