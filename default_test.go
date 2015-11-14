package logger

import "testing"

func TestDefaultFuncs(t *testing.T) {
	m := newMock()
	DefaultLogger = m

	s := "hi"

	checkLnCount(t, m, 0, 0, 0, 0, 0, 0)
	checkFCount(t, m, 0, 0, 0, 0, 0, 0)

	Debugln(s)

	checkLnCount(t, m, 1, 0, 0, 0, 0, 0)
	checkFCount(t, m, 0, 0, 0, 0, 0, 0)

	Debugf(s, s)

	checkLnCount(t, m, 1, 0, 0, 0, 0, 0)
	checkFCount(t, m, 1, 0, 0, 0, 0, 0)

	Verboseln(s)

	checkLnCount(t, m, 1, 1, 0, 0, 0, 0)
	checkFCount(t, m, 1, 0, 0, 0, 0, 0)

	Verbosef(s, s)

	checkLnCount(t, m, 1, 1, 0, 0, 0, 0)
	checkFCount(t, m, 1, 1, 0, 0, 0, 0)

	Infoln(s)

	checkLnCount(t, m, 1, 1, 1, 0, 0, 0)
	checkFCount(t, m, 1, 1, 0, 0, 0, 0)

	Infof(s, s)

	checkLnCount(t, m, 1, 1, 1, 0, 0, 0)
	checkFCount(t, m, 1, 1, 1, 0, 0, 0)

	Okln(s)

	checkLnCount(t, m, 1, 1, 1, 1, 0, 0)
	checkFCount(t, m, 1, 1, 1, 0, 0, 0)

	Okf(s, s)

	checkLnCount(t, m, 1, 1, 1, 1, 0, 0)
	checkFCount(t, m, 1, 1, 1, 1, 0, 0)

	Warnln(s)

	checkLnCount(t, m, 1, 1, 1, 1, 1, 0)
	checkFCount(t, m, 1, 1, 1, 1, 0, 0)

	Warnf(s, s)

	checkLnCount(t, m, 1, 1, 1, 1, 1, 0)
	checkFCount(t, m, 1, 1, 1, 1, 1, 0)

	Fatalln(s)

	checkLnCount(t, m, 1, 1, 1, 1, 1, 1)
	checkFCount(t, m, 1, 1, 1, 1, 1, 0)

	Fatalf(s, s)

	checkLnCount(t, m, 1, 1, 1, 1, 1, 1)
	checkFCount(t, m, 1, 1, 1, 1, 1, 1)

}
