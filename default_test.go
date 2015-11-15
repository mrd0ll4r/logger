// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import "testing"

func TestDefaultFuncs(t *testing.T) {
	m := newMock()
	DefaultLogger = m
	s := "hi"

	SetLevel(Off)
	if m.Level() != Off {
		t.Errorf("Unexpected level: %d, should be %d", m.Level(), Off)
	}

	SetLevel(Everything)
	if m.Level() != Everything {
		t.Errorf("Unexpected level: %d, should be %d", m.Level(), Everything)
	}

	if !Logs(Everything) {
		t.Errorf("DefaultLogger does not log Everything after it was instructed to")
	}

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
