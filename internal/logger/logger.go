package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func NewLogger(prefix string) *Logger {
	writter := io.Writer(os.Stdout)
	l := log.New(writter, prefix, log.LstdFlags)

	return &Logger{
		debug:  log.New(writter, "DEBUG: ", l.Flags()),
		info:   log.New(writter, "INFO: ", l.Flags()),
		warn:   log.New(writter, "WARN: ", l.Flags()),
		err:    log.New(writter, "ERROR: ", l.Flags()),
		writer: writter,
	}
}

func (l *Logger) Debug(s string) {
	l.debug.Println(s)
}

func (l *Logger) Info(s string) {
	l.info.Println(s)
}

func (l *Logger) Warn(s string) {
	l.warn.Println(s)
}

func (l *Logger) Error(s string) {
	l.err.Println(s)
}

func (l *Logger) Debugf(s string, v ...interface{}) {
	l.debug.Printf(s, v...)
}

func (l *Logger) Infof(s string, v ...interface{}) {
	l.info.Printf(s, v...)
}

func (l *Logger) Warnf(s string, v ...interface{}) {
	l.warn.Printf(s, v...)
}

func (l *Logger) Errorf(s string, v ...interface{}) {
	l.err.Printf(s, v...)
}

func GetLogger(p string) *Logger {
	return NewLogger(p)
}
