# Tasks: add-millisecond-timestamp-field

## 实现任务

### main.go
- [x] 修改 `TimestampResponse` 结构体：新增 `UnixMs int64` 字段，JSON tag 为 `unix_ms`
- [x] 修改 `timestampHandler` 函数：使用 `time.Now().UnixMilli()` 计算毫秒级时间戳并赋值给 `UnixMs` 字段

### 测试
- [x] 手动测试：启动服务，请求 `GET /timestamp` 验证响应包含 `unix_ms` 字段
- [x] 验证字段值正确性：`unix_ms` 约等于 `unix * 1000`（误差在毫秒级）

### 集成
- [x] 端到端验证：确认接口向后兼容，现有客户端可正常使用
