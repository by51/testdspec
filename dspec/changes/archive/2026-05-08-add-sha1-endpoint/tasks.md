# Tasks: add-sha1-endpoint

## 实现任务

### main.go
- [x] 实现 SHA1 哈希计算：添加 `Sha1Request`、`Sha1Response` 结构体和 `sha1Handler` 处理函数
- [x] 注册路由：在 `main()` 中添加 `mux.HandleFunc("/sha1", sha1Handler)`
- [x] 单测：覆盖正常哈希计算、空字符串、无效 JSON、非 POST 方法的场景

### API 规范
- [x] 更新 `dspec/specs/api/spec.md`：新增 SHA1 接口文档

### 集成
- [x] 集成测试：启动服务后验证 `POST /sha1` 端点行为符合预期