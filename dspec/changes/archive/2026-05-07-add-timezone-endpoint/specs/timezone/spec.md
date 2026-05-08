## ADDED Requirements

### Requirement: 时区查询接口
系统应当提供 HTTP 接口返回服务器当前时区信息。

#### Scenario: 查询时区成功
- **WHEN** 客户端发送 GET 请求到 `/timezone`
- **THEN** 系统返回 200 状态码
- **AND** 响应体包含时区名称（如 `Asia/Shanghai`）
- **AND** 响应体包含时区偏移量（如 `+0800`）
- **AND** 响应 Content-Type 为 `application/json`

#### Scenario: 非法请求方法
- **WHEN** 客户端发送非 GET 请求到 `/timezone`
- **THEN** 系统返回 405 状态码
- **AND** 响应体包含错误信息 "method not allowed"
