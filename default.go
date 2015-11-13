package logger

func SetLevel(level LogLevel) {
	DefaultLogger.SetLevel(level)
}

func Level() LogLevel {
	return DefaultLogger.Level()
}

func Debugln(val ...interface{}) {
	DefaultLogger.Debugln(val)
}

func Debugf(format string, val ...interface{}) {
	DefaultLogger.Debugf(format, val)
}

func Verboseln(val ...interface{}) {
	DefaultLogger.Verboseln(val)
}

func Verbosef(format string, val ...interface{}) {
	DefaultLogger.Verbosef(val)
}

func Infoln(val ...interface{}) {
	DefaultLogger.Infoln(val)
}

func Infof(format string, val ...interface{}) {
	DefaultLogger.Infof(val)
}

func Okln(val ...interface{}) {
	DefaultLogger.Okln(val)
}

func Okf(format string, val ...interface{}) {
	DefaultLogger.Okf(val)
}

func Warnln(val ...interface{}) {
	DefaultLogger.Warnln(val)
}

func Warnf(format string, val ...interface{}) {
	DefaultLogger.Warnf(val)
}

func Fatalln(val ...interface{}) {
	DefaultLogger.Fatalln(val)
}

func Fatalf(format string, val ...interface{}) {
	DefaultLogger.Fatalf(val)
}
