package convert

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/erickxeno/clib/errors"
)

type VersionInfo struct {
	Major int // 主版本号
	Minor int // 子版本号
	Build int // build版本
}

func (v VersionInfo) String() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Build)
}

// ToVersionE 返回版本信息，包含主版本，子版本, build版本, 支持错误返回值
func ToVersionE(i interface{}) (*VersionInfo, error) {
	i = Indirect(i)

	switch s := i.(type) {
	case VersionInfo:
		return &s, nil
	case string:
		s = strings.TrimSpace(s)
		versionStr := regexp.MustCompile(`[^\d]*([\d\.]+)(.*)`).ReplaceAllString(s, "${1}")
		intArray := ToIntSlice(strings.Split(versionStr, "."))
		if len(intArray) >= 3 {
			return &VersionInfo{
				Major: intArray[0],
				Minor: intArray[1],
				Build: intArray[2],
			}, nil

		}
		if len(intArray) >= 2 {
			return &VersionInfo{
				Major: intArray[0],
				Minor: intArray[1],
				Build: 0,
			}, nil
		}
		if len(intArray) > 0 {
			return &VersionInfo{
				Major: intArray[0],
				Minor: 0,
				Build: 0,
			}, nil
		}
		return &VersionInfo{
			Major: 0,
			Minor: 0,
			Build: 0,
		}, nil
	case nil:
		return &VersionInfo{
			Major: 0,
			Minor: 0,
			Build: 0,
		}, errors.Errorf("unable to cast %#v of type %T to VersionInfo", i, i)
	default:
		v := reflect.ValueOf(i)
		if v.Kind() == reflect.String {
			return ToVersionE(v.String())
		}
		return &VersionInfo{
			Major: 0,
			Minor: 0,
			Build: 0,
		}, errors.Errorf("unable to cast %#v of type %T to VersionInfo", i, i)
	}
}
