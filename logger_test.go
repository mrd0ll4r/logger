// Copyright (C) 2014 Jakob Borg. All rights reserved. Use of this source code
// is governed by an MIT-style license that can be found in the LICENSE file.

package logger

import (
	"log"
	"testing"
)

func TestLevels(t *testing.T) {
	l := New().(*StdlibLogger)
	l.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	l.Debugln("hello debug")
	l.Verboseln("hello verbose")
	l.Infoln("hello info")
	l.Okln("hello OK")
	l.Warnln("hello warn")
}

// All benchmarks should be ran with LOGGER_DISCARD="<non-empty>", to redirect logging to ioutil.Discard

func BenchmarkLshortfile(b *testing.B) {
	l := New().(*StdlibLogger)
	l.SetFlags(log.Lshortfile)
	l.SetLevel(LevelOK)
	for i := 0; i < b.N; i++ {
		l.Okln("hi")
	}
}

func BenchmarkBare(b *testing.B) {
	l := New().(*StdlibLogger)
	l.SetFlags(0)
	l.SetLevel(LevelOK)
	for i := 0; i < b.N; i++ {
		l.Okln("hi")
	}
}

func BenchmarkLevelTooLowBare(b *testing.B) {
	l := New().(*StdlibLogger)
	l.SetFlags(0)
	l.SetLevel(LevelOK)
	for i := 0; i < b.N; i++ {
		l.Debugln("hi")
	}
}
