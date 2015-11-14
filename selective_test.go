package logger

import "testing"

func TestSelective(t *testing.T) {
	m := newMock()
	l := NewSelective(m)
	s := "hi"

	m.checkLnCount(t, 0, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0, 0)

	l.SetLevel(Off)

	l.Fatalln(s)
	m.checkLnCount(t, 0, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0, 0)

	l.Enable(LevelOK)

	l.Debugln(s)
	l.Verboseln(s)
	l.Infoln(s)
	l.Okln(s)
	l.Warnln(s)
	l.Fatalln(s)

	m.checkLnCount(t, 0, 0, 0, 1, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0, 0)

	l.Disable(LevelOK)
	l.Enable(LevelInfo)

	l.Debugf(s, s)
	l.Verbosef(s, s)
	l.Infof(s, s)
	l.Okf(s, s)
	l.Warnf(s, s)
	l.Fatalf(s, s)

	m.checkLnCount(t, 0, 0, 0, 1, 0, 0)
	m.checkFCount(t, 0, 0, 1, 0, 0, 0)

	l.SetLevel(LevelWarn)
	l.Enable(LevelInfo)

	// Enabled should be: Info, Warn, Fatal

	l.Debugln(s)
	l.Verboseln(s)
	l.Infoln(s)
	l.Okln(s)
	l.Warnln(s)
	l.Fatalln(s)

	m.checkLnCount(t, 0, 0, 1, 1, 1, 1)
	m.checkFCount(t, 0, 0, 1, 0, 0, 0)

}
