## ADDED Requirements

### Requirement: 毫秒级时间戳字段
系统应当在时间戳接口响应中提供毫秒级精度的时间戳字段。

#### Scenario: 获取毫秒级时间戳
- **WHEN** 客户端请求 `GET /timestamp` 接口
- **THEN** 响应中包含 `unix_ms` 字段，值为当前时间的毫秒级 Unix 时间戳（从 1970-01-01 00:00:00 UTC 至今的毫秒数）

#### Scenario: 毫秒级时间戳格式
- **WHEN** 系统生成 `unix_ms` 字段
- **THEN** 字段值为 `int64` 类型，精度为毫秒（值为秒级时间戳乘以 1000 加上毫秒部分）

## MODIFIED Requirements

### Requirement: 时间戳接口响应结构
#### Scenario: 完整响应结构
- **WHEN** 客户端请求 `GET /timestamp` 接口
- **THEN** 响应 JSON 包含以下字段：
  - `unix`：秒级 Unix 时间戳（int64）
  - `unix_ms`：毫秒级 Unix 时间戳（int64）
  - `readable`：可读格式时间字符串（RFC3339 格式）
