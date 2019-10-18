package clog

// FieldType indicates which property of the Field struct should be used
type FieldType uint8

const (
	// unknownType is the default field type. If unknownType is received the behaviour will be the same of Any.
	unknownType FieldType = iota
	// stringType indicates that the field carries a string.
	stringType
	//boolType indicates that the field carries a bool.
	boolType
	// int8Type indicates that the field carries an int8.
	int8Type
	// int16Type indicates that the field carries an int16.
	int16Type
	// int32Type indicates that the field carries an int32.
	int32Type
	// int64Type indicates that the field carries an int64.
	int64Type
	// float32Type indicates that the field carries a float32.
	float32Type
	// float64Type indicates that the field carries a float64.
	float64Type
	// uint8Type indicates that the field carries an uint8.
	uint8Type
	// uint16Type indicates that the field carries an uint16.
	uint16Type
	// uint32Type indicates that the field carries an uint32.
	uint32Type
	// uint64Type indicates that the field carries an uint64.
	uint64Type
	// errorType indicates that the field carriers an error.
	errorType
	// anyType indicates that the field carries a interface{}.
	anyType
)

// Field represent a key-value pair to a logger's context.
type Field struct {
	key   string
	fType FieldType

	stringVal  string
	integerVal int64
	floatVal   float64
	boolVal    bool
	anyVal     interface{}
}

// String constructs a field that carries a string.
func String(key, val string) Field {
	return Field{key: key, fType: stringType, stringVal: val}
}

// Bool constructs a field that carries a bool.
func Bool(key string, val bool) Field {
	return Field{key: key, fType: boolType, boolVal: val}
}

// Int constructs a field that carries an int64.
func Int(key string, val int) Field {
	return Int64(key, int64(val))
}

// Int8 constructs a field that carries an int8.
func Int8(key string, val int8) Field {
	return Field{key: key, fType: int8Type, integerVal: int64(val)}
}

// Int16 constructs a field that carries an int16.
func Int16(key string, val int16) Field {
	return Field{key: key, fType: int16Type, integerVal: int64(val)}
}

// Int32 constructs a field that carries an int32.
func Int32(key string, val int32) Field {
	return Field{key: key, fType: int32Type, integerVal: int64(val)}
}

// Int64 constructs a field that carries an int64.
func Int64(key string, val int64) Field {
	return Field{key: key, fType: int64Type, integerVal: val}
}

// Float32 constructs a field that carries a float32.
func Float32(key string, val float32) Field {
	return Field{key: key, fType: float32Type, floatVal: float64(val)}
}

// Float64 constructs a field that carries a float64.
func Float64(key string, val float64) Field {
	return Field{key: key, fType: float64Type, floatVal: val}
}

// Uint constructs a field that carries an int64.
func Uint(key string, val uint) Field {
	return Uint64(key, uint64(val))
}

// Uint8 constructs a field that carries an int8.
func Uint8(key string, val uint8) Field {
	return Field{key: key, fType: uint8Type, integerVal: int64(val)}
}

// Uint16 constructs a field that carries an int16.
func Uint16(key string, val uint16) Field {
	return Field{key: key, fType: uint16Type, integerVal: int64(val)}
}

// Uint32 constructs a field that carries a int32.
func Uint32(key string, val uint32) Field {
	return Field{key: key, fType: uint32Type, integerVal: int64(val)}
}

// Uint64 constructs a field that carries an int64.
func Uint64(key string, val uint64) Field {
	return Field{key: key, fType: uint64Type, integerVal: int64(val)}
}

// Error constructs a field that carries an error.
func Error(key string, val error) Field {
	return Field{key: key, fType: errorType, anyVal: val}
}

// Any constructs a field with any value.
func Any(key string, val interface{}) Field {
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
		return Field{key: key, fType: anyType, anyVal: val}
	}
}
