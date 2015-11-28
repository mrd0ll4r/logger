// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import "testing"

func TestMock(t *testing.T) {
	m := newMock()

	m.checkLnCount(t, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0)

	m.Traceln()

	m.checkLnCount(t, 1, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0)

	m.Tracef("")

	m.checkLnCount(t, 1, 0, 0, 0, 0)
	m.checkFCount(t, 1, 0, 0, 0, 0)

	m.Debugln()

	m.checkLnCount(t, 1, 1, 0, 0, 0)
	m.checkFCount(t, 1, 0, 0, 0, 0)

	m.Debugf("")

	m.checkLnCount(t, 1, 1, 0, 0, 0)
	m.checkFCount(t, 1, 1, 0, 0, 0)

	m.Infoln()

	m.checkLnCount(t, 1, 1, 1, 0, 0)
	m.checkFCount(t, 1, 1, 0, 0, 0)

	m.Infof("")

	m.checkLnCount(t, 1, 1, 1, 0, 0)
	m.checkFCount(t, 1, 1, 1, 0, 0)

	m.Warnln()

	m.checkLnCount(t, 1, 1, 1, 1, 0)
	m.checkFCount(t, 1, 1, 1, 0, 0)

	m.Warnf("")

	m.checkLnCount(t, 1, 1, 1, 1, 0)
	m.checkFCount(t, 1, 1, 1, 1, 0)

	m.Fatalln()

	m.checkLnCount(t, 1, 1, 1, 1, 1)
	m.checkFCount(t, 1, 1, 1, 1, 0)

	m.Fatalf("")

	m.checkLnCount(t, 1, 1, 1, 1, 1)
	m.checkFCount(t, 1, 1, 1, 1, 1)

	m.SetLevel(Off)
	if m.Level() != Off {
		t.Errorf("Unexpected level: %d, should be %d", m.Level(), Off)
	}

	m.SetLevel(Everything)
	if m.Level() != Everything {
		t.Errorf("Unexpected level: %d, should be %d", m.Level(), Everything)
	}
}
