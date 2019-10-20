package clog

import (
	"context"
	"reflect"
	"testing"

	"github.com/friendsofgo/clog/reporter"
	"github.com/friendsofgo/clog/reporter/log"
)

func TestNew(t *testing.T) {
	type args struct {
		opts []LoggerOption
	}
	tests := []struct {
		name string
		args args
		want *Logger
	}{
		{
			"initialize logger without reporters must be return a logger with log reporter",
			args{},
			&Logger{[]reporter.Reporter{log.NewReporter(nil)}},
		},
		{
			"initialize logger with reporter must be return a logger with selected reporters",
			args{
				[]LoggerOption{WithReporters(reporter.NewNoop())},
			},
			&Logger{[]reporter.Reporter{reporter.NewNoop()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, wantLine %v", got, tt.want)
			}
		})
	}
}

func TestLogger_NewLineOnContext(t *testing.T) {
	logger := New()
	expectedLine := emptyLine(logger)
	expectedCtx := newContext(context.Background(), expectedLine)

	gotLine, gotCtx := logger.NewLineOnContext(context.Background())
	if !reflect.DeepEqual(gotLine, expectedLine) {
		t.Errorf("NewLineOnContext() gotLine = %v, wantLine %v", gotLine, expectedLine)
	}

	if !reflect.DeepEqual(gotCtx, expectedCtx) {
		t.Errorf("NewLineOnContext() gotCtx = %v, wantLine %v", gotCtx, expectedCtx)
	}
}

func TestLogger_LineFromContext(t *testing.T) {
	logger := New()
	lineWithTags := &Line{
		logger:   logger,
		severity: reporter.SeverityInfo,
		tags:     map[string]Tag{"testTag": String("test", "test")},
		spans:    make(map[string]*Span),
	}
	ctxWithLine := newContext(context.Background(), lineWithTags)

	tests := []struct {
		name     string
		ctx      context.Context
		wantLine *Line
	}{
		{
			"line doesn't exists on context must be return new line",
			context.Background(),
			emptyLine(logger),
		},
		{
			"line exists on context must be return the line on context",
			ctxWithLine,
			lineWithTags,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			got, _ := l.LineFromContext(tt.ctx)
			if !reflect.DeepEqual(got, tt.wantLine) {
				t.Errorf("LineFromContext() got = %v, wantLine %v", got, tt.wantLine)
			}
		})
	}
}

func emptyLine(l *Logger) *Line {
	return &Line{
		logger:   l,
		severity: reporter.SeverityInfo,
		tags:     make(map[string]Tag),
		spans:    make(map[string]*Span),
	}
}
