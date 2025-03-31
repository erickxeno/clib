package utils

// 线程安全的随机生成器。结合 sync.Pool 使用，避免频繁创建和销毁随机数生成器

import (
	"math/rand"
	"sync"
	"time"
)

// randPool 随机数生成器池
var randPool = sync.Pool{
	New: func() interface{} {
		src := rand.NewSource(time.Now().UnixNano())
		return rand.New(src)
	},
}

// 均匀分布随机数
// 生成[0.0, 1.0)范围内的均匀分布随机数
// 每个数出现的概率相等
// 常用于：
// 基础随机数生成
// 概率模拟
// 随机选择
func SafeRandFloat64() float64 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Float64()
}

// 正态分布随机数
// 生成服从标准正态分布(μ=0, σ=1)的随机数
// 范围：(-∞, +∞)
// 概率密度函数：f(x) = (1/√(2π))e^(-x²/2)
// 常用于：
// 模拟测量误差
// 模拟自然现象
// 统计分析
func SafeRandNormFloat64() float64 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.NormFloat64()
}

// 指数分布随机数
// 生成服从指数分布的随机数
// 范围：(0, +math.MaxFloat64]
// 默认参数λ=1，期望值E=1/λ=1
// 概率密度函数：f(x) = λe^(-λx)
// 常用于模拟：
// 事件间隔时间
// 服务等待时间
// 设备寿命
func SafeRandExpFloat64() float64 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.ExpFloat64()
}

func SafeRandInt63() int64 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Int63()
}

func SafeRandUint32() uint32 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Uint32()
}

func SafeRandUint64() uint64 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Uint64()
}

func SafeRandInt31() int32 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Int31()
}

func SafeRandInt() int {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Int()
}

func SafeRandInt63n(n int64) int64 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Int63n(n)
}

func SafeRandInt31n(n int32) int32 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Int31n(n)
}

func SafeRandIntn(n int) int {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Intn(n)
}

func SafeRandPerm(n int) []int {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Perm(n)
}

func SafeRandFloat32() float32 {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	return r.Float32()
}

func SafeRandShuffle(n int, swap func(i, j int)) {
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)
	r.Shuffle(n, swap)
}
