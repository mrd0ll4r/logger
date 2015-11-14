package logger

import "testing"

func TestMulti(t *testing.T) {
	m1 := newMock()
	m2 := newMock()
	l := NewMultiLogger(m1, m2)

	l.SetLevel(Off)
	if l.Logs(LevelDebug) {
		t.Error("MultiLogger still logs after turing it Off")
	}

	l.PropagateLevel()
	if m1.Level() != Off || m2.Level() != Off {
		t.Errorf("PropagateLevel did not propagate properly, expected: %d,%d - got %d,%d", Off, Off, m1.Level(), m2.Level())
	}

	l.SetLevel(Everything)
	if !l.Logs(Everything) {
		t.Error("MultiLogger does not log Everything after it was instructed to")
	}

	l.PropagateLevel()
	if m1.Level() != Everything || m2.Level() != Everything {
		t.Errorf("PropagateLevel did not propagate properly, expected: %d,%d - got %d,%d", Everything, Everything, m1.Level(), m2.Level())
	}

	m1.checkLnCount(t, 0, 0, 0, 0, 0, 0)
	m1.checkFCount(t, 0, 0, 0, 0, 0, 0)
	m2.checkLnCount(t, 0, 0, 0, 0, 0, 0)
	m2.checkFCount(t, 0, 0, 0, 0, 0, 0)

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
	l.Debugln()
	l.Verboseln()
	l.Infoln()
	l.Okln()

	l.Debugf("")
	l.Verbosef("")
	l.Infof("")
	l.Okf("")

	m1.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m1.checkFCount(t, 0, 1, 0, 0, 0, 0)
	m2.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m2.checkFCount(t, 0, 1, 0, 0, 0, 0)

	l.SetLevel(Off)

	l.Debugln()
	l.Debugf("")
	l.Verboseln()
	l.Verbosef("")
	l.Infoln()
	l.Infof("")
	l.Okln()
	l.Okf("")
	l.Warnln()
	l.Warnf("")
	l.Fatalln()

	m1.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m1.checkFCount(t, 0, 1, 0, 0, 0, 0)
	m2.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m2.checkFCount(t, 0, 1, 0, 0, 0, 0)
}
