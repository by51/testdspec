# Tasks: improve-ping-time-based-response

## 实现任务

### ping handler
- [x] 修改 `main.go` 中 `/ping` handler：获取当前小时数，`hour >= 10 && hour < 18` 时返回 `pong`，否则返回 `pang`
- [x] 单测：在 `main_test.go` 中为 ping 逻辑编写单元测试，覆盖工作时段（hour=10、17）、非工作时段（hour=9、18）两类场景
