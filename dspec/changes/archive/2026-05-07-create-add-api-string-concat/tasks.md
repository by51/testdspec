# Tasks: create-add-api-string-concat

## 实现任务

### main.go
- [x] 实现 `addHandler`：解析 JSON body 中的 `items []string`，用 `strings.Join` 拼接，返回 `{"result": "..."}` JSON 响应
- [x] 注册路由 `POST /add` → `addHandler`
- [x] 单测：覆盖正常多字符串拼接、单字符串、空数组、错误请求体四个场景

### 集成
- [x] 集成测试：启动服务后发送真实 HTTP 请求，验证 `/add` 端到端返回正确拼接结果
