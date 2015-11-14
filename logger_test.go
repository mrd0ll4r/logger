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
