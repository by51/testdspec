# Tasks: add-urldecode-api

## 实现任务

### main.go
- [x] 定义 `UrlDecodeRequest` 请求结构体：包含 `Input string` 字段（JSON tag: `input`）
- [x] 定义 `UrlDecodeResponse` 响应结构体：包含 `Decoded string` 字段（JSON tag: `decoded`）
- [x] 实现 `urldecodeHandler` 函数：处理 POST 请求，调用 `url.QueryUnescape` 解码，返回 JSON 响应
- [x] 注册路由：在 `main()` 函数中添加 `mux.HandleFunc("/urldecode", urldecodeHandler)`

### 单元测试
- [x] 测试正常解码：输入 `hello%20world`，期望输出 `hello world`
- [x] 测试中文字符解码：输入 `%E4%BD%A0%E5%A5%BD`，期望输出 `你好`
- [x] 测试特殊字符解码：输入 `a%3Db%26c%3Dd`，期望输出 `a=b&c=d`
- [x] 测试未编码字符串：输入 `hello world`，期望输出 `hello world`
- [x] 测试空字符串：输入空字符串，期望输出空字符串
- [x] 测试缺少 input 字段：请求体不含 input，期望返回空字符串
- [x] 测试无效 JSON：非 JSON 请求体，期望返回 400 错误
- [x] 测试无效 URL 编码：输入 `%ZZ`，期望返回 400 错误
- [x] 测试不完整的百分号编码：输入 `hello%2`，期望返回 400 错误
- [x] 测试非 POST 方法：GET 请求，期望返回 405 错误

### 集成测试
- [x] 启动服务，使用 curl 测试 `/urldecode` 接口的完整流程