package logger

import "testing"

func TestMock(t *testing.T) {
	m := newMock()
	s := "hi"

	m.checkLnCount(t, 0, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0, 0)

	m.Debugln(s)

	m.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0, 0)

	m.Debugf(s, s)

	m.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m.checkFCount(t, 1, 0, 0, 0, 0, 0)

	m.Verboseln(s)

	m.checkLnCount(t, 1, 1, 0, 0, 0, 0)
	m.checkFCount(t, 1, 0, 0, 0, 0, 0)

	m.Verbosef(s, s)

	m.checkLnCount(t, 1, 1, 0, 0, 0, 0)
	m.checkFCount(t, 1, 1, 0, 0, 0, 0)

	m.Infoln(s)

	m.checkLnCount(t, 1, 1, 1, 0, 0, 0)
	m.checkFCount(t, 1, 1, 0, 0, 0, 0)

	m.Infof(s, s)

	m.checkLnCount(t, 1, 1, 1, 0, 0, 0)
	m.checkFCount(t, 1, 1, 1, 0, 0, 0)

	m.Okln(s)

	m.checkLnCount(t, 1, 1, 1, 1, 0, 0)
	m.checkFCount(t, 1, 1, 1, 0, 0, 0)

	m.Okf(s, s)

	m.checkLnCount(t, 1, 1, 1, 1, 0, 0)
	m.checkFCount(t, 1, 1, 1, 1, 0, 0)

	m.Warnln(s)

	m.checkLnCount(t, 1, 1, 1, 1, 1, 0)
	m.checkFCount(t, 1, 1, 1, 1, 0, 0)

	m.Warnf(s, s)

	m.checkLnCount(t, 1, 1, 1, 1, 1, 0)
	m.checkFCount(t, 1, 1, 1, 1, 1, 0)

	m.Fatalln(s)

	m.checkLnCount(t, 1, 1, 1, 1, 1, 1)
	m.checkFCount(t, 1, 1, 1, 1, 1, 0)

	m.Fatalf(s, s)

	m.checkLnCount(t, 1, 1, 1, 1, 1, 1)
	m.checkFCount(t, 1, 1, 1, 1, 1, 1)

}
