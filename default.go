// Copyright (C) 2014 mrd0ll4r. All rights reserved. Use of this source code
// is governed by the MIT license that can be found in the LICENSE file.

package logger

// SetLevel sets the level for the default logger
func SetLevel(level LogLevel) {
	if level < Everything || level > Off {
		panic("Invalid log level")
	}
	defaultLogger.SetLevel(level)
}

// Logs calls Logs on the default logger
func Logs(level LogLevel) bool {
	return defaultLogger.Logs(level)
}

// Traceln calls Traceln on the default logger
func Traceln(val ...interface{}) {
	defaultLogger.Traceln(val...)
}

// Tracef calls Tracef on the default logger
func Tracef(format string, val ...interface{}) {
	defaultLogger.Tracef(format, val...)
}

// Debugln calls Debugln on the default logger
func Debugln(val ...interface{}) {
	defaultLogger.Debugln(val...)
}

// Debugf calls Debugf on the default logger
func Debugf(format string, val ...interface{}) {
	defaultLogger.Debugf(format, val...)
}

// Infoln calls Infoln on the default logger
func Infoln(val ...interface{}) {
	defaultLogger.Infoln(val...)
}

// Infof calls Infof on the default logger
func Infof(format string, val ...interface{}) {
	defaultLogger.Infof(format, val...)
}

// Warnln calls Warnln on the default logger
func Warnln(val ...interface{}) {
	defaultLogger.Warnln(val...)
}

// Warnf calls Warn of the default logger
func Warnf(format string, val ...interface{}) {
	defaultLogger.Warnf(format, val...)
}

// Fatalln calls Fatalln on the default logger
func Fatalln(val ...interface{}) {
	defaultLogger.Fatalln(val...)
}

// Fatalf calls Fatalf on the default logger
func Fatalf(format string, val ...interface{}) {
	defaultLogger.Fatalf(format, val...)
}
