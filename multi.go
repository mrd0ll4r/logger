package logger

import (
	"os"
	"sync"
)

// A MultiLogger is a Logger that replicates logging events across multiple Loggers
type MultiLogger struct {
	l     []Logger
	level LogLevel
	mu    sync.Mutex
}

// to see if we implement Logger
var _ Logger = NewMultiLogger(DefaultLogger)

// NewMultiLogger wraps the provided loggers into one MultiLogger.
// The inner loggers should not be accessed while the MultiLogger is in use.
func NewMultiLogger(loggers ...Logger) *MultiLogger {
	return &MultiLogger{
		l:     loggers,
		level: LevelOK,
	}
}

func (m *MultiLogger) SetLevel(level LogLevel) {
	if level < Off || level > NumLevels {
		panic("Invalid log level")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.SetLevel(level)
	}
	m.level = level
}

func (m *MultiLogger) Logs(level LogLevel) bool {
	return level <= m.level
}

// Debugln calls Debugln on all the inner loggers
func (m *MultiLogger) Debugln(val ...interface{}) {
	if !m.Logs(LevelDebug) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Debugln(val)
	}
}

// Debugf calls Debugf on all the inner loggers
func (m *MultiLogger) Debugf(format string, val ...interface{}) {
	if !m.Logs(LevelDebug) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Debugf(format, val)
	}
}

// Verboseln calls Verboseln on all the inner loggers
func (m *MultiLogger) Verboseln(val ...interface{}) {
	if !m.Logs(LevelVerbose) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Verboseln(val)
	}
}

// Verbosef calls Verbosef on all the inner loggers
func (m *MultiLogger) Verbosef(format string, val ...interface{}) {
	if !m.Logs(LevelVerbose) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Verbosef(format, val)
	}
}

// Infoln calls Infoln on all the inner loggers
func (m *MultiLogger) Infoln(val ...interface{}) {
	if !m.Logs(LevelInfo) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Infoln(val)
	}
}

// Infof calls Infof on all the inner loggers
func (m *MultiLogger) Infof(format string, val ...interface{}) {
	if !m.Logs(LevelInfo) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Infof(format, val)
	}
}

// Okln calls Okln on all the inner loggers
func (m *MultiLogger) Okln(val ...interface{}) {
	if !m.Logs(LevelOK) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Okln(val)
	}
}

// Okf calls Okf on all the inner loggers
func (m *MultiLogger) Okf(format string, val ...interface{}) {
	if !m.Logs(LevelOK) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Okf(format, val)
	}
}

// Warnln calls Warnln on all the inner loggers
func (m *MultiLogger) Warnln(val ...interface{}) {
	if !m.Logs(LevelWarn) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Warnln(val)
	}
}

// Warnf calls Warnf on all the inner loggers
func (m *MultiLogger) Warnf(format string, val ...interface{}) {
	if !m.Logs(LevelWarn) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Warnf(format, val)
	}
}

// Fatalln calls Warnln on all inner loggers except for the first, including a [MultiLogger-FATAL] tag.
// It then calls Fatalln on the first inner logger, which should exit the program with code 1.
func (m *MultiLogger) Fatalln(val ...interface{}) {
	if !m.Logs(LevelFatal) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.l) < 1 {
		os.Exit(1)
		return
	} else {
		first := m.l[0]
		for _, l := range m.l[1:] {
			l.Warnln(append([]interface{}{"[MultiLogger-FATAL]:"}, val...))
		}
		first.Fatalln(val)
	}
}

// Fatalf calls Warnf on all inner loggers except for the first, including a [MultiLogger-FATAL] tag.
// It then calls Fatalf on the first inner logger, which should exit the program with code 1.
func (m *MultiLogger) Fatalf(format string, val ...interface{}) {
	if !m.Logs(LevelFatal) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.l) < 1 {
		os.Exit(1)
		return
	} else {
		first := m.l[0]
		for _, l := range m.l[1:] {
			l.Warnf("[MultiLogger-FATAL]: "+format, val...)
		}
		first.Fatalf(format, val)
	}
}
