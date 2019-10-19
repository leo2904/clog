package clog

import (
	"bytes"
	"fmt"
	"time"
)
const (
	_defaultTimestampFormat = time.RFC3339
	_headerLineFormat = "[%s] canonical-log-line "
)

func formatHeader() string {
	current := time.Now().Format(_defaultTimestampFormat)
	return fmt.Sprintf(_headerLineFormat, current)
}

func format(fields ...Field) string {
	var b bytes.Buffer
	b.WriteString(formatHeader())

	for _, f := range fields {
		switch f.fType {
		case stringType:
			msg := fmt.Sprintf("%s=%s ", f.key, f.stringVal)
			b.WriteString(msg)
		case boolType:
			msg := fmt.Sprintf("%s=%t ", f.key, f.boolVal)
			b.WriteString(msg)
		case int8Type,
			int16Type,
			int32Type,
			int64Type,
			uint8Type,
			uint16Type,
			uint32Type,
			uint64Type:
			msg := fmt.Sprintf("%s=%d ", f.key, f.integerVal)
			b.WriteString(msg)
		case float32Type,
			float64Type:
			msg := fmt.Sprintf("%s=%f ", f.key, f.floatVal)
			b.WriteString(msg)
		case errorType,
			anyType:
			msg := fmt.Sprintf("%s=%v ", f.key, f.anyVal)
			b.WriteString(msg)
		}
	}

	return b.String()
}
