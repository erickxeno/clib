package uid

import (
	"context"
	"math"

	"github.com/erickxeno/clib/logid"
)

type MockIDGenerator struct{}

func NewMockIDGenerator() Generator {
	return &MockIDGenerator{}
}

func (m *MockIDGenerator) NewID(ctx context.Context) (int64, error) {
	return int64(logid.GetID() & math.MaxInt64), nil
}

func (m *MockIDGenerator) NewIDs(ctx context.Context, count int) []int64 {
	ret := make([]int64, count)
	for i := 0; i < count; i++ {
		ret[i] = int64(logid.GetID() & math.MaxInt64)
	}
	return ret
}
