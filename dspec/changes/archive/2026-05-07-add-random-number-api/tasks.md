# Tasks: add-random-number-api

## 实现任务

### random-api
- [x] 实现 `randomNumberHandler`：生成 `[0, 10^12)` 范围随机整数，用 `fmt.Sprintf("%012d", n)` 格式化为 12 位字符串，写入响应
- [x] 在 `main.go` 中注册路由 `GET /random` → `randomNumberHandler`
- [x] 确保 `rand.Seed(time.Now().UnixNano())` 在 main 函数中初始化（若 Go 版本 ≥ 1.20 可省略）
- [x] 单测：覆盖返回值长度始终为 12 位、值为纯数字、多次调用大概率不同

### 集成
- [x] 集成测试：启动服务后 `curl localhost:<port>/random`，验证响应 200 且长度为 12
