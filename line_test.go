package clog

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/friendsofgo/clog/reporter"
	"github.com/friendsofgo/clog/reporter/mock"
)

func TestLine_AddTag(t *testing.T) {
	l := &Line{
		tags: make(map[string]Tag),
	}

	tag := String("key_test", "test")
	l.AddTag(tag)

	if _, ok := l.tags["key_test"]; !ok {
		t.Errorf("AddTag() want %v but is not found", tag)
	}
}

func TestLine_AddTags(t *testing.T) {
	l := &Line{
		tags: make(map[string]Tag),
	}

	tag := String("key_test", "test")
	tag2 := Int("key_test2", 1)
	l.AddTags(tag, tag2)

	if _, ok := l.tags["key_test"]; !ok {
		t.Errorf("AddTag() want %v but is not found", tag)
	}

	if _, ok := l.tags["key_test2"]; !ok {
		t.Errorf("AddTag() want %v but is not found", tag2)
	}
}

func TestLine_OpenSpan(t *testing.T) {
	type fields struct {
		spans map[string]*Span
	}
	type args struct {
		key  string
		flag uint8
	}

	expectedSpan := &Span{
		key:  "span_exists",
		flag: DurationSpan,
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Span
	}{
		{
			"open a new span",
			fields{spans: make(map[string]*Span)},
			args{"span_test", IncrementalSpan | DurationSpan},
			&Span{key: "span_test", flag: IncrementalSpan | DurationSpan},
		},
		{"return an existing span",
			fields{spans: map[string]*Span{"span_exists": expectedSpan}},
			args{"span_exists", DurationSpan},
			expectedSpan,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				spans: tt.fields.spans,
			}
			if got := l.OpenSpan(tt.args.key, tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenSpan() = %v, want %v", got.key, tt.want.key)
			}
		})
	}
}

func TestLine_MarkLineAs(t *testing.T) {
	tests := []struct {
		name string
		arg  *Line
		want reporter.Severity
	}{
		{"mark line as Info", New().NewLine().MarkAsInfo(), reporter.SeverityInfo},
		{"mark line as Error", New().NewLine().MarkAsError(), reporter.SeverityError},
		{"mark line as Critical", New().NewLine().MarkAsCritical(), reporter.SeverityCritical},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.arg.severity != tt.want {
				t.Errorf("MarkAs() want = %v, got = %v", tt.want, tt.arg.severity)
			}
		})
	}
}

func TestLine_Send(t *testing.T) {
	r := &mock.Reporter{
		Sendfn: func(fmtLine string, severity reporter.Severity) {},
	}

	logger := New(WithReporters(r))
	line := logger.NewLine()

	line.Send()
	if !r.SendInvoked {
		t.Errorf("Send() must be called to reporter Send() but this not happened")
	}
}

func TestLine_format(t *testing.T) {
	type fields struct {
		tags  map[string]Tag
		spans map[string]*Span
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"if the line has not spans nor tags must be return empty string", fields{}, ""},
		{"line with a tag", fields{
			tags: map[string]Tag{
				"key1": String("string-tag", "test"),
			},
		},
			"canonical-log-line string-tag=test"},
		{
			"line with a span", fields{
			spans: map[string]*Span{
				"key1": &Span{key: "test-process", flag: DurationSpan|IncrementalSpan, transactions: []Transaction{{duration: 1 * time.Millisecond}}},
			},
		},
			"canonical-log-line test-process_duration=0.001000 test-process_total=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Line{
				tags:  tt.fields.tags,
				spans: tt.fields.spans,
			}
			if got := l.format(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("format() got = %v, want %v", got, tt.want)
			}
		})
	}
}
