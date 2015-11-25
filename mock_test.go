// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import "testing"

func TestMock(t *testing.T) {
	m := newMock()
	s := "hi"

	m.checkLnCount(t, 0, 0, 0,  0, 0)
	m.checkFCount(t, 0, 0, 0,  0, 0)

	m.Traceln(s)

	m.checkLnCount(t, 1, 0, 0,  0, 0)
	m.checkFCount(t, 0, 0, 0,  0, 0)

	m.Tracef(s, s)

	m.checkLnCount(t, 1, 0, 0,  0, 0)
	m.checkFCount(t, 1, 0, 0,  0, 0)

	m.Debugln(s)

	m.checkLnCount(t, 1, 1, 0,  0, 0)
	m.checkFCount(t, 1, 0, 0,  0, 0)

	m.Debugf(s, s)

	m.checkLnCount(t, 1, 1, 0,  0, 0)
	m.checkFCount(t, 1, 1, 0,  0, 0)

	m.Infoln(s)

	m.checkLnCount(t, 1, 1, 1,  0, 0)
	m.checkFCount(t, 1, 1, 0,  0, 0)

	m.Infof(s, s)

	m.checkLnCount(t, 1, 1, 1,  0, 0)
	m.checkFCount(t, 1, 1, 1,  0, 0)

	m.Warnln(s)

	m.checkLnCount(t, 1, 1, 1,  1, 0)
	m.checkFCount(t, 1, 1, 1,  0, 0)

	m.Warnf(s, s)

	m.checkLnCount(t, 1, 1, 1,  1, 0)
	m.checkFCount(t, 1, 1, 1,  1, 0)

	m.Fatalln(s)

	m.checkLnCount(t, 1, 1, 1,  1, 1)
	m.checkFCount(t, 1, 1, 1,  1, 0)

	m.Fatalf(s, s)

	m.checkLnCount(t, 1, 1, 1,  1, 1)
	m.checkFCount(t, 1, 1, 1,  1, 1)

	m.SetLevel(Off)
	if m.Level() != Off {
		t.Errorf("Unexpected level: %d, should be %d", m.Level(), Off)
	}

	m.SetLevel(Everything)
	if m.Level() != Everything {
		t.Errorf("Unexpected level: %d, should be %d", m.Level(), Everything)
	}
}
