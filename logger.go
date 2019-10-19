package clog

import (
	"context"

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

// WithReporters allow to add multiples reporters to our logger
func WithReporters(reporters ...reporter.Reporter) LoggerOption {
	return func(l *Logger) {
		l.reporters = reporters
	}
}

// LineFromContext returns a canonical line using the line found in
// context. If not found a line in context then new line is created.
// The lines created is mark as severity Info for reporter.
func (l *Logger) LineFromContext(ctx context.Context) (*Line, context.Context) {
	if ctxLine := lineFromContext(ctx); ctxLine != nil {
		return ctxLine, ctx
	}

	line := l.newLine()
	return line, newContext(ctx, line)
}

func (l *Logger) newLine() *Line {
	return &Line{
		logger: l,
		severity: reporter.SeverityInfo,
		tags:   make(map[string]Tag),
		spans:  make(map[string]*Span),
	}
}

func (l *Logger) send(canonicalMsg string, severity reporter.Severity) {
	for _, r := range l.reporters {
		r.Send(canonicalMsg, severity)
	}
}
