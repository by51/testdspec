# Change: 添加毫秒级时间戳字段

## 为什么（Why）
当前 `/timestamp` 接口返回的 `unix` 字段是秒级时间戳（int64），在某些需要更高精度的场景（如性能监控、事件排序、分布式系统时钟同步）中不够用，需要毫秒级精度。

## 变更内容（What Changes）
- 在 `TimestampResponse` 结构体中新增 `UnixMs` 字段，类型为 `int64`，表示毫秒级 Unix 时间戳
- 修改 `timestampHandler` 函数，计算并返回毫秒级时间戳值
- 接口响应将包含三个字段：`unix`（秒级）、`unix_ms`（毫秒级）、`readable`（可读格式）

## 影响范围（Impact）
- **模块**：`main.go` 中的 `TimestampResponse` 结构体和 `timestampHandler` 函数
- **API**：`GET /timestamp` 接口响应结构变更（新增字段，向后兼容）
- **兼容性**：向后兼容，现有客户端可忽略新增字段
- **风险**：低风险，仅新增字段，不影响现有字段
