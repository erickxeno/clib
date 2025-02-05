## 简介
> 对time库的简单封装，优化获取时间的系统调用性能消耗 & 序列化开销

> TODO: 后续可进一步拓展逻辑时间概念，添加offset，时间倍速设置，及逻辑时间标记等。以实现模拟加速时间的效果

## 核心函数
* `SetClock(time.Duration)`: 设置内部ticker的间隔，默认为1毫秒
* `Current()`:  获取包含当前时间信息的Time对象
* `Now()`: 获取time.Time对象
* `String() string / ReadOnlyData() []byte`: Time对象缓存的序列化时间信息（不带时区精确到毫秒，格式如：`2023-08-12 23:12:22,481`）
* `StringWithZone() string / ReadOnlyDataWithZone() []byte`: 带时区版本，格式如：`2023-08-12 23:12:22,481+0800`

> 返回[]byte的版本能够避免转换为string的一次内存拷贝