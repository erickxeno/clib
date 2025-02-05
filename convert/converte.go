package convert

import (
	"encoding"
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/erickxeno/clib/errors"
)

const (
	// match CN phone number
	cnMobileStr = `(0|\+?86)?` + // 匹配 0,86,+86
		`(13[0-9]|` + // 130-139
		`14[0-9]|` + // 140-149
		`15[0-9]|` + // 150-159
		`16[0-9]|` + // 160-169
		`17[0-9]|` + // 170-179
		`18[0-9]|` + // 180-189
		`19[0-9])` + // 190-199
		`[0-9]{8}`
)

var (
	errNegativeNotAllowed = fmt.Errorf("unable to cast negative value")
	cnMobile              = regexp.MustCompile("^" + cnMobileStr + "$")
)

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// Indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func Indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ToBoolE convert interface to bool
func ToBoolE(i interface{}) (bool, error) {
	i = Indirect(i)
	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		return i.(int) != 0, nil
	case int32:
		return i.(int32) != 0, nil
	case int64:
		return i.(int64) != 0, nil
	case uint:
		return i.(uint) != 0, nil
	case uint32:
		return i.(uint32) != 0, nil
	case uint64:
		return i.(uint64) != 0, nil
	case float32:
		return i.(float32) != 0, nil
	case float64:
		return i.(float64) != 0, nil
	case string:
		v, err := strconv.ParseBool(i.(string))
		if err != nil {
			err = errors.WithStack(err)
		}
		return v, err
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToBoolE(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToBoolE(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToBoolE(v.Float())
		case reflect.Bool:
			return v.Bool(), nil
		case reflect.String:
			return ToBoolE(v.String())
		}
		return false, errors.Errorf("unable to cast %#v of type %T to bool", i, i)
	}
}

// ToFloat64E convert interface to float64
func ToFloat64E(i interface{}) (float64, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case uint:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return v, nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v of type %T to float64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToFloat64E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToFloat64E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToFloat64E(v.Float())
		case reflect.Bool:
			return ToFloat64E(v.Bool())
		case reflect.String:
			return ToFloat64E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to float64", i, i)
	}
}

// ToFloat32E convert interface to float32
func ToFloat32E(i interface{}) (float32, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case float64:
		return float32(s), nil
	case float32:
		return s, nil
	case int:
		return float32(s), nil
	case int64:
		return float32(s), nil
	case int32:
		return float32(s), nil
	case int16:
		return float32(s), nil
	case int8:
		return float32(s), nil
	case uint:
		return float32(s), nil
	case uint64:
		return float32(s), nil
	case uint32:
		return float32(s), nil
	case uint16:
		return float32(s), nil
	case uint8:
		return float32(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return float32(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v of type %T to float32", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToFloat32E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToFloat32E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToFloat32E(v.Float())
		case reflect.Bool:
			return ToFloat32E(v.Bool())
		case reflect.String:
			return ToFloat32E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to float32", i, i)
	}
}

// ToInt64E convert interface to int64
func ToInt64E(i interface{}) (int64, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case int:
		return int64(s), nil
	case int64:
		return s, nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case uint:
		return int64(s), nil
	case uint64:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint8:
		return int64(s), nil
	case float64:
		return int64(s), nil
	case float32:
		return int64(s), nil
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToInt64E(s[1:])
		}
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v of type %T to int64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToInt64E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToInt64E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToInt64E(v.Float())
		case reflect.Bool:
			return ToInt64E(v.Bool())
		case reflect.String:
			return ToInt64E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}

// ToInt32E convert interface to int32
func ToInt32E(i interface{}) (int32, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case int:
		return int32(s), nil
	case int64:
		return int32(s), nil
	case int32:
		return s, nil
	case int16:
		return int32(s), nil
	case int8:
		return int32(s), nil
	case uint:
		return int32(s), nil
	case uint64:
		return int32(s), nil
	case uint32:
		return int32(s), nil
	case uint16:
		return int32(s), nil
	case uint8:
		return int32(s), nil
	case float64:
		return int32(s), nil
	case float32:
		return int32(s), nil
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToInt32E(s[1:])
		}
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int32(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v of type %T to int32", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToInt32E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToInt32E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToInt32E(v.Float())
		case reflect.Bool:
			return ToInt32E(v.Bool())
		case reflect.String:
			return ToInt32E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to int32", i, i)
	}
}

// ToInt16E convert interface to int16
func ToInt16E(i interface{}) (int16, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case int:
		return int16(s), nil
	case int64:
		return int16(s), nil
	case int32:
		return int16(s), nil
	case int16:
		return s, nil
	case int8:
		return int16(s), nil
	case uint:
		return int16(s), nil
	case uint64:
		return int16(s), nil
	case uint32:
		return int16(s), nil
	case uint16:
		return int16(s), nil
	case uint8:
		return int16(s), nil
	case float64:
		return int16(s), nil
	case float32:
		return int16(s), nil
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToInt16E(s[1:])
		}
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int16(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v of type %T to int16", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToInt16E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToInt16E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToInt16E(v.Float())
		case reflect.Bool:
			return ToInt16E(v.Bool())
		case reflect.String:
			return ToInt16E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
}

// ToInt8E convert interface to int8
func ToInt8E(i interface{}) (int8, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case int:
		return int8(s), nil
	case int64:
		return int8(s), nil
	case int32:
		return int8(s), nil
	case int16:
		return int8(s), nil
	case int8:
		return s, nil
	case uint:
		return int8(s), nil
	case uint64:
		return int8(s), nil
	case uint32:
		return int8(s), nil
	case uint16:
		return int8(s), nil
	case uint8:
		return int8(s), nil
	case float64:
		return int8(s), nil
	case float32:
		return int8(s), nil
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToInt8E(s[1:])
		}
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int8(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v of type %T to int8", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToInt8E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToInt8E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToInt8E(v.Float())
		case reflect.Bool:
			return ToInt8E(v.Bool())
		case reflect.String:
			return ToInt8E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}

// ToIntE convert interface to int
func ToIntE(i interface{}) (int, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case int:
		return s, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		return int(s), nil
	case float32:
		return int(s), nil
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToIntE(s[1:])
		}
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v of type %T to int", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToIntE(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToIntE(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToIntE(v.Float())
		case reflect.Bool:
			return ToIntE(v.Bool())
		case reflect.String:
			return ToIntE(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

// ToUintE convert interface to uint
func ToUintE(i interface{}) (uint, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToUintE(s[1:])
		}
		v, err := strconv.ParseUint(s, 0, 0)
		if err == nil {
			return uint(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v to uint", i)
	case int:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint(s), nil
	case int64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint(s), nil
	case int32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint(s), nil
	case int16:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint(s), nil
	case int8:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint(s), nil
	case uint:
		return s, nil
	case uint64:
		return uint(s), nil
	case uint32:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint8:
		return uint(s), nil
	case float64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint(s), nil
	case float32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToUintE(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToUintE(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToUintE(v.Float())
		case reflect.Bool:
			return ToUintE(v.Bool())
		case reflect.String:
			return ToUintE(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}

// ToUint64E convert interface to uint64
func ToUint64E(i interface{}) (uint64, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToUint64E(s[1:])
		}
		v, err := strconv.ParseUint(s, 0, 64)
		if err == nil {
			return v, nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v to uint64", i)
	case int:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint64(s), nil
	case int64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint64(s), nil
	case int32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint64(s), nil
	case int16:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint64(s), nil
	case int8:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint64(s), nil
	case uint:
		return uint64(s), nil
	case uint64:
		return s, nil
	case uint32:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case float32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint64(s), nil
	case float64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint64(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToUint64E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToUint64E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToUint64E(v.Float())
		case reflect.Bool:
			return ToUint64E(v.Bool())
		case reflect.String:
			return ToUint64E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}

// ToUint32E convert interface to uint32
func ToUint32E(i interface{}) (uint32, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToUint32E(s[1:])
		}
		v, err := strconv.ParseUint(s, 0, 32)
		if err == nil {
			return uint32(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v to uint32", i)
	case int:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint32(s), nil
	case int64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint32(s), nil
	case int32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint32(s), nil
	case int16:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint32(s), nil
	case int8:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint32(s), nil
	case uint:
		return uint32(s), nil
	case uint64:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint16:
		return uint32(s), nil
	case uint8:
		return uint32(s), nil
	case float64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint32(s), nil
	case float32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint32(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToUint32E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToUint32E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToUint32E(v.Float())
		case reflect.Bool:
			return ToUint32E(v.Bool())
		case reflect.String:
			return ToUint32E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}

// ToUint16E convert interface to uint16
func ToUint16E(i interface{}) (uint16, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToUint16E(s[1:])
		}
		v, err := strconv.ParseUint(s, 0, 16)
		if err == nil {
			return uint16(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v to uint16", i)
	case int:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint16(s), nil
	case int64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint16(s), nil
	case int32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint16(s), nil
	case int16:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint16(s), nil
	case int8:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint16(s), nil
	case uint:
		return uint16(s), nil
	case uint64:
		return uint16(s), nil
	case uint32:
		return uint16(s), nil
	case uint16:
		return s, nil
	case uint8:
		return uint16(s), nil
	case float64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint16(s), nil
	case float32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint16(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToUint16E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToUint16E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToUint16E(v.Float())
		case reflect.Bool:
			return ToUint16E(v.Bool())
		case reflect.String:
			return ToUint16E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
}

// ToUint8E convert interface to uint8
func ToUint8E(i interface{}) (uint8, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case string:
		if strings.HasPrefix(s, "0") && !strings.HasPrefix(s, "0o") {
			return ToUint8E(s[1:])
		}
		v, err := strconv.ParseUint(s, 0, 8)
		if err == nil {
			return uint8(v), nil
		}
		return 0, errors.Wrapf(err, "unable to cast %#v to uint8", i)
	case int:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint8(s), nil
	case int64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint8(s), nil
	case int32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint8(s), nil
	case int16:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint8(s), nil
	case int8:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint8(s), nil
	case uint:
		return uint8(s), nil
	case uint64:
		return uint8(s), nil
	case uint32:
		return uint8(s), nil
	case uint16:
		return uint8(s), nil
	case uint8:
		return s, nil
	case float64:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint8(s), nil
	case float32:
		if s < 0 {
			return 0, errors.WithStack(errNegativeNotAllowed)
		}
		return uint8(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToUint8E(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToUint8E(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToUint8E(v.Float())
		case reflect.Bool:
			return ToUint8E(v.Bool())
		case reflect.String:
			return ToUint8E(v.String())
		}
		return 0, errors.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}

// ToTimeE convert interface to time.Time
func ToTimeE(i interface{}) (tim time.Time, err error) {
	i = Indirect(i)

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		return StringToDateE(v)
	case int:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	default:
		vv := reflect.ValueOf(i)
		switch vv.Kind() {
		case reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToTimeE(vv.Uint())
		case reflect.Int, reflect.Int32, reflect.Int64:
			return ToTimeE(vv.Int())
		case reflect.Float64, reflect.Float32:
			return ToTimeE(vv.Float())
		case reflect.String:
			return ToTimeE(vv.String())
		}
		return time.Time{}, errors.Errorf("unable to cast %#v of type %T to Time", i, i)
	}
}

// ToDurationE convert interface to time.Duration
func ToDurationE(i interface{}) (d time.Duration, err error) {
	i = Indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		d = time.Duration(ToInt64(s))
		return
	case float32, float64:
		d = time.Duration(ToFloat64(s))
		return
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			d, err = time.ParseDuration(s)
		} else {
			d, err = time.ParseDuration(s + "ns")
		}
		err = errors.WithStack(err)
		return
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToDurationE(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToDurationE(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToDurationE(v.Float())
		case reflect.String:
			return ToDurationE(v.String())
		}
		err = errors.Errorf("unable to cast %#v of type %T to Duration", i, i)
		return
	}
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func indirectToStringerOrError(a interface{}) interface{} {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ToStringE convert interface to string
func ToStringE(i interface{}) (string, error) {
	i = indirectToStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatInt(int64(s), 10), nil
	case uint64:
		return strconv.FormatInt(int64(s), 10), nil
	case uint32:
		return strconv.FormatInt(int64(s), 10), nil
	case uint16:
		return strconv.FormatInt(int64(s), 10), nil
	case uint8:
		return strconv.FormatInt(int64(s), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		v := reflect.ValueOf(i)
		if method, ok := reflect.TypeOf(i).MethodByName("String"); ok && method.Type.NumIn() == 0 &&
			method.Type.NumOut() == 1 && method.Type.Out(0).Kind() == reflect.String {
			return method.Func.Call([]reflect.Value{v})[0].String(), nil
		}
		switch v.Kind() {
		case reflect.Func:
			fullName := runtime.FuncForPC(v.Pointer()).Name()
			ss := strings.Split(fullName, ".")
			if len(ss) > 0 {
				return ss[len(ss)-1], nil
			} else {
				return fullName, nil
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return ToStringE(v.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return ToStringE(v.Int())
		case reflect.Float64, reflect.Float32:
			return ToStringE(v.Float())
		case reflect.Bool:
			return ToStringE(v.Bool())
		case reflect.String:
			return v.String(), nil
		}
		return "", errors.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}

// ToStringMapStringE convert interface to map[string]string
func ToStringMapStringE(i interface{}) (map[string]string, error) {
	i = Indirect(i)
	var m = map[string]string{}

	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		mapKeys := value.MapKeys()
		for _, key := range mapKeys {
			sKey, err := ToStringE(key.Interface())
			if err != nil {
				return m, err
			}
			iVal, err := ToStringE(value.MapIndex(key).Interface())
			if err != nil {
				return m, err
			}
			m[sKey] = iVal
		}
		return m, nil
	}

	switch v := i.(type) {
	case string:
		return m, errors.WithStack(jsonStringToObject(v, &m))
	default:
		return m, errors.Errorf("unable to cast %#v of type %T to map[string]string", i, i)
	}
}

// ToStringMapStringSliceE convert interface to map[string][]string
func ToStringMapStringSliceE(i interface{}) (map[string][]string, error) {
	i = Indirect(i)
	var m = map[string][]string{}

	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		mapKeys := value.MapKeys()
		for _, key := range mapKeys {
			sKey, err := ToStringE(key.Interface())
			if err != nil {
				return m, err
			}
			iVal, err := ToStringSliceE(value.MapIndex(key).Interface())
			if err != nil {
				return m, err
			}
			m[sKey] = iVal
		}
		return m, nil
	}

	switch v := i.(type) {
	case string:
		return m, errors.WithStack(jsonStringToObject(v, &m))
	default:
		return m, errors.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
	}
}

// ToStringMapBoolE convert interface to map[string]bool
func ToStringMapBoolE(i interface{}) (map[string]bool, error) {
	i = Indirect(i)
	var m = map[string]bool{}

	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		mapKeys := value.MapKeys()
		for _, key := range mapKeys {
			sKey, err := ToStringE(key.Interface())
			if err != nil {
				return m, err
			}
			iVal, err := ToBoolE(value.MapIndex(key).Interface())
			if err != nil {
				return m, err
			}
			m[sKey] = iVal
		}
		return m, nil
	}

	switch v := i.(type) {
	case string:
		return m, errors.WithStack(jsonStringToObject(v, &m))
	default:
		return m, errors.Errorf("unable to cast %#v of type %T to map[string]bool", i, i)
	}
}

// ToStringMapE convert interface to map[string]interface{}
func ToStringMapE(i interface{}) (map[string]interface{}, error) {
	i = Indirect(i)

	var m = map[string]interface{}{}

	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		mapKeys := value.MapKeys()
		for _, key := range mapKeys {
			sKey, err := ToStringE(key.Interface())
			if err != nil {
				return m, err
			}
			m[sKey] = value.MapIndex(key).Interface()
		}
		return m, nil
	}

	switch v := i.(type) {
	case string:
		return m, errors.WithStack(jsonStringToObject(v, &m))
	default:
		return m, errors.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	}
}

// ToStringMapIntE convert interface to map[string]int
func ToStringMapIntE(i interface{}) (map[string]int, error) {
	i = Indirect(i)
	var m = map[string]int{}

	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		mapKeys := value.MapKeys()
		for _, key := range mapKeys {
			sKey, err := ToStringE(key.Interface())
			if err != nil {
				return m, err
			}
			iVal, err := ToIntE(value.MapIndex(key).Interface())
			if err != nil {
				return m, err
			}
			m[sKey] = iVal
		}
		return m, nil
	}

	switch v := i.(type) {
	case string:
		return m, errors.WithStack(jsonStringToObject(v, &m))
	default:
		return m, errors.Errorf("unable to cast %#v of type %T to map[string]int", i, i)
	}
}

// ToStringMapInt64E convert interface to map[string]int64
func ToStringMapInt64E(i interface{}) (map[string]int64, error) {
	i = Indirect(i)
	var m = map[string]int64{}

	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		mapKeys := value.MapKeys()
		for _, key := range mapKeys {
			sKey, err := ToStringE(key.Interface())
			if err != nil {
				return m, err
			}
			iVal, err := ToInt64E(value.MapIndex(key).Interface())
			if err != nil {
				return m, err
			}
			m[sKey] = iVal
		}
		return m, nil
	}

	switch v := i.(type) {
	case string:
		return m, errors.WithStack(jsonStringToObject(v, &m))
	default:
		return m, errors.Errorf("unable to cast %#v of type %T to map[string]int64", i, i)
	}
}

// ToSliceE convert interface to []interface{}
func ToSliceE(i interface{}) ([]interface{}, error) {
	i = Indirect(i)
	if i == nil {
		return []interface{}{}, errors.Errorf("unable to cast %#v of type %T to []interface{}", i, i)
	}
	r := reflect.ValueOf(i)
	if r.Kind() == reflect.Slice {
		s := make([]interface{}, r.Len())
		for i := 0; i < r.Len(); i++ {
			s[i] = r.Index(i).Interface()
		}
		return s, nil
	}
	return []interface{}{}, errors.Errorf("unable to cast non Slice to []interface{}")
}

// ToBoolSliceE convert interface to []bool
func ToBoolSliceE(i interface{}) ([]bool, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []bool:
		return v, nil
	case nil:
		return []bool{}, errors.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToBoolE(s.Index(j).Interface())
			if err != nil {
				return []bool{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, errors.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}
}

// ToStringSliceE convert interface to []string
func ToStringSliceE(i interface{}) ([]string, error) {
	i = Indirect(i)

	switch v := i.(type) {
	case []string:
		return v, nil
	case string:
		return strings.Fields(v), nil
	case nil:
		return []string{}, errors.Errorf("unable to cast %#v of type %T to []string", i, i)
	}
	kind := reflect.TypeOf(i).Kind()
	var a []string
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a = make([]string, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToStringE(s.Index(j).Interface())
			if err != nil {
				return []string{}, err
			}
			a[j] = val
		}
		return a, nil
	}
	str, err := ToStringE(i)
	if err != nil {
		return a, err
	}
	return []string{str}, nil
}

// ToIntSliceE convert interface to []int
func ToIntSliceE(i interface{}) ([]int, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []int:
		return v, nil
	case nil:
		return []int{}, errors.Errorf("unable to cast %#v of type %T to []int", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToIntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, errors.Errorf("unable to cast %#v of type %T to []int", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, errors.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
}

// ToDurationSliceE convert interface to []time.Duration
func ToDurationSliceE(i interface{}) ([]time.Duration, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []time.Duration:
		return v, nil
	case nil:
		return []time.Duration{}, errors.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]time.Duration, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToDurationE(s.Index(j).Interface())
			if err != nil {
				return []time.Duration{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []time.Duration{}, errors.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}
}

// ToByteSliceE convert interface to []byte
func ToByteSliceE(i interface{}) ([]byte, error) {
	if bm, ok := i.(encoding.BinaryMarshaler); ok {
		bb, err := bm.MarshalBinary()
		if err != nil {
			return []byte{}, errors.WithStack(err)
		}
		return bb, nil
	}
	i = Indirect(i)
	switch v := i.(type) {
	case nil:
		return []byte{}, nil
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	default:
		bb, err := json.Marshal(i)
		if err != nil {
			return []byte{}, errors.WithStack(err)
		}
		return bb, nil
	}
}

// ToInt64SliceE convert interface to []int64
func ToInt64SliceE(i interface{}) ([]int64, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []int64:
		return v, nil
	case nil:
		return []int64{}, errors.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToInt64E(s.Index(j).Interface())
			if err != nil {
				return []int64{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int64{}, errors.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}
}

// ToInt32SliceE convert interface to []int32
func ToInt32SliceE(i interface{}) ([]int32, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []int32:
		return v, nil
	case nil:
		return []int32{}, errors.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int32, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToInt32E(s.Index(j).Interface())
			if err != nil {
				return []int32{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int32{}, errors.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}
}

// ToInt16SliceE convert interface to []int16
func ToInt16SliceE(i interface{}) ([]int16, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []int16:
		return v, nil
	case nil:
		return []int16{}, errors.Errorf("unable to cast %#v of type %T to []int16", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int16, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToInt16E(s.Index(j).Interface())
			if err != nil {
				return []int16{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int16{}, errors.Errorf("unable to cast %#v of type %T to []int16", i, i)
	}
}

// ToInt8SliceE convert interface to []int8
func ToInt8SliceE(i interface{}) ([]int8, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []int8:
		return v, nil
	case nil:
		return []int8{}, errors.Errorf("unable to cast %#v of type %T to []int8", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int8, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToInt8E(s.Index(j).Interface())
			if err != nil {
				return []int8{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int8{}, errors.Errorf("unable to cast %#v of type %T to []int8", i, i)
	}
}

// ToUint64SliceE convert interface to []uint64
func ToUint64SliceE(i interface{}) ([]uint64, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []uint64:
		return v, nil
	case nil:
		return []uint64{}, errors.Errorf("unable to cast %#v of type %T to []uint64", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUint64E(s.Index(j).Interface())
			if err != nil {
				return []uint64{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint64{}, errors.Errorf("unable to cast %#v of type %T to []uint64", i, i)
	}
}

// ToUint32SliceE convert interface to []uint32
func ToUint32SliceE(i interface{}) ([]uint32, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []uint32:
		return v, nil
	case nil:
		return []uint32{}, errors.Errorf("unable to cast %#v of type %T to []uint32", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint32, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUint32E(s.Index(j).Interface())
			if err != nil {
				return []uint32{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint32{}, errors.Errorf("unable to cast %#v of type %T to []uint32", i, i)
	}
}

// ToUint16SliceE convert interface to []uint16
func ToUint16SliceE(i interface{}) ([]uint16, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []uint16:
		return v, nil
	case nil:
		return []uint16{}, errors.Errorf("unable to cast %#v of type %T to []uint16", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint16, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUint16E(s.Index(j).Interface())
			if err != nil {
				return []uint16{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint16{}, errors.Errorf("unable to cast %#v of type %T to []uint16", i, i)
	}
}

// ToUint8SliceE convert interface to []uint8
func ToUint8SliceE(i interface{}) ([]uint8, error) {
	i = Indirect(i)
	switch v := i.(type) {
	case []uint8:
		return v, nil
	case nil:
		return []uint8{}, errors.Errorf("unable to cast %#v of type %T to []uint8", i, i)
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]uint8, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUint8E(s.Index(j).Interface())
			if err != nil {
				return []uint8{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []uint8{}, errors.Errorf("unable to cast %#v of type %T to []uint8", i, i)
	}
}

// ToStructData Unmarshal bb bytes data to struct i
func BinaryToStructDataE(bb []byte, i interface{}) error {
	if bm, ok := i.(encoding.BinaryUnmarshaler); ok {
		// use struct i's UnmarshalBinary to Unmarshal
		err := bm.UnmarshalBinary(bb)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
	err := json.Unmarshal(bb, i)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// jsonStringToObject attempts to unmarshall a string as JSON into
// the object passed as pointer.
func jsonStringToObject(s string, v interface{}) error {
	return errors.WithStack(json.Unmarshal([]byte(s), v))
}

// ToJSONStringE convert interface to json string
func ToJSONStringE(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	return Bytes2Str(bytes), errors.WithStack(err)
}

// ToStyledJSONStringE  convert interface to styled json string
func ToStyledJSONStringE(v interface{}) (string, error) {
	bytes, err := json.MarshalIndent(v, "", "\t")
	return Bytes2Str(bytes), errors.WithStack(err)
}

// MapToKeySliceE convert map interface's key to interface slice
func MapToKeySliceE(i interface{}) ([]interface{}, error) {
	i = Indirect(i)
	value := reflect.ValueOf(i)
	if value.Kind() != reflect.Map {
		return []interface{}{}, errors.Errorf("unable to cast %#v of type %T to map", i, i)
	}
	mapKeys := value.MapKeys()
	arr := make([]interface{}, len(mapKeys))
	for i, key := range mapKeys {
		arr[i] = key.Interface()
	}
	return arr, nil
}

// MapToValueSliceE convert map interface's value to interface slice
func MapToValueSliceE(i interface{}) ([]interface{}, error) {
	i = Indirect(i)
	value := reflect.ValueOf(i)
	if value.Kind() != reflect.Map {
		return []interface{}{}, errors.Errorf("unable to cast %#v of type %T to map", i, i)
	}
	mapKeys := value.MapKeys()
	arr := make([]interface{}, len(mapKeys))
	for i, key := range mapKeys {
		arr[i] = value.MapIndex(key).Interface()
	}
	return arr, nil
}

// ToMobileE convert interface to phone number string without +86 profix
func ToMobileE(i interface{}) (string, error) {
	str, err := ToStringE(i)
	if err != nil {
		return "", err
	}
	str = strings.TrimSpace(str)
	if IsCnMobile(str) {
		return strings.TrimPrefix(str, "+86"), nil
	}
	return "", errors.Errorf("invalid mobile format:%v", str)
}

func IsCnMobile(str string) bool {
	return cnMobile.MatchString(str)
}
