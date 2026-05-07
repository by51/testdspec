## Purpose
随机数字生成接口能力域，提供生成随机数字字符串的 HTTP 接口。

---

### Requirement: 随机数字生成接口
系统应当提供一个 HTTP 接口，生成并返回随机的 12 位纯数字字符串。

#### Scenario: 正常请求
- **WHEN** 客户端发送 `GET /random`
- **THEN** 返回 HTTP 200，响应体包含一个恰好 12 位的纯数字字符串

#### Scenario: 数字位数精确
- **WHEN** 接口返回随机数
- **THEN** 数字长度始终为 12 位（首位可以为 0）

#### Scenario: 每次请求结果不同
- **WHEN** 连续多次请求 `GET /random`
- **THEN** 大概率返回不同的值（允许极低概率碰撞）
