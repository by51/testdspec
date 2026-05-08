# Spec: encryption

## Purpose

提供加密相关的能力，包括哈希计算等功能。

## Requirements

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

### Requirement: SHA256 哈希计算
系统应当提供 SHA256 哈希计算能力，将输入字符串转换为 64 位小写十六进制哈希值。SHA256 是比 MD5 更安全的哈希算法，适用于安全敏感场景。

#### Scenario: 正常计算 SHA256 哈希
- **WHEN** 客户端发送 `POST /sha256` 请求，请求体包含 `{"input": "hello world"}`
- **THEN** 系统返回 `{"hash": "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"}`，状态码为 200

#### Scenario: 空字符串输入
- **WHEN** 客户端发送 `POST /sha256` 请求，请求体包含 `{"input": ""}`
- **THEN** 系统返回 `{"hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}`，状态码为 200

#### Scenario: 缺少 input 字段
- **WHEN** 客户端发送 `POST /sha256` 请求，请求体不包含 `input` 字段
- **THEN** 系统返回空字符串的 SHA256 哈希值，状态码为 200

#### Scenario: 无效 JSON 格式
- **WHEN** 客户端发送 `POST /sha256` 请求，请求体不是有效的 JSON
- **THEN** 系统返回 `{"error": "invalid request body"}`，状态码为 400

#### Scenario: 使用非 POST 方法
- **WHEN** 客户端使用 GET 或其他方法请求 `/sha256`
- **THEN** 系统返回状态码 405（Method Not Allowed）

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