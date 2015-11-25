// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import "testing"

func TestSelective(t *testing.T) {
	m := newMock()
	l := NewSelective(m)
	s := "hi"

	// check base state
	m.checkLnCount(t, 0, 0, 0,  0, 0)
	m.checkFCount(t, 0, 0, 0,  0, 0)

	l.SetLevel(Off)
	// check disable everything
	l.Traceln(s)
	l.Tracef(s, s)
	l.Debugln(s)
	l.Debugf(s, s)
	l.Infoln(s)
	l.Infof(s, s)
	l.Warnln(s)
	l.Warnf(s, s)
	l.Fatalln(s)
	l.Fatalf(s)
	m.checkLnCount(t, 0, 0, 0,  0, 0)
	m.checkFCount(t, 0, 0, 0,  0, 0)

	l.Enable(LevelInfo)

	// enabled should be: Info

	l.Traceln(s)
	l.Debugln(s)
	l.Infoln(s)
	l.Warnln(s)
	l.Fatalln(s)

	m.checkLnCount(t, 0, 0, 1, 0, 0)
	m.checkFCount(t, 0, 0, 0,  0, 0)

	l.Disable(LevelInfo)
	l.Enable(LevelDebug)

	// enabled should be: Debug

	l.Tracef(s,s)
	l.Debugf(s, s)
	l.Infof(s, s)
	l.Warnf(s, s)
	l.Fatalf(s, s)

	m.checkLnCount(t, 0, 0, 1,  0, 0)
	m.checkFCount(t, 0, 1, 0,  0, 0)

	l.SetLevel(LevelWarn)
	l.Enable(LevelInfo)

	// Enabled should be: Info, Warn, Fatal

	l.Traceln(s)
	l.Debugln(s)
	l.Infoln(s)
	l.Warnln(s)
	l.Fatalln(s)

	m.checkLnCount(t, 0, 0, 2,  1, 1)
	m.checkFCount(t, 0, 1, 0,  0, 0)
}
