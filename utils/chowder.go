package utils

import (
	"fmt"
	"strconv"
)

func IsString(v interface{}) bool {
	switch v.(type) {
	case string, fmt.Stringer:
		return true
	default:
		return false
	}
}

func IsStrNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsStrNumDot(s string) bool {
	if len(s) == 0 {
		return false
	}
	dotFound := false
	for _, v := range s {
		if v == '.' {
			if dotFound {
				return false
			}
			dotFound = true
		} else if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
