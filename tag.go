package clog

import (
	"fmt"
)

// TagType indicates which property of the Tag struct should be used
type TagType uint8

const (
	// unknownType is the default tag type. If unknownType is received the behaviour will be the same of Any.
	unknownType TagType = iota
	// stringType indicates that the tag carries a string.
	stringType
	//boolType indicates that the tag carries a bool.
	boolType
	// int8Type indicates that the tag carries an int8.
	int8Type
	// int16Type indicates that the tag carries an int16.
	int16Type
	// int32Type indicates that the tag carries an int32.
	int32Type
	// int64Type indicates that the tag carries an int64.
	int64Type
	// float32Type indicates that the tag carries a float32.
	float32Type
	// float64Type indicates that the tag carries a float64.
	float64Type
	// uint8Type indicates that the tag carries an uint8.
	uint8Type
	// uint16Type indicates that the tag carries an uint16.
	uint16Type
	// uint32Type indicates that the tag carries an uint32.
	uint32Type
	// uint64Type indicates that the tag carries an uint64.
	uint64Type
	// errorType indicates that the tag carriers an error.
	errorType
	// anyType indicates that the tag carries a interface{}.
	anyType
)

// Tag represent a key-value pair to a logger's context.
type Tag struct {
	key   string
	tType TagType

	stringVal  string
	integerVal int64
	floatVal   float64
	boolVal    bool
	anyVal     interface{}
}

// String transform tag to `logFmt format (https://brandur.org/logfmt)
func (t Tag) String() string {
	switch t.tType {
	case stringType:
		return fmt.Sprintf("%s=%s", t.key, t.stringVal)
	case boolType:
		return fmt.Sprintf("%s=%t", t.key, t.boolVal)
	case int8Type,
		int16Type,
		int32Type,
		int64Type,
		uint8Type,
		uint16Type,
		uint32Type,
		uint64Type:
		return fmt.Sprintf("%s=%d", t.key, t.integerVal)
	case float32Type,
		float64Type:
		return fmt.Sprintf("%s=%f ", t.key, t.floatVal)
	case errorType,
		anyType:
		return fmt.Sprintf("%s=%v ", t.key, t.anyVal)
	default:
		return ""
	}
}

// String constructs a tag that carries a string.
func String(key, val string) Tag {
	return Tag{key: key, tType: stringType, stringVal: val}
}

// Bool constructs a tag that carries a bool.
func Bool(key string, val bool) Tag {
	return Tag{key: key, tType: boolType, boolVal: val}
}

// Int constructs a tag that carries an int64.
func Int(key string, val int) Tag {
	return Int64(key, int64(val))
}

// Int8 constructs a tag that carries an int8.
func Int8(key string, val int8) Tag {
	return Tag{key: key, tType: int8Type, integerVal: int64(val)}
}

// Int16 constructs a tag that carries an int16.
func Int16(key string, val int16) Tag {
	return Tag{key: key, tType: int16Type, integerVal: int64(val)}
}

// Int32 constructs a tag that carries an int32.
func Int32(key string, val int32) Tag {
	return Tag{key: key, tType: int32Type, integerVal: int64(val)}
}

// Int64 constructs a tag that carries an int64.
func Int64(key string, val int64) Tag {
	return Tag{key: key, tType: int64Type, integerVal: val}
}

// Float32 constructs a tag that carries a float32.
func Float32(key string, val float32) Tag {
	return Tag{key: key, tType: float32Type, floatVal: float64(val)}
}

// Float64 constructs a tag that carries a float64.
func Float64(key string, val float64) Tag {
	return Tag{key: key, tType: float64Type, floatVal: val}
}

// Uint constructs a tag that carries an int64.
func Uint(key string, val uint) Tag {
	return Uint64(key, uint64(val))
}

// Uint8 constructs a tag that carries an int8.
func Uint8(key string, val uint8) Tag {
	return Tag{key: key, tType: uint8Type, integerVal: int64(val)}
}

// Uint16 constructs a tag that carries an int16.
func Uint16(key string, val uint16) Tag {
	return Tag{key: key, tType: uint16Type, integerVal: int64(val)}
}

// Uint32 constructs a tag that carries a int32.
func Uint32(key string, val uint32) Tag {
	return Tag{key: key, tType: uint32Type, integerVal: int64(val)}
}

// Uint64 constructs a tag that carries an int64.
func Uint64(key string, val uint64) Tag {
	return Tag{key: key, tType: uint64Type, integerVal: int64(val)}
}

// Error constructs a tag that carries an error.
func Error(key string, val error) Tag {
	return Tag{key: key, tType: errorType, anyVal: val}
}

// Any constructs a tag with any value.
func Any(key string, val interface{}) Tag {
	switch value := val.(type) {
	case string:
		return String(key, value)
	case bool:
		return Bool(key, value)
	case int8:
		return Int8(key, value)
	case int16:
		return Int16(key, value)
	case int32:
		return Int32(key, value)
	case int64:
		return Int64(key, value)
	case uint8:
		return Uint8(key, value)
	case uint16:
		return Uint16(key, value)
	case uint32:
		return Uint32(key, value)
	case uint64:
		return Uint64(key, value)
	case float32:
		return Float32(key, value)
	case float64:
		return Float64(key, value)
	case error:
		return Error(key, value)
	default:
		return Tag{key: key, tType: anyType, anyVal: val}
	}
}
