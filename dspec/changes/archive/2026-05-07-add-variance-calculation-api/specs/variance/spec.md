## ADDED Requirements

### Requirement: 方差计算接口
系统应当提供 HTTP 接口，接受一组浮点数列表，计算并返回其总体方差。

#### Scenario: 正常计算方差
- **WHEN** 客户端发送 `POST /variance`，请求体为 `{"items": [2, 4, 4, 4, 5, 5, 7, 9]}`
- **THEN** 响应状态码 200，响应体包含 `{"variance": 4.0}`

#### Scenario: 单个元素列表
- **WHEN** 请求体 `items` 只含一个数字
- **THEN** 返回方差 `0.0`

#### Scenario: 空列表
- **WHEN** 请求体 `items` 为空数组
- **THEN** 返回状态码 400，包含错误信息 `"items must not be empty"`

#### Scenario: 请求体非法
- **WHEN** 请求体无法解析为合法 JSON，或 `items` 字段缺失
- **THEN** 返回状态码 400，包含错误信息 `"invalid request body"`

#### Scenario: 方法不允许
- **WHEN** 使用 GET 等非 POST 方法访问 `/variance`
- **THEN** 返回状态码 405
