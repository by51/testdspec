# Design: 字符串按空格拆分接口

## 技术方案
使用 Go 标准库 `strings.Fields(s string) []string`，该函数天然支持：
- 按任意空白字符（空格、tab 等）拆分
- 自动忽略连续空白和首尾空白
- 输入为空或仅含空白时返回空切片

无需引入外部依赖，直接封装为公共函数暴露。

## 模块结构
```
stringutils/
  split.go        # SplitBySpace 函数实现
  split_test.go   # 单元测试
```

函数签名：
```go
func SplitBySpace(s string) []string
```

## 关键决策
- **选择 `strings.Fields` 而非 `strings.Split(s, " ")`**：后者在连续空格时会产生空元素，需要额外过滤；`strings.Fields` 行为与规格完全一致，代码更简洁。
- **新建 `stringutils` 包**：保持职责单一，未来可扩展其他字符串工具函数。
- **函数命名用 `SplitBySpace`**：语义明确，区别于按其他分隔符拆分的场景。
