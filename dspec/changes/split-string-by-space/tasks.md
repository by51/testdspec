# Tasks: split-string-by-space

## 实现任务

### stringutils 包
- [x] 创建 `stringutils/split.go`：实现 `SplitBySpace(s string) []string`，使用 `strings.Fields` 拆分
- [x] 创建 `stringutils/split_test.go`：单测覆盖以下场景
  - 普通多词字符串拆分
  - 连续空格（不产生空元素）
  - 首尾有空格
  - 空字符串输入 → 返回空切片
  - 仅含空格 → 返回空切片
  - 单词无空格 → 返回单元素数组

### 集成
- [x] 在 `main.go` 中添加调用示例，验证函数可正常导入和使用
