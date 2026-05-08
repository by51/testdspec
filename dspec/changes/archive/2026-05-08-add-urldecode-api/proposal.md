# Change: 添加 URL 解码 API

## 为什么（Why）

当前系统缺少 URL 解码功能。在 Web 开发中，URL 编码的字符串（如 `%E4%BD%A0%E5%A5%BD`）需要解码后才能正常使用。提供一个 URL 解码 API 可以让客户端便捷地处理编码字符串，无需在各自系统中重复实现解码逻辑。

## 变更内容（What Changes）

1. 新增 `/urldecode` HTTP 接口，支持 POST 方法
2. 接受 JSON 格式请求体，包含 `input` 字段（待解码字符串）
3. 返回 JSON 格式响应，包含 `decoded` 字段（解码后字符串）
4. 支持标准 URL 解码（`net/url` 的 `QueryUnescape`）
5. 处理无效编码、无效 JSON 等错误场景

## 影响范围（Impact）

- **模块影响**：`main.go`（新增 handler）、可能扩展 `stringutils` 包
- **规范影响**：`api` 能力域需新增接口定义
- **依赖**：使用 Go 标准库 `net/url`，无外部依赖
- **风险点**：
  - 无效的 URL 编码格式需返回明确错误
  - 需考虑编码格式的边界情况（如 `%` 后缺少两位十六进制字符）