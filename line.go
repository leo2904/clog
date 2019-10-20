package clog

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/clog/reporter"
)

const (
	_defaultTimestampFormat = time.RFC3339
	_headerLineFormat       = "[%s] canonical-log-line "
)

// Line represent a Canonical Log Lines in which is accumulated all the information about the process.
type Line struct {
	sync.RWMutex

	logger *Logger
	severity reporter.Severity
	tags   map[string]Tag
	spans  map[string]*Span
}

// AddTag add new tag to the line.
func (l *Line) AddTag(t Tag) {
	l.Lock()
	l.tags[t.key] = t
	l.Unlock()
}

// AddTags add multiples tags to the line at once.
func (l *Line) AddTags(tags ...Tag) {
	for _, t := range tags {
		l.AddTag(t)
	}
}

// OpenSpan return an existing span into the line otherwise creates a new one.
// You need to specify which is the purpose of the Span with the flags:
// IncrementalSpan or DurationSpan or both.
func (l *Line) OpenSpan(key string, flag uint8) *Span {
	l.Lock()

	if sp, ok := l.spans[key]; ok {
		l.Unlock()
		return sp
	}

	s := &Span{key: key, flag: flag}
	l.spans[key] = s
	l.Unlock()

	return s
}

// MarkAsInfo change the severity of the line to INFO.
func (l *Line) MarkAsInfo() *Line {
	l.Lock()
	l.severity = reporter.SeverityInfo
	l.Unlock()
	return l
}

// MarkAsError change the severity of the line to ERROR.
func (l *Line) MarkAsError() *Line {
	l.Lock()
	l.severity = reporter.SeverityError
	l.Unlock()
	return l
}

// MarkAsCritical change the severity of the line to CRITICAL.
func (l *Line) MarkAsCritical() *Line {
	l.Lock()
	l.severity = reporter.SeverityCritical
	l.Unlock()
	return l
}

// Send preparing and send the line throw to all reporters initialized on the logger.
func (l *Line) Send() {
	l.send()
}

func (l *Line) send() {
	l.Lock()
	canonicalLine := l.format()
	l.logger.send(canonicalLine, l.severity)
	l.Unlock()
}

func (l *Line) format() string {
	var b bytes.Buffer

	totalTags := len(l.tags)
	totalSpans := len(l.spans)

	if totalTags == 0 && totalSpans == 0 {
		return ""
	}

	current := time.Now().Format(_defaultTimestampFormat)
	b.WriteString(fmt.Sprintf(_headerLineFormat, current))

	for _, t := range l.tags {
		b.WriteString(t.String() + " ")
	}

	for _, s := range l.spans {
		b.WriteString(s.String())
	}

	return strings.TrimSpace(b.String())
}
