## ADDED Requirements

### Requirement: 字符串拼接接口
系统应当提供 `POST /add` 接口，接收字符串数组，将所有元素按顺序拼接后返回。

#### Scenario: 正常拼接多个字符串
- **WHEN** 客户端发送 `POST /add`，body 为 `{"items": ["hello", " ", "world"]}`
- **THEN** 响应 200，body 为 `{"result": "hello world"}`

#### Scenario: 单个字符串
- **WHEN** 客户端发送 `POST /add`，body 为 `{"items": ["only"]}`
- **THEN** 响应 200，body 为 `{"result": "only"}`

#### Scenario: 空数组
- **WHEN** 客户端发送 `POST /add`，body 为 `{"items": []}`
- **THEN** 响应 200，body 为 `{"result": ""}`

#### Scenario: 请求体格式错误
- **WHEN** 客户端发送非 JSON body 或缺少 `items` 字段
- **THEN** 响应 400，body 含错误描述
