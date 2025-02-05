package convert

import (
	"time"
	"unsafe"
)

// ToBool convert interface to bool
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToFloat64 convert interface to float64
func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToFloat32 convert interface to float32
func ToFloat32(i interface{}) float32 {
	v, _ := ToFloat32E(i)
	return v
}

// ToInt64 convert interface to int64
func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

// ToInt32 convert interface to int32
func ToInt32(i interface{}) int32 {
	v, _ := ToInt32E(i)
	return v
}

// ToInt16 convert interface to int16
func ToInt16(i interface{}) int16 {
	v, _ := ToInt16E(i)
	return v
}

// ToInt8 convert interface to int8
func ToInt8(i interface{}) int8 {
	v, _ := ToInt8E(i)
	return v
}

// ToInt casts an interface to an int type.
func ToInt(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

// ToUint convert interface to uint
func ToUint(i interface{}) uint {
	v, _ := ToUintE(i)
	return v
}

// ToUint64 convert interface to uint64
func ToUint64(i interface{}) uint64 {
	v, _ := ToUint64E(i)
	return v
}

// ToUint32 convert interface to uint32
func ToUint32(i interface{}) uint32 {
	v, _ := ToUint32E(i)
	return v
}

// ToUint16 convert interface to uint16
func ToUint16(i interface{}) uint16 {
	v, _ := ToUint16E(i)
	return v
}

// ToUint8 convert interface to uint8
func ToUint8(i interface{}) uint8 {
	v, _ := ToUint8E(i)
	return v
}

// ToString convert interface to string
func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}

// ToStringMapString convert interface to map[string]string
func ToStringMapString(i interface{}) map[string]string {
	v, _ := ToStringMapStringE(i)
	return v
}

// ToStringMapStringSlice convert interface to map[string][]string
func ToStringMapStringSlice(i interface{}) map[string][]string {
	v, _ := ToStringMapStringSliceE(i)
	return v
}

// ToStringMapBool convert interface to map[string]boo
func ToStringMapBool(i interface{}) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

// ToStringMapInt convert interface to map[string]int
func ToStringMapInt(i interface{}) map[string]int {
	v, _ := ToStringMapIntE(i)
	return v
}

// ToStringMapInt64 convert interface to map[string]int64
func ToStringMapInt64(i interface{}) map[string]int64 {
	v, _ := ToStringMapInt64E(i)
	return v
}

// ToStringMap convert interface to map[string]interface{}
func ToStringMap(i interface{}) map[string]interface{} {
	v, _ := ToStringMapE(i)
	return v
}

// ToSlice convert interface to []interface{}
func ToSlice(i interface{}) []interface{} {
	v, _ := ToSliceE(i)
	return v
}

// ToBoolSlice convert interface to []bool
func ToBoolSlice(i interface{}) []bool {
	v, _ := ToBoolSliceE(i)
	return v
}

// ToStringSlice convert interface to []string
func ToStringSlice(i interface{}) []string {
	v, _ := ToStringSliceE(i)
	return v
}

// ToIntSlice convert interface to []int
func ToIntSlice(i interface{}) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToDurationSlice convert interface to []time.Duration
func ToDurationSlice(i interface{}) []time.Duration {
	v, _ := ToDurationSliceE(i)
	return v
}

// ToByteSlice convert interface to []byte
func ToByteSlice(i interface{}) []byte {
	v, _ := ToByteSliceE(i)
	return v
}

// BinaryToStructData convert interface to 结构化
func BinaryToStructData(bb []byte, i interface{}) {
	_ = BinaryToStructDataE(bb, i)
}

// ToJSONString convert interface to Json String
func ToJSONString(i interface{}) string {
	s, _ := ToJSONStringE(i)
	return s
}

// ToStyledJSONString convert interface to styled Json String
func ToStyledJSONString(i interface{}) string {
	s, _ := ToStyledJSONStringE(i)
	return s
}

// ToUint64Slice convert interface to []uint64
func ToUint64Slice(i interface{}) []uint64 {
	v, _ := ToUint64SliceE(i)
	return v
}

// ToUint32Slice convert interface to []uint32
func ToUint32Slice(i interface{}) []uint32 {
	v, _ := ToUint32SliceE(i)
	return v
}

// ToUint16Slice convert interface to []uint16
func ToUint16Slice(i interface{}) []uint16 {
	v, _ := ToUint16SliceE(i)
	return v
}

// ToUint8Slice convert interface to []uint8
func ToUint8Slice(i interface{}) []uint8 {
	v, _ := ToUint8SliceE(i)
	return v
}

// ToInt64Slice convert interface to []int64
func ToInt64Slice(i interface{}) []int64 {
	v, _ := ToInt64SliceE(i)
	return v
}

// ToInt32Slice convert interface to []int32
func ToInt32Slice(i interface{}) []int32 {
	v, _ := ToInt32SliceE(i)
	return v
}

// ToInt16Slice convert interface to []int16
func ToInt16Slice(i interface{}) []int16 {
	v, _ := ToInt16SliceE(i)
	return v
}

// ToInt8Slice convert interface to []int8
func ToInt8Slice(i interface{}) []int8 {
	v, _ := ToInt8SliceE(i)
	return v
}

// Str2Bytes convert string to []byte. NOTICE: the []byte return can only read!
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2Str convert []byte to string
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
