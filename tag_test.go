package clog

import (
	"errors"
	"reflect"
	"testing"
)

func TestTag_Constructors(t *testing.T) {
	tests := []struct {
		name string
		want Tag
		got  Tag
	}{
		{"Skip", Tag{tType: skipType}, Skip()},
		{"String", Tag{key: "key_test", tType: stringType, stringVal: "test"}, String("key_test", "test")},
		{"Bool", Tag{key: "key_test", tType: boolType, boolVal: true}, Bool("key_test", true)},
		{"Int", Tag{key: "key_test", tType: int64Type, integerVal: 1}, Int("key_test", 1)},
		{"Int8", Tag{key: "key_test", tType: int8Type, integerVal: 1}, Int8("key_test", 1)},
		{"Int16", Tag{key: "key_test", tType: int16Type, integerVal: 1}, Int16("key_test", 1)},
		{"Int32", Tag{key: "key_test", tType: int32Type, integerVal: 1}, Int32("key_test", 1)},
		{"Int64", Tag{key: "key_test", tType: int64Type, integerVal: 1}, Int64("key_test", 1)},
		{"Uint", Tag{key: "key_test", tType: uint64Type, integerVal: 1}, Uint("key_test", 1)},
		{"Uint8", Tag{key: "key_test", tType: uint8Type, integerVal: 1}, Uint8("key_test", 1)},
		{"Uint16", Tag{key: "key_test", tType: uint16Type, integerVal: 1}, Uint16("key_test", 1)},
		{"Uint32", Tag{key: "key_test", tType: uint32Type, integerVal: 1}, Uint32("key_test", 1)},
		{"Uint64", Tag{key: "key_test", tType: uint64Type, integerVal: 1}, Uint64("key_test", 1)},
		{"Float32", Tag{key: "key_test", tType: float32Type, floatVal: 1.0}, Float32("key_test", 1.0)},
		{"Float64", Tag{key: "key_test", tType: float64Type, floatVal: 1.0}, Float64("key_test", 1.0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%s got = %v, want = %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestTag_ErrorConstructor(t *testing.T) {

	tests := []struct {
		name string
		want Tag
		got  Tag
	}{
		{"Error:Nil", Skip(), Error("key_test", nil)},
		{"Error", Tag{key: "key_test", tType: errorType, tInterface: errors.New("errors happens")}, Error("key_test", errors.New("errors happens"))},
		{"Any:Error", Error("key_test", errors.New("errors happens")), Any("key_test", errors.New("errors happens"))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%s got = %v, want = %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestTag_AnyConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Tag
		got  Tag
	}{
		{"Any:String", String("key_test", "test"), Any("key_test", "test")},
		{"Any:Bool", Bool("key_test", true), Any("key_test", true)},
		{"Any:Int", Int("key_test", 1), Any("key_test", 1)},
		{"Any:Int8", Int8("key_test", 1), Any("key_test", int8(1))},
		{"Any:Int16", Int16("key_test", 1), Any("key_test", int16(1))},
		{"Any:Int32", Int32("key_test", 1), Any("key_test", int32(1))},
		{"Any:Int64", Int64("key_test", 1), Any("key_test", int64(1))},
		{"Any:Uint", Uint("key_test", 1), Any("key_test", uint(1))},
		{"Any:Uint8", Uint8("key_test", 1), Any("key_test", uint8(1))},
		{"Any:Uint16", Uint16("key_test", 1), Any("key_test", uint16(1))},
		{"Any:Uint32", Uint32("key_test", 1), Any("key_test", uint32(1))},
		{"Any:Uint64", Uint64("key_test", 1), Any("key_test", uint64(1))},
		{"Any:Float32", Float32("key_test", 1.0), Any("key_test", float32(1.0))},
		{"Any:Float64", Float64("key_test", 1.0), Any("key_test", 1.0)},
		{"Any", Tag{key: "key_test", tType: anyType, tInterface: struct{}{}}, Any("key_test", struct{}{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%s got = %T, want = %T", tt.name, tt.got.tInterface, tt.want.tInterface)
			}
		})
	}
}

func TestTag_String(t *testing.T) {
	tests := []struct {
		name string
		tag  Tag
		want string
	}{
		{"Skip", Skip(), ""},
		{"String", String("key_test", "test"), "key_test=test"},
		{"Bool", Bool("key_test", true), "key_test=true"},
		{"Int", Int("key_test", 1), "key_test=1"},
		{"Int8", Int8("key_test", 1), "key_test=1"},
		{"Int16", Int16("key_test", 1), "key_test=1"},
		{"Int32", Int32("key_test", 1), "key_test=1"},
		{"Int64", Int64("key_test", 1), "key_test=1"},
		{"Uint", Uint("key_test", 1), "key_test=1"},
		{"Uint8", Uint8("key_test", 1), "key_test=1"},
		{"Uint16", Uint16("key_test", 1), "key_test=1"},
		{"Uint32", Uint32("key_test", 1), "key_test=1"},
		{"Uint64", Uint64("key_test", 1), "key_test=1"},
		{"Float32", Float32("key_test", 1.0), "key_test=1.000000"},
		{"Float64", Float64("key_test", 1.0), "key_test=1.000000"},
		{"Error:Nil", Error("key_test", nil), ""},
		{"Error", Error("key_test", errors.New("errors happens")), "key_test=errors happens"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tag.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
