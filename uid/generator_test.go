package uid

import (
	"context"
	"fmt"
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	ctx = context.Background()
)

func TestUID(t *testing.T) {
	PatchConvey("test UID", t, func() {
		id, err := GetDefaultGenerator().NewID(ctx)
		So(err, ShouldBeNil)
		fmt.Println("id", id)
		ids := GetDefaultGenerator().NewIDs(ctx, 10)
		//So(len(ids), ShouldEqual, 10)
		fmt.Println("ids", ids)
	})
}
