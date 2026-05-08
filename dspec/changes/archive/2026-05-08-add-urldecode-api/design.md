# Design: 添加 URL 解码 API

## 技术方案

采用直接在 `main.go` 中添加 handler 的方式实现，与现有 md5/sha256/sha1 等接口保持一致的风格。使用 Go 标准库 `net/url.QueryUnescape` 进行 URL 解码，无需引入外部依赖。

## 模块结构

### main.go
- 新增 `UrlDecodeRequest` 请求结构体，包含 `Input` 字段
- 新增 `UrlDecodeResponse` 响应结构体，包含 `Decoded` 字段
- 新增 `urldecodeHandler` 函数，处理 POST 请求并返回解码结果
- 在 `main()` 函数中注册 `/urldecode` 路由

### 请求处理流程
1. 校验请求方法为 POST，否则返回 405
2. 解析 JSON 请求体，解析失败返回 400
3. 调用 `url.QueryUnescape` 解码字符串
4. 解码失败（无效编码）返回 400，错误信息为 "invalid URL encoding"
5. 成功返回 200，JSON 响应包含解码后的字符串

## 关键决策

### 为什么不扩展 stringutils 包
现有 stringutils 包提供的是纯字符串处理功能（如 SplitBySpace），而 URL 解码是 HTTP 协议相关的编码处理，直接在 handler 中调用标准库更简洁，无需额外抽象层。

### 错误处理策略
- 无效 JSON：返回 400，错误信息 "invalid request body"
- 无效 URL 编码：返回 400，错误信息 "invalid URL encoding"
- 缺少 input 字段：返回 200，空字符串作为默认值（与 md5/sha256 等接口行为一致）

### 响应字段命名
使用 `decoded` 而非 `result`，语义更明确，表明这是解码操作的结果。