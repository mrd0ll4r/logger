package logger

import "sync"

type MultiLogger struct {
	l     []Logger
	level LogLevel
	mu    sync.Mutex
}

// to see if we implement Logger
var ml Logger = NewMultiLogger(DefaultLogger)

func NewMultiLogger(loggers ...Logger) *MultiLogger {
	return &MultiLogger{
		l:     loggers,
		level: LevelOK,
	}
}

func (m *MultiLogger) SetLevel(level LogLevel) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, l := range m.l {
		l.SetLevel(level)
	}
	m.level = level
}

func (m *MultiLogger) Level() LogLevel {
	return m.level
}

func (m *MultiLogger) Logs(level LogLevel) bool {
	return level <= m.level
}

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

func (m *MultiLogger) Fatalln(val ...interface{}) {
	// Fatal logging can not work with multi loggers
}

func (m *MultiLogger) Fatalf(format string, val ...interface{}) {
	// Fatal logging can not work with multi loggers
}
