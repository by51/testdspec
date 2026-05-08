# Tasks: md5-encryption-endpoint

## 实现任务

### main.go
- [x] 定义请求/响应结构体：`Md5Request`（含 `Input string` 字段）和 `Md5Response`（含 `Hash string` 字段）
- [x] 实现 `md5Handler` 函数：校验方法为 POST，解析 JSON，计算 MD5 哈希，返回响应
- [x] 在 `main()` 中注册路由：`mux.HandleFunc("/md5", md5Handler)`

### 单元测试
- [x] 测试正常输入：验证 `hello world` 返回正确的哈希值
- [x] 测试空字符串：验证返回空字符串的 MD5 值
- [x] 测试缺少 input 字段：验证返回 400 错误
- [x] 测试无效 JSON：验证返回 400 错误
- [x] 测试非 POST 方法：验证返回 405 错误

### 集成
- [x] 启动服务，手动测试 `POST /md5` 端点