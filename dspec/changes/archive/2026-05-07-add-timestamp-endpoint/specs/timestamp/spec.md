## ADDED Requirements

### Requirement: 获取当前时间戳
系统应当提供 HTTP 接口返回服务器当前时间戳信息。

#### Scenario: 正常请求时间戳
- **WHEN** 客户端请求 GET `/timestamp`
- **THEN** 系统返回 200 状态码，JSON 格式包含 Unix 时间戳和可读时间字符串

#### Scenario: 返回格式正确
- **WHEN** 客户端请求时间戳
- **THEN** 响应 JSON 包含字段 `unix`（整数，秒级时间戳）和 `readable`（字符串，RFC3339 格式）
