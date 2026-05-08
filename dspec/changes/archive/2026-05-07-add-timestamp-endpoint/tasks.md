# Tasks: add-timestamp-endpoint

## 实现任务

### 时间戳接口
- [x] 实现 timestampHandler：处理 GET `/timestamp` 请求，返回 JSON 格式时间戳信息
- [x] 定义响应结构体：包含 `unix`（int64）和 `readable`（string）字段

### 路由注册
- [x] 在 main 函数中注册 `/timestamp` 路由

### 测试
- [x] 手动测试：启动服务，使用 curl 验证接口返回正确的时间戳信息
