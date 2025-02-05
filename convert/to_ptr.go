package convert

import "time"

func ToBoolPtr(b bool) *bool {
	return &b
}

func ToStringPtr(s string) *string {
	return &s
}

func ToIntPtr(i int) *int {
	return &i
}

func ToInt8Ptr(i int8) *int8 {
	return &i
}

func ToInt16Ptr(i int16) *int16 {
	return &i
}

func ToInt32Ptr(i int32) *int32 {
	return &i
}

func ToInt64Ptr(i int64) *int64 {
	return &i
}

func ToUintPtr(i uint) *uint {
	return &i
}

func ToUint8Ptr(i uint8) *uint8 {
	return &i
}

func ToUint16Ptr(i uint16) *uint16 {
	return &i
}

func ToUint32Ptr(i uint32) *uint32 {
	return &i
}

func ToUint64Ptr(i uint64) *uint64 {
	return &i
}

func ToFloat32Ptr(f float32) *float32 {
	return &f
}

func ToFloat64Ptr(f float64) *float64 {
	return &f
}

func ToTimePtr(t time.Time) *time.Time {
	return &t
}
