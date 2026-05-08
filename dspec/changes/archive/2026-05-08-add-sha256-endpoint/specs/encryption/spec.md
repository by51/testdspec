## ADDED Requirements

### Requirement: SHA256 哈希计算
系统应当提供 SHA256 哈希计算能力，将输入字符串转换为 64 位小写十六进制哈希值。

#### Scenario: 正常计算 SHA256 哈希
- **WHEN** 客户端发送 `POST /sha256` 请求，请求体包含 `{"input": "hello world"}`
- **THEN** 系统返回 `{"hash": "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"}`，状态码为 200

#### Scenario: 空字符串输入
- **WHEN** 客户端发送 `POST /sha256` 请求，请求体包含 `{"input": ""}`
- **THEN** 系统返回 `{"hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}`，状态码为 200

#### Scenario: 缺少 input 字段
- **WHEN** 客户端发送 `POST /sha256` 请求，请求体不包含 `input` 字段
- **THEN** 系统返回 `{"error": "missing input field"}`，状态码为 400

#### Scenario: 无效 JSON 格式
- **WHEN** 客户端发送 `POST /sha256` 请求，请求体不是有效的 JSON
- **THEN** 系统返回 `{"error": "invalid request body"}`，状态码为 400

#### Scenario: 使用非 POST 方法
- **WHEN** 客户端使用 GET 或其他方法请求 `/sha256`
- **THEN** 系统返回状态码 405（Method Not Allowed）