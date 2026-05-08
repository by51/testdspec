# Design: 增加时间戳接口

## 技术方案
在 main.go 中新增一个 HTTP handler，使用 Go 标准库 `net/http` 和 `time` 包实现。

## 模块结构
- `main.go`：新增 `/timestamp` 路由和对应 handler 函数
- 响应结构体定义在 main.go 中

## 关键决策
- 使用标准库，无外部依赖
- 时间戳精度为秒级，满足一般同步需求
- 使用 RFC3339 格式作为可读时间字符串，便于解析和阅读
