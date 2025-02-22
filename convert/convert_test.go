package convert

import (
	"encoding/json"
	"fmt"
	"html/template"
	"testing"
	"time"

	"github.com/facebookgo/ensure"
	"github.com/stretchr/testify/assert"
)

type Int64T int64
type Int32T int32
type Int16T int16
type Int8T int8
type Uint64T int64
type Uint32T int32
type Uint16T int16
type Uint8T int8
type IntT int
type UintT uint
type FloatT64 float64
type FloatT32 float32
type StringT string
type BoolT bool

type SInt64T int64

const (
	int64T   Int64T   = 8
	int32T   Int32T   = 8
	int16T   Int16T   = 8
	int8T    Int8T    = 8
	intT     IntT     = 8
	uint64T  Uint64T  = 8
	uint32T  Uint32T  = 8
	uint16T  Uint16T  = 8
	uint8T   Uint8T   = 8
	uintT    UintT    = 8
	floatT64 FloatT64 = 8.0
	floatT32 FloatT32 = 8.0
	stringT  StringT  = "8"
	boolT    BoolT    = true
	sint64T  SInt64T  = 8
)

func (s SInt64T) String() string {
	return "888"
}
func TestToUintE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToUintE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test:
		v = ToUint(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint64
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToUint64E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test:
		v = ToUint64(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint32E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint32
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToUint32E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test:
		v = ToUint32(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint16E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint16
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToUint16E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToUint16(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToUint8E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint8
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToUint8E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToUint8(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToIntE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToIntE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToInt(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToInt64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int64
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{int64(6716803370692116486), 6716803370692116486, false},
		{"6716803370692116486", 6716803370692116486, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToInt64E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToInt64(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToInt32E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int32
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToInt32E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToInt32(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToInt16E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int16
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToInt16E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToInt16(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToInt8E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int8
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"09", 9, false},
		{"0009", 9, false},
		{"0o9", 9, true},
		{nil, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToInt8E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToInt8(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToFloat64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float64
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8), 8, false},
		{float64(8.31), 8.31, false},
		{"8", 8, false},
		{true, 1, false},
		{false, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToFloat64E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToFloat64(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToFloat32E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float32
		iserr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8.31, false},
		{float64(8.31), 8.31, false},
		{"8", 8, false},
		{true, 1, false},
		{false, 0, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
		{int64T, 8, false},
		{int32T, 8, false},
		{int16T, 8, false},
		{int8T, 8, false},
		{intT, 8, false},
		{uint64T, 8, false},
		{uint32T, 8, false},
		{uint16T, 8, false},
		{uint8T, 8, false},
		{uintT, 8, false},
		{floatT64, 8, false},
		{floatT64, 8, false},
		{stringT, 8, false},
		{boolT, 1, false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToFloat32E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToFloat32(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToStringE(t *testing.T) {
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	tests := []struct {
		input  interface{}
		expect string
		iserr  bool
	}{
		{int(8), "8", false},
		{int8(8), "8", false},
		{int16(8), "8", false},
		{int32(8), "8", false},
		{int64(8), "8", false},
		{uint(8), "8", false},
		{uint8(8), "8", false},
		{uint16(8), "8", false},
		{uint32(8), "8", false},
		{uint64(8), "8", false},
		{float32(8.31), "8.31", false},
		{float64(8.31), "8.31", false},
		{true, "true", false},
		{false, "false", false},
		{nil, "", false},
		{[]byte("one time"), "one time", false},
		{"one more time", "one more time", false},
		{template.HTML("one time"), "one time", false},
		{template.URL("http://somehost.foo"), "http://somehost.foo", false},
		{template.JS("(1+2)"), "(1+2)", false},
		{template.CSS("a"), "a", false},
		{template.HTMLAttr("a"), "a", false},
		// errors
		{testing.T{}, "", true},
		{key, "", true},
		{int64T, "8", false},
		{int32T, "8", false},
		{int16T, "8", false},
		{int8T, "8", false},
		{intT, "8", false},
		{uint64T, "8", false},
		{uint32T, "8", false},
		{uint16T, "8", false},
		{uint8T, "8", false},
		{uintT, "8", false},
		{floatT64, "8", false},
		{floatT64, "8", false},
		{stringT, "8", false},
		{boolT, "true", false},
		{sint64T, "888", false},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToStringE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToString(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

type foo struct {
	val string
}

func (x foo) String() string {
	return x.val
}

func TestStringerToString(t *testing.T) {
	var x foo
	x.val = "bar"
	assert.Equal(t, "bar", ToString(x))
}

type fu struct {
	val string
}

func (x fu) Error() string {
	return x.val
}

func TestErrorToString(t *testing.T) {
	var x fu
	x.val = "bar"
	assert.Equal(t, "bar", ToString(x))
}

func TestStringMapStringSliceE(t *testing.T) {
	// ToStringMapString inputs/outputs
	var stringMapString = map[string]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var stringMapInterface = map[string]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapString = map[interface{}]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapInterface = map[interface{}]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}

	// ToStringMapStringSlice inputs/outputs
	var stringMapStringSlice = map[string][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var stringMapInterfaceSlice = map[string][]interface{}{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var stringMapInterfaceInterfaceSlice = map[string]interface{}{"key 1": []interface{}{"value 1", "value 2", "value 3"}, "key 2": []interface{}{"value 1", "value 2", "value 3"}, "key 3": []interface{}{"value 1", "value 2", "value 3"}}
	var stringMapStringSingleSliceFieldsResult = map[string][]string{"key 1": {"value", "1"}, "key 2": {"value", "2"}, "key 3": {"value", "3"}}
	var interfaceMapStringSlice = map[interface{}][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var interfaceMapInterfaceSlice = map[interface{}][]interface{}{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}

	var stringMapStringSliceMultiple = map[string][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}

	var stringMapInterface1 = map[string]interface{}{"key 1": []string{"value 1"}, "key 2": []string{"value 2"}}
	var stringMapInterfaceResult1 = map[string][]string{"key 1": {"value 1"}, "key 2": {"value 2"}}

	var jsonStringMapString = `{"key 1": "value 1", "key 2": "value 2"}`
	var jsonStringMapStringArray = `{"key 1": ["value 1"], "key 2": ["value 2", "value 3"]}`
	var jsonStringMapStringArrayResult = map[string][]string{"key 1": {"value 1"}, "key 2": {"value 2", "value 3"}}

	type Key struct {
		k string
	}

	tests := []struct {
		input  interface{}
		expect map[string][]string
		iserr  bool
	}{
		{stringMapStringSlice, stringMapStringSlice, false},
		{stringMapInterfaceSlice, stringMapStringSlice, false},
		{stringMapInterfaceInterfaceSlice, stringMapStringSlice, false},
		{stringMapStringSliceMultiple, stringMapStringSlice, false},
		{stringMapStringSliceMultiple, stringMapStringSlice, false},
		{stringMapString, stringMapStringSingleSliceFieldsResult, false},
		{stringMapInterface, stringMapStringSingleSliceFieldsResult, false},
		{stringMapInterface1, stringMapInterfaceResult1, false},
		{interfaceMapStringSlice, stringMapStringSlice, false},
		{interfaceMapInterfaceSlice, stringMapStringSlice, false},
		{interfaceMapString, stringMapStringSingleSliceFieldsResult, false},
		{interfaceMapInterface, stringMapStringSingleSliceFieldsResult, false},
		{jsonStringMapStringArray, jsonStringMapStringArrayResult, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{map[interface{}]interface{}{"foo": testing.T{}}, nil, true},
		{map[interface{}]interface{}{Key{"foo"}: "bar"}, nil, true}, // ToStringE(Key{"foo"}) should fail
		{jsonStringMapString, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToStringMapStringSliceE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToStringMapStringSlice(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToStringMapE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]interface{}
		iserr  bool
	}{
		{map[interface{}]interface{}{"tag": "tags", "group": "groups"}, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{map[string]interface{}{"tag": "tags", "group": "groups"}, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{`{"tag": "tags", "group": "groups"}`, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{`{"tag": "tags", "group": true}`, map[string]interface{}{"tag": "tags", "group": true}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToStringMapE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToStringMap(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToStringMapBoolE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]bool
		iserr  bool
	}{
		{map[interface{}]interface{}{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{map[string]interface{}{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{map[string]bool{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{`{"v1": true, "v2": false}`, map[string]bool{"v1": true, "v2": false}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToStringMapBoolE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToStringMapBool(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToStringMapIntE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]int
		iserr  bool
	}{
		{map[interface{}]interface{}{"v1": 1, "v2": 222}, map[string]int{"v1": 1, "v2": 222}, false},
		{map[string]interface{}{"v1": 342, "v2": 5141}, map[string]int{"v1": 342, "v2": 5141}, false},
		{map[string]int{"v1": 33, "v2": 88}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]int32{"v1": int32(33), "v2": int32(88)}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]uint16{"v1": uint16(33), "v2": uint16(88)}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]float64{"v1": float64(8.22), "v2": float64(43.32)}, map[string]int{"v1": 8, "v2": 43}, false},
		{`{"v1": 67, "v2": 56}`, map[string]int{"v1": 67, "v2": 56}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToStringMapIntE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToStringMapInt(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToStringMapInt64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]int64
		iserr  bool
	}{
		{map[interface{}]interface{}{"v1": int32(8), "v2": int32(888)}, map[string]int64{"v1": int64(8), "v2": int64(888)}, false},
		{map[string]interface{}{"v1": int64(45), "v2": int64(67)}, map[string]int64{"v1": 45, "v2": 67}, false},
		{map[string]int64{"v1": 33, "v2": 88}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]int{"v1": 33, "v2": 88}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]int32{"v1": int32(33), "v2": int32(88)}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]uint16{"v1": uint16(33), "v2": uint16(88)}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]float64{"v1": float64(8.22), "v2": float64(43.32)}, map[string]int64{"v1": 8, "v2": 43}, false},
		{`{"v1": 67, "v2": 56}`, map[string]int64{"v1": 67, "v2": 56}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToStringMapInt64E(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToStringMapInt64(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToStringMapStringE(t *testing.T) {
	var stringMapString = map[string]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var stringMapInterface = map[string]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapString = map[interface{}]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapInterface = map[interface{}]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var jsonString = `{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}`
	var invalidJsonString = `{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"`
	var emptyString = ""

	tests := []struct {
		input  interface{}
		expect map[string]string
		iserr  bool
	}{
		{stringMapString, stringMapString, false},
		{stringMapInterface, stringMapString, false},
		{interfaceMapString, stringMapString, false},
		{interfaceMapInterface, stringMapString, false},
		{jsonString, stringMapString, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{invalidJsonString, nil, true},
		{emptyString, nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToStringMapStringE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToStringMapString(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToBoolSliceE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []bool
		iserr  bool
	}{
		{[]bool{true, false, true}, []bool{true, false, true}, false},
		{[]interface{}{true, false, true}, []bool{true, false, true}, false},
		{[]int{1, 0, 1}, []bool{true, false, true}, false},
		{[]int64{1, 0, 1}, []bool{true, false, true}, false},
		{[]int32{1, 0, 1}, []bool{true, false, true}, false},
		{[]uint64{1, 0, 1}, []bool{true, false, true}, false},
		{[]uint32{1, 0, 1}, []bool{true, false, true}, false},
		{[]float64{1, 0, 1}, []bool{true, false, true}, false},
		{[]float32{1, 0, 1}, []bool{true, false, true}, false},
		{[]string{"true", "false", "true"}, []bool{true, false, true}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"foo", "bar"}, nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToBoolSliceE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToBoolSlice(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToIntSliceE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []int
		iserr  bool
	}{
		{[]int{1, 3}, []int{1, 3}, false},
		{[]interface{}{1.2, 3.2}, []int{1, 3}, false},
		{[]string{"2", "3"}, []int{2, 3}, false},
		{[2]string{"2", "3"}, []int{2, 3}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"foo", "bar"}, nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToIntSliceE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToIntSlice(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToSliceE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []interface{}
		iserr  bool
	}{
		{[]interface{}{1, 3}, []interface{}{1, 3}, false},
		{[]map[string]interface{}{{"k1": 1}, {"k2": 2}}, []interface{}{map[string]interface{}{"k1": 1}, map[string]interface{}{"k2": 2}}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToSliceE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToSlice(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToStringSliceE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []string
		iserr  bool
	}{
		{[]string{"a", "b"}, []string{"a", "b"}, false},
		{[]interface{}{1, 3}, []string{"1", "3"}, false},
		{interface{}(1), []string{"1"}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToStringSliceE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToStringSlice(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToDurationSliceE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []time.Duration
		iserr  bool
	}{
		{[]string{"1s", "1m"}, []time.Duration{time.Second, time.Minute}, false},
		{[]int{1, 2}, []time.Duration{1, 2}, false},
		{[]interface{}{1, 3}, []time.Duration{1, 3}, false},
		{[]time.Duration{1, 3}, []time.Duration{1, 3}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"invalid"}, nil, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToDurationSliceE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToDurationSlice(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToBoolE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect bool
		iserr  bool
	}{
		{0, false, false},
		{nil, false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{false, false, false},

		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{true, true, false},
		{-1, true, false},

		// errors
		{"test", false, true},
		{testing.T{}, false, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToBoolE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToBool(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func BenchmarkTooBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !ToBool(true) {
			b.Fatal("ToBool returned false")
		}
	}
}

func TestIndirectPointers(t *testing.T) {
	x := 13
	y := &x
	z := &y

	assert.Equal(t, ToInt(y), 13)
	assert.Equal(t, ToInt(z), 13)
}

func TestToTimeEE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect time.Time
		iserr  bool
	}{
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},   // Time.String()
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},        // ANSIC
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},    // UnixDate
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},  // RubyDate
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},             // RFC822
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},           // RFC822Z
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false}, // RFC850
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},   // RFC1123
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false}, // RFC1123Z
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},            // RFC3339
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, time.UTC), false},      // RFC3339 without timezone hh:mm colon
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},            // RFC3339Nano
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC), false},                              // Kitchen
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},                    // Stamp
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},                // StampMilli
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},             // StampMicro
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},          // StampNano
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},        // RFC3339 without T
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},         // RFC3339 without T or timezone hh:mm colon
		{"2016-03-06 15:28:01", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2006-01-02", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"02 Jan 2006", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{1472574600, time.Date(2016, 8, 30, 16, 30, 0, 0, time.UTC), false},
		{int(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{int64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{int32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{uint64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		// errors
		{"2006", time.Time{}, true},
		{testing.T{}, time.Time{}, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToTimeE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v.UTC(), errmsg)

		// Non-E test
		v = ToTime(test.input)
		assert.Equal(t, test.expect, v.UTC(), errmsg)
	}
}

func TestToDurationE(t *testing.T) {
	var td time.Duration = 5

	tests := []struct {
		input  interface{}
		expect time.Duration
		iserr  bool
	}{
		{time.Duration(5), td, false},
		{int(5), td, false},
		{int64(5), td, false},
		{int32(5), td, false},
		{int16(5), td, false},
		{int8(5), td, false},
		{uint(5), td, false},
		{uint64(5), td, false},
		{uint32(5), td, false},
		{uint16(5), td, false},
		{uint8(5), td, false},
		{float64(5), td, false},
		{float32(5), td, false},
		{string("5"), td, false},
		{string("5ns"), td, false},
		{string("5us"), time.Microsecond * td, false},
		{string("5µs"), time.Microsecond * td, false},
		{string("5ms"), time.Millisecond * td, false},
		{string("5s"), time.Second * td, false},
		{string("5m"), time.Minute * td, false},
		{string("5h"), time.Hour * td, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for i, test := range tests {
		errmsg := fmt.Sprintf("i = %d", i) // assert helper message

		v, err := ToDurationE(test.input)
		if test.iserr {
			assert.Error(t, err, errmsg)
			continue
		}

		assert.NoError(t, err, errmsg)
		assert.Equal(t, test.expect, v, errmsg)

		// Non-E test
		v = ToDuration(test.input)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestToStringSlice(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3", "4"}, ToStringSlice("1 2 3 4"))
	assert.Equal(t, []string{"1", "2", "3", "4"}, ToStringSlice([]int{1, 2, 3, 4}))
	assert.Equal(t, []string{"1.2"}, ToStringSlice(1.2))

}

func TestToIntSlice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, ToIntSlice([]string{"1", "2", "3", "4"}))
}

func TestToSlice(t *testing.T) {
	assert.Equal(t, []interface{}{}, ToSlice("1 2 3 4"))
	assert.Equal(t, []interface{}{"1", "2", "3", "4"}, ToSlice([]string{"1", "2", "3", "4"}))
	assert.Equal(t, []interface{}{1, 2, 3, 4}, ToSlice([]int{1, 2, 3, 4}))
	assert.Equal(t, []interface{}{}, ToSlice(1.2))
}

func TestToMobile(t *testing.T) {
	assert.Equal(t, "15821446338", ToMobile("+8615821446338"))
	assert.Equal(t, "15821446338", ToMobile("15821446338"))
	assert.Equal(t, "15821446338", ToMobile(15821446338))
	assert.Equal(t, "", ToMobile(315821446338))
	assert.Equal(t, "", ToMobile("315821446338"))
}

func TestToVersion(t *testing.T) {
	versionInfo := ToVersion("3.2.1beta")
	ensure.DeepEqual(t, versionInfo.Major, 3)
	ensure.DeepEqual(t, versionInfo.Minor, 2)
	ensure.DeepEqual(t, versionInfo.Build, 1)

	versionInfo = ToVersion("3.2.1-beta")
	ensure.DeepEqual(t, versionInfo.Major, 3)
	ensure.DeepEqual(t, versionInfo.Minor, 2)
	ensure.DeepEqual(t, versionInfo.Build, 1)

	versionInfo = ToVersion(versionInfo)
	ensure.DeepEqual(t, versionInfo.Major, 3)
	ensure.DeepEqual(t, versionInfo.Minor, 2)
	ensure.DeepEqual(t, versionInfo.Build, 1)

	versionInfo, err := ToVersionE(nil)
	ensure.DeepEqual(t, versionInfo.Major, 0)
	ensure.DeepEqual(t, versionInfo.Minor, 0)
	ensure.DeepEqual(t, versionInfo.Build, 0)
	ensure.NotNil(t, err)

	versionInfo = ToVersion(nil)
	ensure.DeepEqual(t, versionInfo.Major, 0)
	ensure.DeepEqual(t, versionInfo.Minor, 0)
	ensure.DeepEqual(t, versionInfo.Build, 0)

	versionInfo = ToVersion("kitool - version v3.7.21\n")
	ensure.DeepEqual(t, versionInfo.Major, 3)
	ensure.DeepEqual(t, versionInfo.Minor, 7)
	ensure.DeepEqual(t, versionInfo.Build, 21)

	versionInfo = ToVersion("v1.4.2")
	ensure.DeepEqual(t, versionInfo.Major, 1)
	ensure.DeepEqual(t, versionInfo.Minor, 4)
	ensure.DeepEqual(t, versionInfo.Build, 2)
	ensure.DeepEqual(t, fmt.Sprintf("%v", versionInfo), "1.4.2")
	ensure.DeepEqual(t, fmt.Sprintf("%v", *versionInfo), "1.4.2")

}

func TestMapToKeySlice(t *testing.T) {
	src := map[string]struct{}{
		"1": struct{}{},
		"2": struct{}{},
	}
	ensure.DeepEqual(t, len(MapToKeySlice(src)), 2)
}

func TestMapToValueSlice(t *testing.T) {
	src := map[string]string{
		"1": "1",
		"2": "1",
	}
	ensure.DeepEqual(t, len(MapToValueSlice(src)), 2)
}

type BB struct {
	A int `json:"a"`
}

func (b BB) MarshalBinary() ([]byte, error) {
	return []byte("BB"), nil
}
func (b *BB) UnmarshalBinary([]byte) error {
	b.A = 1
	return nil
}

type CC struct {
	A int `json:"a"`
}

func TestToByteSlice(t *testing.T) {
	src := []int{1, 2, 3}
	val, _ := json.Marshal(src)
	ensure.DeepEqual(t, ToByteSlice(src), val)
	ensure.DeepEqual(t, ToByteSlice(&BB{}), []byte("BB"))
	ensure.DeepEqual(t, ToByteSlice(&CC{A: 1}), []byte(`{"a":1}`))
}

func TestToStructData(t *testing.T) {
	bb := &BB{}
	cc := &CC{}
	BinaryToStructData([]byte("BB"), bb)
	BinaryToStructData([]byte(`{"a":2}`), cc)
	ensure.DeepEqual(t, bb.A, 1)
	ensure.DeepEqual(t, cc.A, 2)
}

func TestStructToSlice(t *testing.T) {
	bb := []BB{
		{
			A: 1,
		},
		{
			A: 2,
		},
		{
			A: 3,
		},
		{
			A: 4,
		},
	}
	strSlice := ExtractStructToStringSlice(bb, func(i int) string {
		return ToString(bb[i].A)
	})
	ensure.DeepEqual(t, strSlice, []string{"1", "2", "3", "4"})
	int64Slice := ExtractStructToInt64Slice(bb, func(i int) int64 {
		return int64(bb[i].A)
	})
	ensure.DeepEqual(t, int64Slice, []int64{1, 2, 3, 4})
	int32Slice := ExtractStructToInt32Slice(bb, func(i int) int32 {
		return int32(bb[i].A)
	})
	ensure.DeepEqual(t, int32Slice, []int32{1, 2, 3, 4})
	int8Slice := ExtractStructToInt8Slice(bb, func(i int) int8 {
		return int8(bb[i].A)
	})
	ensure.DeepEqual(t, int8Slice, []int8{1, 2, 3, 4})
	uint64Slice := ExtractStructToUint64Slice(bb, func(i int) uint64 {
		return uint64(bb[i].A)
	})
	ensure.DeepEqual(t, uint64Slice, []uint64{1, 2, 3, 4})
	uint32Slice := ExtractStructToUint32Slice(bb, func(i int) uint32 {
		return uint32(bb[i].A)
	})
	ensure.DeepEqual(t, uint32Slice, []uint32{1, 2, 3, 4})
	uint8Slice := ExtractStructToUint8Slice(bb, func(i int) uint8 {
		return uint8(bb[i].A)
	})
	ensure.DeepEqual(t, uint8Slice, []uint8{1, 2, 3, 4})
}
