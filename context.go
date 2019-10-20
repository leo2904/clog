package clog

import (
	"context"
)

var contextKeyCanonicalLine = contextKey("canonical-log-line")

type contextKey string

func newContext(ctx context.Context, l *Line) context.Context {
	return context.WithValue(ctx, contextKeyCanonicalLine, l)
}

func lineFromContext(ctx context.Context) *Line {
	if l, ok := ctx.Value(contextKeyCanonicalLine).(*Line); ok {
		return l
	}

	return nil
}
