# Delta Spec: API

## ADDED Requirements

### Requirement: URL 解码接口
系统应当提供 HTTP 接口，对 URL 编码的字符串进行解码。

#### Scenario: 正常解码 URL 编码字符串
- **WHEN** 客户端发送 `POST /urldecode`，请求体为 `{"input": "hello%20world"}`
- **THEN** 系统返回状态码 200，响应体为 `{"decoded": "hello world"}`

#### Scenario: 解码中文字符
- **WHEN** 客户端发送 `POST /urldecode`，请求体为 `{"input": "%E4%BD%A0%E5%A5%BD"}`
- **THEN** 系统返回状态码 200，响应体为 `{"decoded": "你好"}`

#### Scenario: 解码特殊字符
- **WHEN** 客户端发送 `POST /urldecode`，请求体为 `{"input": "a%3Db%26c%3Dd"}`
- **THEN** 系统返回状态码 200，响应体为 `{"decoded": "a=b&c=d"}`

#### Scenario: 解码未编码字符串
- **WHEN** 客户端发送 `POST /urldecode`，请求体为 `{"input": "hello world"}`
- **THEN** 系统返回状态码 200，响应体为 `{"decoded": "hello world"}`

#### Scenario: 解码空字符串
- **WHEN** 客户端发送 `POST /urldecode`，请求体为 `{"input": ""}`
- **THEN** 系统返回状态码 200，响应体为 `{"decoded": ""}`

#### Scenario: 无效的 URL 编码格式
- **WHEN** 客户端发送 `POST /urldecode`，请求体包含无效编码如 `{"input": "%ZZ"}`
- **THEN** 系统返回状态码 400，响应体为 `{"error": "invalid URL encoding"}`

#### Scenario: 百分号后缺少两位十六进制字符
- **WHEN** 客户端发送 `POST /urldecode`，请求体为 `{"input": "hello%2"}`
- **THEN** 系统返回状态码 400，响应体为 `{"error": "invalid URL encoding"}`

#### Scenario: 缺少 input 字段
- **WHEN** 客户端发送 `POST /urldecode`，请求体不包含 `input` 字段
- **THEN** 系统返回状态码 200，响应体为 `{"decoded": ""}`（空字符串作为默认值）

#### Scenario: 无效 JSON 格式
- **WHEN** 客户端发送 `POST /urldecode`，请求体不是有效的 JSON
- **THEN** 系统返回状态码 400，响应体为 `{"error": "invalid request body"}`

#### Scenario: 使用非 POST 方法
- **WHEN** 客户端使用 GET 或其他方法请求 `/urldecode`
- **THEN** 系统返回状态码 405（Method Not Allowed）