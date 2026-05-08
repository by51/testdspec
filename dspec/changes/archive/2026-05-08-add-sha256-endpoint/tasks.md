# Tasks: add-sha256-endpoint

## 实现任务

### main.go
- [x] 添加 `crypto/sha256` 导入
- [x] 实现 Sha256Request 结构体：包含 `input` 字段
- [x] 实现 Sha256Response 结构体：包含 `hash` 字段
- [x] 实现 sha256Handler 函数：处理 POST 请求，计算 SHA256 哈希，返回 JSON 响应
- [x] 在 main() 中注册路由：`mux.HandleFunc("/sha256", sha256Handler)`

### 单元测试
- [x] 测试正常输入：验证 "hello world" 返回正确的哈希值
- [x] 测试空字符串输入：验证返回空字符串的 SHA256 哈希
- [x] 测试缺少 input 字段：验证返回错误响应
- [x] 测试无效 JSON：验证返回错误响应
- [x] 测试非 POST 方法：验证返回 405 状态码

### 主规范同步
- [x] 更新 dspec/specs/api/spec.md：添加 SHA256 接口文档（第 11 节）
- [x] 更新 dspec/specs/encryption/spec.md：添加 SHA256 需求

### 集成验证
- [x] 运行测试：`go test ./...`
- [x] 启动服务验证：`go run main.go` 后调用 `/sha256` 接口