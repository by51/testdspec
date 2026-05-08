# Server 规范

## Purpose

描述 HTTP 服务器的基本行为，包括监听端口配置等。

---

### Requirement: HTTP 服务监听端口
系统应当在端口 10001 上监听 HTTP 请求。

#### Scenario: 服务启动
- **WHEN** 服务启动
- **THEN** HTTP Server 绑定到 `:10001`，启动日志输出 `服务启动，监听 :10001`
