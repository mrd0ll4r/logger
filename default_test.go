// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import "testing"

func TestDefaultFuncs(t *testing.T) {
	m := newMock()
	defaultLogger = m

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

	m.checkLnCount(t, 0, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0)

	Traceln()

	m.checkLnCount(t, 1, 0, 0, 0, 0)
	m.checkFCount(t, 0, 0, 0, 0, 0)

	Tracef("")

	m.checkLnCount(t, 1, 0, 0, 0, 0)
	m.checkFCount(t, 1, 0, 0, 0, 0)

	Debugln()

	m.checkLnCount(t, 1, 1, 0, 0, 0)
	m.checkFCount(t, 1, 0, 0, 0, 0)

	Debugf("")

	m.checkLnCount(t, 1, 1, 0, 0, 0)
	m.checkFCount(t, 1, 1, 0, 0, 0)

	Infoln()

	m.checkLnCount(t, 1, 1, 1, 0, 0)
	m.checkFCount(t, 1, 1, 0, 0, 0)

	Infof("")

	m.checkLnCount(t, 1, 1, 1, 0, 0)
	m.checkFCount(t, 1, 1, 1, 0, 0)

	Warnln()

	m.checkLnCount(t, 1, 1, 1, 1, 0)
	m.checkFCount(t, 1, 1, 1, 0, 0)

	Warnf("")

	m.checkLnCount(t, 1, 1, 1, 1, 0)
	m.checkFCount(t, 1, 1, 1, 1, 0)

	Fatalln()

	m.checkLnCount(t, 1, 1, 1, 1, 1)
	m.checkFCount(t, 1, 1, 1, 1, 0)

	Fatalf("")

	m.checkLnCount(t, 1, 1, 1, 1, 1)
	m.checkFCount(t, 1, 1, 1, 1, 1)
}
