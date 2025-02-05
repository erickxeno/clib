package errors

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestNew(t *testing.T) {
	type args struct {
		code int32
		msg  string
		args []interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantCode int32
		wantMsg  string
	}{
		{
			name: "success",
			args: args{
				code: 1,
				msg:  "111",
			},
			wantCode: 1,
			wantMsg:  "111",
		},
		{
			name: "format",
			args: args{
				code: 1,
				msg:  "1%d3",
				args: []interface{}{2},
			},
			wantCode: 1,
			wantMsg:  "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.code, tt.args.msg, tt.args.args...)
			if !reflect.DeepEqual(GetCode(got), tt.wantCode) || !reflect.DeepEqual(GetMsg(got), tt.wantMsg) {
				t.Errorf("New() got = %v, want %v:%v", got, tt.wantCode, tt.wantMsg)
			}
		})
	}
}

type testError struct {
	code int32
	msg  string
}

func (e testError) Code() int32 {
	return e.code
}

func (e testError) Msg() string {
	return e.msg
}

func (e testError) Error() string {
	return e.msg
}

func TestGetCode(t *testing.T) {
	ee := New(100, "10000")
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "struct",
			args: args{
				err: ee,
			},
			want: 100,
		},
		{
			name: "ptr",
			args: args{
				err: ee,
			},
			want: 100,
		},
		{
			name: "interface",
			args: args{
				err: testError{code: 100},
			},
			want: 100,
		},
		{
			name: "wrap",
			args: args{
				err: errors.WithStack(ee),
			},
			want: 100,
		},
		{
			name: "not match",
			args: args{
				err: errors.New("error"),
			},
			want: UnknownErrorCode,
		},
		{
			name: "nil",
			args: args{
				err: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCode(tt.args.err); got != tt.want {
				t.Errorf("Name = %v, GetCode() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestGetMsg(t *testing.T) {
	ee := New(100, "111")

	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "struct",
			args: args{
				err: ee,
			},
			want: "111",
		},
		{
			name: "ptr",
			args: args{
				err: ee,
			},
			want: "111",
		},
		{
			name: "interface",
			args: args{
				err: testError{msg: "111"},
			},
			want: "111",
		},
		{
			name: "wrap",
			args: args{
				err: errors.WithStack(errors.WithStack(ee)),
			},
			want: "111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMsg(tt.args.err); got != tt.want {
				t.Errorf("GetMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIs(t *testing.T) {
	ee := New(100, "111")

	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "struct",
			args: args{
				err: ee,
			},
			want: true,
		},
		{
			name: "ptr",
			args: args{
				err: ee,
			},
			want: true,
		},
		{
			name: "interface",
			args: args{
				err: testError{},
			},
			want: true,
		},
		{
			name: "wrap",
			args: args{
				err: errors.WithStack(ee),
			},
			want: true,
		},
		{
			name: "not match",
			args: args{
				err: errors.New("error"),
			},
			want: false,
		},
		{
			name: "nil",
			args: args{
				err: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBizError(tt.args.err); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
