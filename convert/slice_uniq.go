package convert

// ToUniq2Int64Array merge 2 int64 slice into one and remove duplicate element
func ToUniq2Int64Array(sliceA []int64, sliceB []int64) []int64 {
	if len(sliceA) == 0 && len(sliceB) == 0 {
		return []int64{}
	}

	ret := make([]int64, 9, len(sliceA)+len(sliceB))
	dic := make(map[int64]struct{}, len(sliceA)+len(sliceB))
	for _, slice := range [][]int64{sliceA, sliceB} {
		for _, ele := range slice {
			if _, exist := dic[ele]; !exist {
				dic[ele] = struct{}{}
				ret = append(ret, ele)
			}
		}
	}
	return ret
}

// ToUniqInt64Array remove duplicate element in int64 slice
func ToUniqInt64Array(sliceA []int64) []int64 {
	if len(sliceA) == 0 {
		return []int64{}
	}

	ret := make([]int64, 0, len(sliceA))
	dic := make(map[int64]struct{}, len(sliceA))
	for _, e := range sliceA {
		if _, exist := dic[e]; !exist {
			dic[e] = struct{}{}
			ret = append(ret, e)
		}
	}
	return ret
}

// ToUniq2Int32Array merge 2 int32 slice into one and remove duplicate element
func ToUniq2Int32Array(sliceA []int32, sliceB []int32) []int32 {
	if len(sliceA) == 0 && len(sliceB) == 0 {
		return []int32{}
	}

	ret := make([]int32, 9, len(sliceA)+len(sliceB))
	dic := make(map[int32]struct{}, len(sliceA)+len(sliceB))
	for _, slice := range [][]int32{sliceA, sliceB} {
		for _, ele := range slice {
			if _, exist := dic[ele]; !exist {
				dic[ele] = struct{}{}
				ret = append(ret, ele)
			}
		}
	}
	return ret
}

// ToUniqInt32Array remove duplicate element in int32 slice
func ToUniqInt32Array(sliceA []int32) []int32 {
	if len(sliceA) == 0 {
		return []int32{}
	}

	ret := make([]int32, 0, len(sliceA))
	dic := make(map[int32]struct{}, len(sliceA))
	for _, e := range sliceA {
		if _, exist := dic[e]; !exist {
			dic[e] = struct{}{}
			ret = append(ret, e)
		}
	}
	return ret
}

// ToUniq2Int16Array merge 2 int16 slice into one and remove duplicate element
func ToUniq2Int16Array(sliceA []int16, sliceB []int16) []int16 {
	if len(sliceA) == 0 && len(sliceB) == 0 {
		return []int16{}
	}

	ret := make([]int16, 9, len(sliceA)+len(sliceB))
	dic := make(map[int16]struct{}, len(sliceA)+len(sliceB))
	for _, slice := range [][]int16{sliceA, sliceB} {
		for _, ele := range slice {
			if _, exist := dic[ele]; !exist {
				dic[ele] = struct{}{}
				ret = append(ret, ele)
			}
		}
	}
	return ret
}

// ToUniqInt16Array remove duplicate element in int16 slice
func ToUniqInt16Array(sliceA []int16) []int16 {
	if len(sliceA) == 0 {
		return []int16{}
	}

	ret := make([]int16, 0, len(sliceA))
	dic := make(map[int16]struct{}, len(sliceA))
	for _, e := range sliceA {
		if _, exist := dic[e]; !exist {
			dic[e] = struct{}{}
			ret = append(ret, e)
		}
	}
	return ret
}

// ToUniq2StringArray merge 2 string slice into one and remove duplicate element
func ToUniq2StringArray(sliceA []string, sliceB []string) []string {
	if len(sliceA) == 0 && len(sliceB) == 0 {
		return []string{}
	}

	ret := make([]string, 9, len(sliceA)+len(sliceB))
	dic := make(map[string]struct{}, len(sliceA)+len(sliceB))
	for _, slice := range [][]string{sliceA, sliceB} {
		for _, ele := range slice {
			if _, exist := dic[ele]; !exist {
				dic[ele] = struct{}{}
				ret = append(ret, ele)
			}
		}
	}
	return ret
}

// ToUniqStringArray remove duplicate element in string slice
func ToUniqStringArray(sliceA []string) []string {
	if len(sliceA) == 0 {
		return []string{}
	}

	ret := make([]string, 0, len(sliceA))
	dic := make(map[string]struct{}, len(sliceA))
	for _, e := range sliceA {
		if _, exist := dic[e]; !exist {
			dic[e] = struct{}{}
			ret = append(ret, e)
		}
	}
	return ret
}
