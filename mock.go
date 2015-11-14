package logger

import "sync"

type mock struct {
	lnCount map[LogLevel]int
	fCount  map[LogLevel]int
	mu      sync.Mutex
}

var _ Logger = newMock()

func newMock() *mock {
	return &mock{
		lnCount: make(map[LogLevel]int),
		fCount:  make(map[LogLevel]int),
	}
}

func (m *mock) LnCount(level LogLevel) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.lnCount[level]
}

func (m *mock) FCount(level LogLevel) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.fCount[level]
}

func (m *mock) Logs(level LogLevel) bool {
	return level >= LevelOK
}

func (m *mock) SetLevel(LogLevel) {}

func (m *mock) Level() LogLevel { return LevelOK }

func (m *mock) Debugln(val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lnCount[LevelDebug]++
}

func (m *mock) Debugf(format string, val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.fCount[LevelDebug]++
}

func (m *mock) Verboseln(val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lnCount[LevelVerbose]++
}

func (m *mock) Verbosef(format string, val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.fCount[LevelVerbose]++
}

func (m *mock) Infoln(val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lnCount[LevelInfo]++
}

func (m *mock) Infof(format string, val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.fCount[LevelInfo]++
}

func (m *mock) Okln(val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lnCount[LevelOK]++
}

func (m *mock) Okf(format string, val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.fCount[LevelOK]++
}

func (m *mock) Warnln(val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lnCount[LevelWarn]++
}

func (m *mock) Warnf(format string, val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.fCount[LevelWarn]++
}

func (m *mock) Fatalln(val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lnCount[LevelFatal]++
}

func (m *mock) Fatalf(format string, val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.fCount[LevelFatal]++
}
