# Design: 创建 Add 接口——字符串拼接

## 技术方案
在 `main.go` 中直接注册 `POST /add` handler，使用标准库 `encoding/json` 解析请求体，`strings.Join` 拼接字符串并返回 JSON 响应。不引入额外文件或包，保持与现有代码风格一致。

## 模块结构
- `main.go`
  - 新增 `addHandler(w http.ResponseWriter, r *http.Request)` 函数
  - 在 `main()` 中注册路由 `mux.HandleFunc("/add", addHandler)`

请求/响应结构：
```
Request  Body: {"items": ["str1", "str2", ...]}
Response Body: {"result": "str1str2..."}
```

## 关键决策
- **POST 而非 GET**：字符串列表作为 body 传递，避免 URL 长度限制和编码问题
- **`strings.Join` 拼接**：直接按顺序拼接，无分隔符，符合"拼接到一起"的需求
- **只用标准库**：项目目前无外部依赖，维持零依赖原则
