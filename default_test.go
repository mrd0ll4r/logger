package logger

import "testing"

func TestDefaultFuncs(t *testing.T) {
	m := newMock()
	DefaultLogger = m

	s := "hi"

	m.checkLnCount(t, 0, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0, 0)

	Debugln(s)

	m.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0, 0)

	Debugf(s, s)

	m.checkLnCount(t, 1, 0, 0, 0, 0, 0)
	m.checkFCount(t, 1, 0, 0, 0, 0, 0)

	Verboseln(s)

	m.checkLnCount(t, 1, 1, 0, 0, 0, 0)
	m.checkFCount(t, 1, 0, 0, 0, 0, 0)

	Verbosef(s, s)

	m.checkLnCount(t, 1, 1, 0, 0, 0, 0)
	m.checkFCount(t, 1, 1, 0, 0, 0, 0)

	Infoln(s)

	m.checkLnCount(t, 1, 1, 1, 0, 0, 0)
	m.checkFCount(t, 1, 1, 0, 0, 0, 0)

	Infof(s, s)

	m.checkLnCount(t, 1, 1, 1, 0, 0, 0)
	m.checkFCount(t, 1, 1, 1, 0, 0, 0)

	Okln(s)

	m.checkLnCount(t, 1, 1, 1, 1, 0, 0)
	m.checkFCount(t, 1, 1, 1, 0, 0, 0)

	Okf(s, s)

	m.checkLnCount(t, 1, 1, 1, 1, 0, 0)
	m.checkFCount(t, 1, 1, 1, 1, 0, 0)

	Warnln(s)

	m.checkLnCount(t, 1, 1, 1, 1, 1, 0)
	m.checkFCount(t, 1, 1, 1, 1, 0, 0)

	Warnf(s, s)

	m.checkLnCount(t, 1, 1, 1, 1, 1, 0)
	m.checkFCount(t, 1, 1, 1, 1, 1, 0)

	Fatalln(s)

	m.checkLnCount(t, 1, 1, 1, 1, 1, 1)
	m.checkFCount(t, 1, 1, 1, 1, 1, 0)

	Fatalf(s, s)

	m.checkLnCount(t, 1, 1, 1, 1, 1, 1)
	m.checkFCount(t, 1, 1, 1, 1, 1, 1)

}
