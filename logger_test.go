// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import (
	"log"
	"testing"
)

func TestLevels(t *testing.T) {
	l := NewStdlibLogger()
	l.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	l.SetLevel(Everything)
	l.Traceln("hello trace")
	l.Debugln("hello debug")
	l.Infoln("hello info")
	l.Warnln("hello warn")
}

func TestByName(t *testing.T) {
	cases := []struct {
		in   string
		want LogLevel
	}{
		{"e", Everything},
		{"everything", Everything},
		{"t", LevelTrace},
		{"trace", LevelTrace},
		{"LevelTrace", LevelTrace},
		{"d", LevelDebug},
		{"debug", LevelDebug},
		{"LevelDebug", LevelDebug},
		{"i", LevelInfo},
		{"info", LevelInfo},
		{"LevelInfo", LevelInfo},
		{"w", LevelWarn},
		{"warn", LevelWarn},
		{"LevelWarn", LevelWarn},
		{"f", LevelFatal},
		{"fatal", LevelFatal},
		{"LevelFatal", LevelFatal},
		{"o", Off},
		{"off", Off},
	}

	for _, c := range cases {
		got, err := ByName(c.in)
		if err != nil {
			t.Errorf("ByName(%s) returned error %s", c.in, err)
		}
		if got != c.want {
			t.Errorf("ByName(%s) == %d, want %d", c.in, got, c.want)
		}
	}

	got, err := ByName("garbage?")
	if err == nil {
		t.Error("ByName(garbage?) did not return an error")
	}
	if got != Everything {
		t.Errorf("ByName(garbage?) == %d, want %d", got, Everything)
	}

}
