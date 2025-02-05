package convert

import (
	"reflect"
	"testing"
)

type A struct {
	A string
	B string
}

func TestNewTuple(t *testing.T) {
	want := "{\"A\":\"aaa\",\"B\":\"bbb\"}"
	t.Run("new", func(t *testing.T) {
		if got := NewTuple(ToJSONStringE(&A{
			A: "aaa",
			B: "bbb",
		})).GetString(0); !reflect.DeepEqual(got, want) {
			t.Errorf("NewTuple() = %v, want %v", got, want)
		}
	})
	t.Run("len", func(t *testing.T) {
		if got := NewTuple(ToJSONStringE(&A{
			A: "aaa",
			B: "bbb",
		})).Len(); !reflect.DeepEqual(got, 2) {
			t.Errorf("Len() = %v, want %v", got, 2)
		}
	})

}
