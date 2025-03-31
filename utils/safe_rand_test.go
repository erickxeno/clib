package utils

import (
	"sync"
	"testing"

	//. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
)

func TestSafeRandFloat64(t *testing.T) {
	ast := require.New(t)
	//testSrc := rand.NewSource(time.Now().UnixNano())
	//testRand := rand.New(testSrc)

	// 测试随机数生成器在100次请求中是否有重复
	hitRatios := sync.Map{}
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			hitRatio := SafeRandFloat64()
			//hitRatio := testRand.Float64()
			// 检查是否有重复的随机数
			if _, ok := hitRatios.Load(hitRatio); ok {
				t.Errorf("发现重复的随机数: %f", hitRatio)
			}
			hitRatios.Store(hitRatio, true)
		}()
	}
	wg.Wait()
	// 验证生成了100个不同的随机数
	count := 0
	hitRatios.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	//So(count, ShouldEqual, 100, "应该生成100个不同的随机数")
	ast.Equal(100, count, "应该生成100个不同的随机数")
}
