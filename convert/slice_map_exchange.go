package convert

// Int32SliceToMap convert []int32 slice to map[int32]struct{}
func Int32SliceToMap(s []int32) map[int32]struct{} {
	m := make(map[int32]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// Int64SliceToMap convert []int64 slice to map[int64]struct{}
func Int64SliceToMap(s []int64) map[int64]struct{} {
	m := make(map[int64]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// StringSliceToMap convert []string slice to map[string]struct{}
func StringSliceToMap(s []string) map[string]struct{} {
	m := make(map[string]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToKeySlice convert map interface's key to interface slice
func MapToKeySlice(i interface{}) []interface{} {
	s, _ := MapToKeySliceE(i)
	return s
}

// MapToValueSlice convert map interface's value to interface slice
func MapToValueSlice(i interface{}) []interface{} {
	s, _ := MapToValueSliceE(i)
	return s
}
