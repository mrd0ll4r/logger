// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import (
	"sync"
	"testing"
)

// mock is a mock logger that just counts how often <Level>(ln|f) was called
type mock struct {
	lnCount map[LogLevel]int
	fCount  map[LogLevel]int
	level   LogLevel
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

func (m *mock) checkLnCount(t *testing.T, trace, debug, info, warn, fatal int) {
	if !(m.LnCount(LevelTrace) == trace &&
		m.LnCount(LevelDebug) == debug &&
		m.LnCount(LevelInfo) == info &&
		m.LnCount(LevelWarn) == warn &&
		m.LnCount(LevelFatal) == fatal) {
		t.Errorf("LnCount is wrong, should be %d,%d,%d,%d,%d - is %d,%d,%d,%d,%d", trace, debug, info, warn, fatal,
			m.LnCount(LevelTrace),
			m.LnCount(LevelDebug),
			m.LnCount(LevelInfo),
			m.LnCount(LevelWarn),
			m.LnCount(LevelFatal))
	}
}

func (m *mock) checkFCount(t *testing.T, trace, debug, info, warn, fatal int) {
	if !(m.FCount(LevelTrace) == trace &&
		m.FCount(LevelDebug) == debug &&
		m.FCount(LevelInfo) == info &&
		m.FCount(LevelWarn) == warn &&
		m.FCount(LevelFatal) == fatal) {
		t.Errorf("FCount is wrong, should be %d,%d,%d,%d,%d - is %d,%d,%d,%d,%d", trace, debug, info, warn, fatal,
			m.FCount(LevelTrace),
			m.FCount(LevelDebug),
			m.FCount(LevelInfo),
			m.FCount(LevelWarn),
			m.FCount(LevelFatal))
	}
}

func (m *mock) Logs(level LogLevel) bool {
	return level >= m.level
}

func (m *mock) SetLevel(level LogLevel) {
	m.level = level
}

func (m *mock) Level() LogLevel { return m.level }

func (m *mock) Traceln(val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lnCount[LevelTrace]++
}

func (m *mock) Tracef(format string, val ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.fCount[LevelTrace]++
}

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
