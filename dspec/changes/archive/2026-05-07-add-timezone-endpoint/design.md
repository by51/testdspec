# Design: 增加时区查询接口

## 技术方案
使用 Go 标准库 `time` 包获取系统时区：
- `time.Now().Location()` 获取当前时区
- `time.Now().Zone()` 获取时区名称和偏移秒数
- 偏移秒数转换为 `+HHMM` 格式

## 模块结构
```
main.go
├── TimezoneResponse 结构体  # 响应数据结构
├── timezoneHandler()        # 处理函数
└── main()                   # 路由注册
```

## 关键决策
- **响应格式**：JSON，与现有 `/timestamp` 接口风格一致
- **时区偏移格式**：使用 `+0800` 而非 `UTC+8`，更符合 ISO 8601 标准
- **错误处理**：仅支持 GET 方法，其他方法返回 405
