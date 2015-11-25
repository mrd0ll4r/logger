// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

import (
	"log"
	"testing"
)

// All benchmarks should be ran with LOGGER_DISCARD="<non-empty>", to redirect logging to ioutil.Discard

func BenchmarkLshortfileLn(b *testing.B) {
	l := NewStdlibLogger()
	l.SetFlags(log.Lshortfile)
	l.SetLevel(LevelInfo)
	s := "hi"
	for i := 0; i < b.N; i++ {
		l.Infoln(s)
	}
}

func BenchmarkNoFlagLn(b *testing.B) {
	l := NewStdlibLogger()
	l.SetFlags(0)
	l.SetLevel(LevelInfo)
	s := "hi"
	for i := 0; i < b.N; i++ {
		l.Infoln(s)
	}
}

func BenchmarkNoFlagF(b *testing.B) {
	l := NewStdlibLogger()
	l.SetFlags(0)
	l.SetLevel(LevelInfo)
	s := ""
	s1 := "hi"
	for i := 0; i < b.N; i++ {
		l.Infof(s, s1)
	}
}

func BenchmarkLevelTooLow(b *testing.B) {
	l := NewStdlibLogger()
	l.SetFlags(0)
	l.SetLevel(LevelInfo)
	s := "hi"
	for i := 0; i < b.N; i++ {
		l.Debugln(s)
	}
}
