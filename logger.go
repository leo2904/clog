package clog

import (
	"github.com/friendsofgo/clog/reporter"
	"github.com/friendsofgo/clog/reporter/log"
)

// Logger is our Logger for Canonical Log Lines implementation. You should initialize it using
// the New method.
type Logger struct {
	reporters []reporter.Reporter
}

// LoggerOption allow to adjust behaviour of the Logger
// to be created with New() method.
type LoggerOption func(*Logger)

// New returns a new Logger
// If any report has ben added the log reporter will be the default
func New(opts ...LoggerOption) *Logger {
	logger := &Logger{
		reporters: []reporter.Reporter{log.NewReporter(nil)},
	}

	for _, opt := range opts {
		opt(logger)
	}

	return logger
}

func WithReporters(reporters ...reporter.Reporter) LoggerOption {
	return func(l *Logger) {
		l.reporters = reporters
	}
}

func (l *Logger) Print(fields ...Field) {
	canonicalLine := format(fields...)
	for _, r := range l.reporters {
		r.Send(canonicalLine)
	}
}
