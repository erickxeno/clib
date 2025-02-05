package utils

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
)

const (
	SymbolDot                  = "."
	SymbolComma                = ","
	SymbolAsterisk             = "*"
	SymbolSpace                = " "
	SymbolForwardSlash         = "/"
	SymbolBackwardSlash        = "\\"
	SymbolPlus                 = "+"
	SymbolMinus                = "-"
	SymbolEqual                = "="
	MurmurDefaultSeed   uint32 = 0xe31e8a70
)

func IsLocal() bool {
	return runtime.GOOS == "darwin" || runtime.GOOS == "windows"
}

func Marshal(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func IsDir(path string) bool {
	if f, err := os.Stat(path); err == nil {
		return f.Mode().IsDir()
	}
	return false
}

func FileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFile 判断所给的路径是否为文件
func IsFile(path string) bool {
	return FileExist(path) && !IsDir(path)
}

func Md5Hash(in string) string {
	h := md5.New()
	_, _ = io.WriteString(h, in)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha512(in string) string {
	s := sha512.New()
	_, _ = io.WriteString(s, in)
	return fmt.Sprintf("%x", s.Sum(nil))
}

func IsNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)
	if isNilableKind && value.IsNil() {
		return true
	}
	return false
}

// containsKind checks if a specified kind in the slice of kinds.
func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}
	return false
}

func GetFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "NilErrFunc"
	}
	return runtime.FuncForPC(pc).Name()
}
