package clog

import (
	"fmt"
	"io"
	"time"
)
const lineFormat = "[%s] canonical-log-line "

func printHeader(out io.Writer) {
	current := time.Now().Format(time.RFC3339)
	_, _ = io.WriteString(out, fmt.Sprintf(lineFormat, current))
}

func printCanonicalLine(out io.Writer, fields ...Field) {
	printHeader(out)

	for _, f := range fields {
		switch f.fType {
		case stringType:
			msg := fmt.Sprintf("%s=%s ", f.key, f.stringVal)
			_, _ = io.WriteString(out, msg)
		case boolType:
			msg := fmt.Sprintf("%s=%t ", f.key, f.boolVal)
			_, _ = io.WriteString(out, msg)
		case int8Type,
			int16Type,
			int32Type,
			int64Type,
			uint8Type,
			uint16Type,
			uint32Type,
			uint64Type:
			msg := fmt.Sprintf("%s=%d ", f.key, f.integerVal)
			_, _ = io.WriteString(out, msg)
		case float32Type,
			float64Type:
			msg := fmt.Sprintf("%s=%f ", f.key, f.floatVal)
			_, _ = io.WriteString(out, msg)
		case errorType,
			anyType:
			msg := fmt.Sprintf("%s=%v ", f.key, f.anyVal)
			_, _ = io.WriteString(out, msg)
		}
	}
}
