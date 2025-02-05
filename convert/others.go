package convert

import (
	"encoding/json"

	"github.com/erickxeno/clib/errors"
)

// DeepCopy
func DeepCopy(dst, src interface{}) error {
	if src == nil || dst == nil {
		return errors.Errorf("src or dst is nil")
	}
	byt, err := json.Marshal(src)
	if err != nil {
		return errors.WithStack(err)
	}
	return json.Unmarshal(byt, dst)
}

// ToMobile convert interface to phone number string without +86 profix, if err return empty string
func ToMobile(i interface{}) string {
	m, _ := ToMobileE(i)
	return m
}

// ToVersion convert i to VersionInfo struct, with Major, Minor, and Build info
func ToVersion(i interface{}) *VersionInfo {
	v, _ := ToVersionE(i)
	return v
}

// JSONEscape escape json string
func JSONEscape(str string) string {
	bytes, _ := json.Marshal(str)
	if bytes == nil {
		return ""
	}
	eStr := string(bytes)
	return eStr[1 : len(eStr)-1]
}
