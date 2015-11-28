// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import "testing"

func TestSelective(t *testing.T) {
	m := newMock()
	l := NewSelective(m)

	// check base state
	m.checkLnCount(t, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0)

	l.SetLevel(Off)
	// check disable everything
	l.Traceln()
	l.Tracef("")
	l.Debugln()
	l.Debugf("")
	l.Infoln()
	l.Infof("")
	l.Warnln()
	l.Warnf("")
	l.Fatalln()
	l.Fatalf("")
	m.checkLnCount(t, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0)

	l.Enable(LevelInfo)

	// enabled should be: Info

	l.Traceln()
	l.Tracef("")
	l.Debugln()
	l.Debugf("")
	l.Infoln()
	l.Infof("")
	l.Warnln()
	l.Warnf("")
	l.Fatalln()
	l.Fatalf("")

	m.checkLnCount(t, 0, 0, 1, 0, 0)
	m.checkFCount(t, 0, 0, 1, 0, 0)

	l.Disable(LevelInfo)
	l.Enable(LevelDebug)

	// enabled should be: Debug

	l.Traceln()
	l.Tracef("")
	l.Debugln()
	l.Debugf("")
	l.Infoln()
	l.Infof("")
	l.Warnln()
	l.Warnf("")
	l.Fatalln()
	l.Fatalf("")

	m.checkLnCount(t, 0, 1, 1, 0, 0)
	m.checkFCount(t, 0, 1, 1, 0, 0)

	l.SetLevel(LevelWarn)
	l.Enable(LevelInfo)

	// Enabled should be: Info, Warn, Fatal

	l.Traceln()
	l.Tracef("")
	l.Debugln()
	l.Debugf("")
	l.Infoln()
	l.Infof("")
	l.Warnln()
	l.Warnf("")
	l.Fatalln()
	l.Fatalf("")

	m.checkLnCount(t, 0, 1, 2, 1, 1)
	m.checkFCount(t, 0, 1, 2, 1, 1)

	l.Enable(LevelTrace)

	l.Traceln()
	l.Tracef("")

	m.checkLnCount(t, 1, 1, 2, 1, 1)
	m.checkFCount(t, 1, 1, 2, 1, 1)
}
