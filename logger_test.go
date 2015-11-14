// Copyright (C) 2014 Jakob Borg. All rights reserved. Use of this source code
// is governed by an MIT-style license that can be found in the LICENSE file.

package logger

import (
	"log"
	"testing"
)

func TestLevels(t *testing.T) {
	l := NewStdlibLogger()
	l.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	l.Debugln("hello debug")
	l.Verboseln("hello verbose")
	l.Infoln("hello info")
	l.Okln("hello OK")
	l.Warnln("hello warn")
}

// All benchmarks should be ran with LOGGER_DISCARD="<non-empty>", to redirect logging to ioutil.Discard

func BenchmarkLshortfileLn(b *testing.B) {
	l := NewStdlibLogger()
	l.SetFlags(log.Lshortfile)
	l.SetLevel(LevelOK)
	s := "hi"
	for i := 0; i < b.N; i++ {
		l.Okln(s)
	}
}

func BenchmarkNoFlagLn(b *testing.B) {
	l := NewStdlibLogger()
	l.SetFlags(0)
	l.SetLevel(LevelOK)
	s := "hi"
	for i := 0; i < b.N; i++ {
		l.Okln(s)
	}
}

func BenchmarkNoFlagF(b *testing.B) {
	l := NewStdlibLogger()
	l.SetFlags(0)
	l.SetLevel(LevelOK)
	s := ""
	s1 := "hi"
	for i := 0; i < b.N; i++ {
		l.Okf(s, s1)
	}
}

func BenchmarkLevelTooLow(b *testing.B) {
	l := NewStdlibLogger()
	l.SetFlags(0)
	l.SetLevel(LevelOK)
	s := "hi"
	for i := 0; i < b.N; i++ {
		l.Debugln(s)
	}
}
