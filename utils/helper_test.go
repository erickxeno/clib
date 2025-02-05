package utils

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

type testStruct struct {
	I32       int32  `json:"i32"`
	I64       int32  `json:"i64"`
	StringVal string `json:"string_val"`
}

func TestMurmur(t *testing.T) {
	var data = []struct {
		seed uint32
		h32  uint32
		h64  uint64
		s    string
	}{
		{0x00, 0x00000000, 0x0000000000000000, ""},
		{0x00, 0x248bfa47, 0xcbd8a7b341bd9b02, "hello"},
		{0x00, 0x149bbb7f, 0x342fac623a5ebc8e, "hello, world"},
		{0x00, 0xe31e8a70, 0xb89e5988b737affc, "19 Jan 2038 at 3:14:07 AM"},
		{0x00, 0xd5c48bfc, 0xcd99481f9ee902c9, "The quick brown fox jumps over the lazy dog."},

		{0x01, 0x514e28b7, 0x4610abe56eff5cb5, ""},
		{0x01, 0xbb4abcad, 0xa78ddff5adae8d10, "hello"},
		{0x01, 0x6f5cb2e9, 0x8b95f808840725c6, "hello, world"},
		{0x01, 0xf50e1f30, 0x2a929de9c8f97b2f, "19 Jan 2038 at 3:14:07 AM"},
		{0x01, 0x846f6a36, 0xfb3325171f9744da, "The quick brown fox jumps over the lazy dog."},

		{0x2a, 0x087fcd5c, 0xf02aa77dfa1b8523, ""},
		{0x2a, 0xe2dbd2e1, 0xc4b8b3c960af6f08, "hello"},
		{0x2a, 0x7ec7c6c2, 0xb91864d797caa956, "hello, world"},
		{0x2a, 0x58f745f6, 0xfd8f19ebdc8c6b6a, "19 Jan 2038 at 3:14:07 AM"},
		{0x2a, 0xc02d1434, 0x74f33c659cda5af7, "The quick brown fox jumps over the lazy dog."},

		{MurmurDefaultSeed, 0xfafd39d1, 0x7ba23256b7b863af, ""},
		{MurmurDefaultSeed, 0xeac790e5, 0x8ea93274ec6bc204, "hello"},
		{MurmurDefaultSeed, 0xe51e255f, 0x4406e2d44aa3e91d, "hello, world"},
		{MurmurDefaultSeed, 0x391ffaf1, 0xc2789fbe1efcf19e, "19 Jan 2038 at 3:14:07 AM"},
		{MurmurDefaultSeed, 0xa5ef7ef0, 0x55aa24df07dddbd, "The quick brown fox jumps over the lazy dog."},
	}
	var data2 = []struct {
		seed uint32
		h32  uint32
		h64  uint64
		s    string
	}{
		{MurmurDefaultSeed, 0xfafd39d1, 0x7ba23256b7b863af, ""},
		{MurmurDefaultSeed, 0xeac790e5, 0x8ea93274ec6bc204, "hello"},
		{MurmurDefaultSeed, 0xe51e255f, 0x4406e2d44aa3e91d, "hello, world"},
		{MurmurDefaultSeed, 0x391ffaf1, 0xc2789fbe1efcf19e, "19 Jan 2038 at 3:14:07 AM"},
		{MurmurDefaultSeed, 0xa5ef7ef0, 0x55aa24df07dddbd, "The quick brown fox jumps over the lazy dog."},
	}

	PatchConvey("test murmur3", t, func() {
		t.Logf("test murmur3")
		for _, ele := range data {
			val32 := MurmurSum32WithSeed(ele.s, ele.seed)
			So(val32, ShouldEqual, ele.h32)

			val64 := MurmurSum64WithSeed(ele.s, ele.seed)
			So(val64, ShouldEqual, ele.h64)
		}
		for _, ele := range data2 {
			val32 := MurmurSum32(ele.s)
			So(val32, ShouldEqual, ele.h32)

			val64 := MurmurSum64(ele.s)
			So(val64, ShouldEqual, ele.h64)
		}
	})

	PatchConvey("test murmur3 with other type", t, func() {
		t.Logf("test murmur3 with other type")
		var v1 int = 999
		var v2 bool = true
		var v3 int8 = 9
		var v4 int64 = 999
		var v5 int32 = 999
		var v6 uint8 = 99
		var v7 byte = 9
		var v8 uint32 = 999
		var v9 uint64 = 999
		var v10 float32 = 999
		var v11 float64 = 999
		var v12 = &testStruct{
			I32:       32,
			I64:       64,
			StringVal: "string_val",
		}
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v1, MurmurSum32(v1), MurmurSum64(v1))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v2, MurmurSum32(v2), MurmurSum64(v2))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v3, MurmurSum32(v3), MurmurSum64(v3))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v4, MurmurSum32(v4), MurmurSum64(v4))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v5, MurmurSum32(v5), MurmurSum64(v5))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v6, MurmurSum32(v6), MurmurSum64(v6))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v7, MurmurSum32(v7), MurmurSum64(v7))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v8, MurmurSum32(v8), MurmurSum64(v8))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v9, MurmurSum32(v9), MurmurSum64(v9))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v10, MurmurSum32(v10), MurmurSum64(v10))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v11, MurmurSum32(v11), MurmurSum64(v11))
		t.Logf("%v, murmur3 32 res:0x%x, 64 res:%x", v12, MurmurSum32(v12), MurmurSum64(v12))
	})
}

type errorImpl struct{}

func (e *errorImpl) Error() string {
	return ""
}

var ei *errorImpl

func errorImplFun() error {
	return ei
}

func toInterface(i interface{}) interface{} {
	return i
}

func TestIsNil(t *testing.T) {
	PatchConvey("test IsNil", t, func() {
		t.Logf("test IsNil")

		var a1 *struct{} = nil
		var a2 []int = nil
		var a3 map[int]bool = nil
		var a4 chan string = nil
		var a5 func() = nil
		var a6 interface{} = nil

		So(IsNil(a1), ShouldBeTrue)
		So(IsNil(a2), ShouldBeTrue)
		So(IsNil(a3), ShouldBeTrue)
		So(IsNil(a4), ShouldBeTrue)
		So(IsNil(a5), ShouldBeTrue)
		So(IsNil(a6), ShouldBeTrue)
		So(a1 == nil, ShouldBeTrue)
		So(a2 == nil, ShouldBeTrue)
		So(a3 == nil, ShouldBeTrue)
		So(a4 == nil, ShouldBeTrue)
		So(a5 == nil, ShouldBeTrue)
		So(a6 == nil, ShouldBeTrue)

		So(IsNil(toInterface(a1)), ShouldBeTrue)
		So(IsNil(toInterface(a2)), ShouldBeTrue)
		So(IsNil(toInterface(a3)), ShouldBeTrue)
		So(IsNil(toInterface(a4)), ShouldBeTrue)
		So(IsNil(toInterface(a5)), ShouldBeTrue)
		So(IsNil(toInterface(a6)), ShouldBeTrue)
		// these are false
		So(toInterface(a1) == nil, ShouldBeFalse)
		So(toInterface(a2) == nil, ShouldBeFalse)
		So(toInterface(a3) == nil, ShouldBeFalse)
		So(toInterface(a4) == nil, ShouldBeFalse)
		So(toInterface(a5) == nil, ShouldBeFalse)
		// this is true
		So(toInterface(a6) == nil, ShouldBeTrue)

		// this is false
		So(errorImplFun() == nil, ShouldBeFalse)
		So(IsNil(errorImplFun()), ShouldBeTrue)

		// string cannot be nil
		var s string
		So(IsNil(s), ShouldBeFalse)
		var i interface{}
		So(IsNil(i), ShouldBeTrue)
	})
}
