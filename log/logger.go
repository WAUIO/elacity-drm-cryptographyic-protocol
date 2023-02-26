package log

import (
	"log"
	"os"
)

var (
	warnLog  = log.New(os.Stderr, "[WARN] ", log.Ltime|log.Lmicroseconds)
	infoLog  = log.New(os.Stdout, "[INFO] ", log.Ltime|log.Lmicroseconds)
	traceLog = log.New(os.Stdout, "[TRACE] ", log.Ltime|log.Lmicroseconds)
	errLog   = log.New(os.Stderr, "[ERROR] ", log.Ltime|log.Lmicroseconds)
)

func Error(v ...interface{}) {
	errLog.Println(v...)
}

func Errorf(f string, v ...interface{}) {
	errLog.Printf(f, v...)
}

func Warn(v ...interface{}) {
	warnLog.Println(v...)
}

func Warnf(f string, v ...interface{}) {
	warnLog.Printf(f, v...)
}

func Info(v ...interface{}) {
	infoLog.Println(v...)
}

func Infof(f string, v ...interface{}) {
	infoLog.Printf(f, v...)
}

func Trace(v ...interface{}) {
	traceLog.Println(v...)
}

func Tracef(f string, v ...interface{}) {
	traceLog.Printf(f, v...)
}
