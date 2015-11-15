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
	l.Debugln("hello debug")
	l.Verboseln("hello verbose")
	l.Infoln("hello info")
	l.Okln("hello OK")
	l.Warnln("hello warn")
}
