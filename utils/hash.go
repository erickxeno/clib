package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spaolacci/murmur3"
	"hash/crc32"
	"reflect"
	"strconv"
)

func HashCode(obj interface{}) uint32 {
	bytes := typeToByte(obj)
	hash := crc32.ChecksumIEEE(bytes)
	return hash
}

func MurmurSum32(obj interface{}) uint32 {
	bytes := typeToByte(obj)
	return murmur3.Sum32WithSeed(bytes, MurmurDefaultSeed)
}

func MurmurSum32WithSeed(obj interface{}, seed uint32) uint32 {
	bytes := typeToByte(obj)
	return murmur3.Sum32WithSeed(bytes, seed)
}

func MurmurSum64(obj interface{}) uint64 {
	bytes := typeToByte(obj)
	return murmur3.Sum64WithSeed(bytes, MurmurDefaultSeed)
}
func MurmurSum64WithSeed(obj interface{}, seed uint32) uint64 {
	bytes := typeToByte(obj)
	return murmur3.Sum64WithSeed(bytes, seed)
}

func typeToByte(obj interface{}) []byte {
	objValue := reflect.ValueOf(obj)
	var bytes []byte
	var err error
	switch objValue.Kind() {
	case reflect.Bool:
		if objValue.Bool() {
			bytes = []byte{1}
		} else {
			bytes = []byte{0}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num := objValue.Int()
		bytes = []byte(strconv.FormatInt(num, 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		num := objValue.Uint()
		bytes = []byte(strconv.FormatUint(num, 10))
	case reflect.Float32, reflect.Float64:
		num := objValue.Float()
		bytes = []byte(strconv.FormatFloat(num, 'f', -1, 64))
	case reflect.Complex64, reflect.Complex128:
		num := objValue.Complex()
		bytes = []byte(fmt.Sprintf("%v", num))
	case reflect.String:
		str := objValue.String()
		bytes = []byte(str)
	default:
		bytes, err = json.Marshal(obj)
		if err != nil {
			panic(err)
		}
	}
	return bytes
}
