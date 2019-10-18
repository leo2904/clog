package clog

import (
	"io"
	"io/ioutil"
	"os"
)


type Logger struct {
	stdOutput io.Writer
}

func New() *Logger {
	logger := &Logger{
		stdOutput: os.Stdout,
	}
	return logger
}

func NewNop() *Logger {
	logger := &Logger{
		stdOutput: ioutil.Discard,
	}
	return logger
}

func (l *Logger) Print(fields ...Field) {
	printCanonicalLine(l.stdOutput, fields...)
}