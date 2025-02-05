package convert

import "reflect"

// ExtractStructToStringSlice 通过rangV函数，提取interface slice的结构体中某个成员，组成string slice
func ExtractStructToStringSlice(x interface{}, rangeV func(i int) string) []string {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return []string{}
	}

	result := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = rangeV(i)
	}
	return result
}

// ExtractStructToInt64Slice 通过rangV函数，提取interface slice的结构体中某个成员，int64 slice
func ExtractStructToInt64Slice(x interface{}, rangeV func(i int) int64) []int64 {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return []int64{}
	}

	result := make([]int64, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = rangeV(i)
	}
	return result
}

// ExtractStructToInt32Slice 通过rangV函数，提取interface slice的结构体中某个成员，int32 slice
func ExtractStructToInt32Slice(x interface{}, rangeV func(i int) int32) []int32 {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return []int32{}
	}

	result := make([]int32, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = rangeV(i)
	}
	return result
}

// ExtractStructToInt8Slice 通过rangV函数，提取interface slice的结构体中某个成员，int8 slice
func ExtractStructToInt8Slice(x interface{}, rangeV func(i int) int8) []int8 {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return []int8{}
	}

	result := make([]int8, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = rangeV(i)
	}
	return result
}

// ExtractStructToUint64Slice 通过rangV函数，提取interface slice的结构体中某个成员，uint64 slice
func ExtractStructToUint64Slice(x interface{}, rangeV func(i int) uint64) []uint64 {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return []uint64{}
	}

	result := make([]uint64, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = rangeV(i)
	}
	return result
}

// ExtractStructToUint32Slice 通过rangV函数，提取interface slice的结构体中某个成员，uint32 slice
func ExtractStructToUint32Slice(x interface{}, rangeV func(i int) uint32) []uint32 {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return []uint32{}
	}

	result := make([]uint32, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = rangeV(i)
	}
	return result
}

// ExtractStructToUint8Slice 通过rangV函数，提取interface slice的结构体中某个成员，uint8 slice
func ExtractStructToUint8Slice(x interface{}, rangeV func(i int) uint8) []uint8 {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		return []uint8{}
	}

	result := make([]uint8, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = rangeV(i)
	}
	return result
}
