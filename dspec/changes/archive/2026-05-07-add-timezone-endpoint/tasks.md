# Tasks: add-timezone-endpoint

## 实现任务

### main.go
- [x] 定义 TimezoneResponse 结构体：包含 timezone（时区名称）和 offset（偏移量）字段
- [x] 实现 timezoneHandler 函数：获取系统时区并返回 JSON 响应
- [x] 注册路由：在 main() 中添加 `/timezone` 路由

### 测试
- [x] 手动测试：使用 curl 测试 GET 请求返回正确的时区信息
- [x] 手动测试：验证非 GET 请求返回 405 错误
