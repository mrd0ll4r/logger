// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import (
	"sync"
)

// A MultiLogger is a Logger that replicates logging events across multiple Loggers
type MultiLogger struct {
	l     []Logger
	level LogLevel
	mu    sync.Mutex
}

// to see if we implement Logger
var _ Logger = NewMultiLogger(defaultLogger)

// NewMultiLogger wraps the provided loggers into one MultiLogger with LevelInfo logging.
// The inner loggers should not be accessed while the MultiLogger is in use.
// See SetLevel and PropagateLevel for specific level settings
func NewMultiLogger(loggers ...Logger) *MultiLogger {
	if len(loggers) < 1 {
		panic("MultiLogger: No loggers provided")
	}
	return &MultiLogger{
		l:     loggers,
		level: LevelInfo,
	}
}

// PropagateLevel propagates this loggers level to all underlying loggers
func (m *MultiLogger) PropagateLevel() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, v := range m.l {
		v.SetLevel(m.level)
	}
}

// SetLevel sets the level for this logger only, it does not change the underlying loggers.
func (m *MultiLogger) SetLevel(level LogLevel) {
	if level < Everything || level > Off {
		panic("Invalid log level")
	}
	m.level = level
}

// See Logger.Logs
func (m *MultiLogger) Logs(level LogLevel) bool {
	return level >= m.level
}

// Traceln calls Traceln on all the inner loggers
func (m *MultiLogger) Traceln(val ...interface{}) {
	if !m.Logs(LevelTrace) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Traceln(val)
	}
}

// Tracef calls Tracef on all the inner loggers
func (m *MultiLogger) Tracef(format string, val ...interface{}) {
	if !m.Logs(LevelTrace) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.Tracef(format, val)
	}
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
	first := m.l[0]
	for _, l := range m.l[1:] {
		l.Warnln(append([]interface{}{"[MultiLogger-FATAL]:"}, val...))
	}
	first.Fatalln(val)

}

// Fatalf calls Warnf on all inner loggers except for the first, including a [MultiLogger-FATAL] tag.
// It then calls Fatalf on the first inner logger, which should exit the program with code 1.
func (m *MultiLogger) Fatalf(format string, val ...interface{}) {
	if !m.Logs(LevelFatal) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	first := m.l[0]
	for _, l := range m.l[1:] {
		l.Warnf("[MultiLogger-FATAL]: "+format, val...)
	}
	first.Fatalf(format, val)

}
