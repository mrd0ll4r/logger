package logger

// SetLevel sets the level for the default logger
func SetLevel(level LogLevel) {
	if level < Off || level > NumLevels {
		panic("Invalid log level")
	}
	DefaultLogger.SetLevel(level)
}

// Debugln calls Debugln on the default logger
func Debugln(val ...interface{}) {
	DefaultLogger.Debugln(val)
}

// Debugf calls Debugf on the default logger
func Debugf(format string, val ...interface{}) {
	DefaultLogger.Debugf(format, val)
}

// Verboseln calls Verboseln on the default logger
func Verboseln(val ...interface{}) {
	DefaultLogger.Verboseln(val)
}

// Verbosef calls Verbosef on the default logger
func Verbosef(format string, val ...interface{}) {
	DefaultLogger.Verbosef(format, val)
}

// Infoln calls Infoln on the default logger
func Infoln(val ...interface{}) {
	DefaultLogger.Infoln(val)
}

// Infof calls Infof on the default logger
func Infof(format string, val ...interface{}) {
	DefaultLogger.Infof(format, val)
}

// Okln calls Okln on the default logger
func Okln(val ...interface{}) {
	DefaultLogger.Okln(val)
}

// Okf calls Okf on the default logger
func Okf(format string, val ...interface{}) {
	DefaultLogger.Okf(format, val)
}

// Warnln calls Warnln on the default logger
func Warnln(val ...interface{}) {
	DefaultLogger.Warnln(val)
}

// Warnf calls Warn of the default logger
func Warnf(format string, val ...interface{}) {
	DefaultLogger.Warnf(format, val)
}

// Fatalln calls Fatalln on the default logger
func Fatalln(val ...interface{}) {
	DefaultLogger.Fatalln(val)
}

// Fatalf calls Fatalf on the default logger
func Fatalf(format string, val ...interface{}) {
	DefaultLogger.Fatalf(format, val)
}
