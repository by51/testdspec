# Delta Spec: encryption

## ADDED Requirements

### Requirement: MD5 哈希计算
系统应当提供 MD5 哈希计算能力，将输入字符串转换为 32 位小写十六进制哈希值。

#### Scenario: 正常计算 MD5 哈希
- **WHEN** 客户端发送 `POST /md5` 请求，请求体包含 `{"input": "hello world"}`
- **THEN** 系统返回 `{"hash": "5eb63bbbe01eeed093cb22bb8f5acdc3"}`，状态码为 200

#### Scenario: 空字符串输入
- **WHEN** 客户端发送 `POST /md5` 请求，请求体包含 `{"input": ""}`
- **THEN** 系统返回 `{"hash": "d41d8cd98f00b204e9800998ecf8427e"}`，状态码为 200

#### Scenario: 缺少 input 字段
- **WHEN** 客户端发送 `POST /md5` 请求，请求体不包含 `input` 字段
- **THEN** 系统返回 `{"error": "missing input field"}`，状态码为 400

#### Scenario: 无效 JSON 格式
- **WHEN** 客户端发送 `POST /md5` 请求，请求体不是有效的 JSON
- **THEN** 系统返回 `{"error": "invalid request body"}`，状态码为 400

#### Scenario: 使用非 POST 方法
- **WHEN** 客户端使用 GET 或其他方法请求 `/md5`
- **THEN** 系统返回状态码 405（Method Not Allowed）