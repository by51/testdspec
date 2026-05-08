## ADDED Requirements

### Requirement: SHA1 哈希计算
系统应当提供 SHA1 哈希计算能力，将输入字符串转换为 40 位小写十六进制哈希值。

#### Scenario: 正常计算 SHA1 哈希
- **WHEN** 客户端发送 `POST /sha1` 请求，请求体包含 `{"input": "hello world"}`
- **THEN** 系统返回 `{"hash": "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"}`，状态码为 200

#### Scenario: 空字符串输入
- **WHEN** 客户端发送 `POST /sha1` 请求，请求体包含 `{"input": ""}`
- **THEN** 系统返回 `{"hash": "da39a3ee5e6b4b0d3255bfef95601890afd80709"}`，状态码为 200

#### Scenario: 缺少 input 字段
- **WHEN** 客户端发送 `POST /sha1` 请求，请求体不包含 `input` 字段
- **THEN** 系统返回空字符串的 SHA1 哈希值，状态码为 200

#### Scenario: 无效 JSON 格式
- **WHEN** 客户端发送 `POST /sha1` 请求，请求体不是有效的 JSON
- **THEN** 系统返回 `{"error": "invalid request body"}`，状态码为 400

#### Scenario: 使用非 POST 方法
- **WHEN** 客户端使用 GET 或其他方法请求 `/sha1`
- **THEN** 系统返回状态码 405（Method Not Allowed）