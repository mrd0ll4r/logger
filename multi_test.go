package logger

import "testing"

func TestMulti(t *testing.T) {
	m1 := newMock()
	m2 := newMock()
	l := NewMultiLogger(m1, m2)

	m1.checkLnCount(t, 0, 0, 0, 0, 0, 0)
	m1.checkFCount(t, 0, 0, 0, 0, 0, 0)
	m2.checkLnCount(t, 0, 0, 0, 0, 0, 0)
	m2.checkFCount(t, 0, 0, 0, 0, 0, 0)

	l.SetLevel(Everything)
	l.Debugln()

	m1.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m1.checkFCount(t, 0, 0, 0, 0, 0, 0)
	m2.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m2.checkFCount(t, 0, 0, 0, 0, 0, 0)

	l.Verbosef("")

	m1.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m1.checkFCount(t, 0, 1, 0, 0, 0, 0)
	m2.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m2.checkFCount(t, 0, 1, 0, 0, 0, 0)

	l.SetLevel(LevelWarn)

	// should not log
	l.Infoln()

	m1.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m1.checkFCount(t, 0, 1, 0, 0, 0, 0)
	m2.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m2.checkFCount(t, 0, 1, 0, 0, 0, 0)

}
