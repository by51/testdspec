# Design: 添加毫秒级时间戳字段

## 技术方案
采用直接扩展 `TimestampResponse` 结构体的方案，新增 `UnixMs` 字段。这是最简单直接的实现方式，Go 标准库的 `time.Time` 已提供 `UnixMilli()` 方法（Go 1.17+），可以直接获取毫秒级时间戳。

## 模块结构

### main.go
- **TimestampResponse 结构体**：新增 `UnixMs int64` 字段，JSON tag 为 `unix_ms`
- **timestampHandler 函数**：调用 `time.Now().UnixMilli()` 获取毫秒级时间戳并赋值

```
main.go
├── TimestampResponse (结构体)
│   ├── Unix     int64  `json:"unix"`      // 秒级时间戳（已有）
│   ├── UnixMs   int64  `json:"unix_ms"`   // 毫秒级时间戳（新增）
│   └── Readable string `json:"readable"` // 可读格式（已有）
└── timestampHandler (函数)
    └── 计算并返回三个字段的值
```

## 关键决策

### 为什么使用 `UnixMilli()` 方法
- Go 1.17+ 提供原生支持，无需手动计算
- 性能优于 `UnixNano() / 1e6` 的方式
- 代码更简洁、可读性更好

### 为什么不删除秒级时间戳
- 保持向后兼容，现有客户端依赖 `unix` 字段
- 秒级时间戳在某些场景下仍有使用价值（如日志聚合、简单计时）
- 两个字段共存，让客户端按需选择

### 字段命名选择
- 使用 `unix_ms` 而非 `unixMilli` 或 `milliseconds`
- 与现有 `unix` 字段命名风格一致
- 下划线风格符合 JSON 命名惯例，且与现有 `readable` 字段风格统一
