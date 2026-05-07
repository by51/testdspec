# Change: 创建 Add 接口——字符串拼接

## 为什么（Why）
项目当前缺少通用字符串处理接口。需要新增一个 `/add` HTTP 接口，接收多个字符串参数，拼接后返回结果。

## 变更内容（What Changes）
- 新增 HTTP 路由 `POST /add`
- 接口接收 JSON body，字段为字符串数组 `items []string`
- 将所有字符串按顺序拼接，以 JSON 格式返回拼接结果

## 影响范围（Impact）
- `main.go`：新增路由注册和 handler 实现
- 无外部依赖变更
