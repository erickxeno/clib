package time

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	osTime "time"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTime(t *testing.T) {
	PatchConvey("test time", t, func() {
		{
			ti := Current()
			t.Logf("%s", ti.String())
			t.Logf("%s", ti.StringWithZone())
			t.Logf("%s", ti.ReadOnlyDataWithoutZone())
			t.Logf("%s", ti.ReadOnlyDataWithZone())
		}
		{
			ti := Now()
			t.Logf("%s", ti)
		}

		t.Logf("test time")
		now := Current()
		osTime.Sleep(10 * osTime.Millisecond)
		t.Logf("%v", Current())
		So(Current().Sub(now.Time) > osTime.Millisecond, ShouldBeTrue)
		osTime.Sleep(10 * osTime.Millisecond)
		t.Logf("%v", Current())
		So(Current().Sub(now.Time) > osTime.Millisecond, ShouldBeTrue)
		osTime.Sleep(10 * osTime.Millisecond)
		t.Logf("%v", Current())
		So(Current().Sub(now.Time) > osTime.Millisecond, ShouldBeTrue)

		t.Logf("sleep 10 micro second")
		osTime.Sleep(10 * osTime.Microsecond)
		t.Logf("%v", Current())
		osTime.Sleep(10 * osTime.Microsecond)
		t.Logf("%v", Current())
		osTime.Sleep(10 * osTime.Microsecond)
		t.Logf("%v", Current())

		t.Logf("sleep mill second")
		osTime.Sleep(10 * osTime.Millisecond)
		t.Logf("%v", Current())

		osTime.Sleep(10 * osTime.Microsecond)
		t.Logf("%v", Current())
		osTime.Sleep(10 * osTime.Microsecond)
		t.Logf("%v", Current())
		osTime.Sleep(10 * osTime.Microsecond)
		t.Logf("%v", Current())

		t.Logf("set clock to 10 millisecond")
		SetClock(osTime.Millisecond * 10)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())
		osTime.Sleep(osTime.Millisecond)
		t.Logf("%v", Current())

	})
}

func TestTimer2(t *testing.T) {
	now := Current()
	osTime.Sleep(10 * osTime.Millisecond)
	fmt.Println(Current())
	assert.True(t, Current().Sub(now.Time) > osTime.Millisecond)
	osTime.Sleep(10 * osTime.Millisecond)
	fmt.Println(Current())
	assert.True(t, Current().Sub(now.Time) > osTime.Millisecond)
	osTime.Sleep(10 * osTime.Millisecond)
	fmt.Println(Current())
	assert.True(t, Current().Sub(now.Time) > osTime.Millisecond)
}

func BenchmarkTimer(b *testing.B) {
	b.ReportAllocs()
	b.Run("ticker", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Current()
			}
		})
	})
	b.Run("std", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				osTime.Now()
			}
		})
	})
	b.Run("ticker2", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			SetClock(osTime.Millisecond * 10)
			for pb.Next() {
				Current()
			}
		})
	})
}
