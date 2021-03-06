// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import "sync"

// Selective is a Logger that can be turned on and off for every level
type Selective struct {
	l      Logger
	levels map[LogLevel]bool
	mu     sync.Mutex
}

// to check if we implement Logger
var _ Logger = NewSelective(defaultLogger)

// NewSelective returns a new Selective, wrapping the provided Logger.
// The Selective returned will be enabled for every level the inner Logger logs
func NewSelective(inner Logger) *Selective {
	toReturn := Selective{
		l:      inner,
		levels: make(map[LogLevel]bool),
	}
	for _, level := range Levels() {
		if inner.Logs(level) {
			toReturn.levels[level] = true
		}
	}
	return &toReturn
}

// SetLevel clears all enables/disables and enables all levels above and including the level provided
func (p *Selective) SetLevel(level LogLevel) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, l := range Levels() {
		if l < level {
			p.levels[l] = false
		} else {
			p.levels[l] = true
		}
	}
}

// See Logger.Logs
func (p *Selective) Logs(level LogLevel) bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.levels[level]
}

// Enable enables a level for logging
func (p *Selective) Enable(level LogLevel) {
	if level < Everything || level > Off {
		panic("Invalid log level")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.levels[level] = true
}

// Disable disables a level for logging
func (p *Selective) Disable(level LogLevel) {
	if level < Everything || level > Off {
		panic("Invalid log level")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.levels[level] = false
}

// Traceln calls Traceln on the inner logger if this Selective is enabled for that level
func (p *Selective) Traceln(val ...interface{}) {
	if !p.Logs(LevelTrace) {
		return
	}
	p.l.Traceln(val)
}

// Tracef calls Tracef on the inner logger if this Selective is enabled for that level
func (p *Selective) Tracef(format string, val ...interface{}) {
	if !p.Logs(LevelTrace) {
		return
	}
	p.l.Tracef(format, val)
}

// Debugln calls Debugln on the inner logger if this Selective is enabled for that level
func (p *Selective) Debugln(val ...interface{}) {
	if !p.Logs(LevelDebug) {
		return
	}
	p.l.Debugln(val)
}

// Debugf calls Debugf on the inner logger if this Selective is enabled for that level
func (p *Selective) Debugf(format string, val ...interface{}) {
	if !p.Logs(LevelDebug) {
		return
	}
	p.l.Debugf(format, val)
}

// Infoln calls Infoln on the inner logger if this Selective is enabled for that level
func (p *Selective) Infoln(val ...interface{}) {
	if !p.Logs(LevelInfo) {
		return
	}
	p.l.Infoln(val)
}

// Infof calls Infof on the inner logger if this Selective is enabled for that level
func (p *Selective) Infof(format string, val ...interface{}) {
	if !p.Logs(LevelInfo) {
		return
	}
	p.l.Infof(format, val)
}

// Warnln calls Warnln on the inner logger if this Selective is enabled for that level
func (p *Selective) Warnln(val ...interface{}) {
	if !p.Logs(LevelWarn) {
		return
	}
	p.l.Warnln(val)
}

// Warnf calls Warnf on the inner logger if this Selective is enabled for that level
func (p *Selective) Warnf(format string, val ...interface{}) {
	if !p.Logs(LevelWarn) {
		return
	}
	p.l.Warnf(format, val)
}

// Fatalln calls Fatalln on the inner logger if this Selective is enabled for that level
func (p *Selective) Fatalln(val ...interface{}) {
	if !p.Logs(LevelFatal) {
		return
	}
	p.l.Fatalln(val)
}

// Fatalf calls Fatalf on the inner logger if this Selective is enabled for that level
func (p *Selective) Fatalf(format string, val ...interface{}) {
	if !p.Logs(LevelFatal) {
		return
	}
	p.l.Fatalf(format, val)
}
