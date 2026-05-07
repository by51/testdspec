# Design: 增加随机数接口

## 技术方案
使用 Go 标准库 `net/http` 实现路由，`fmt.Sprintf` + `math/rand` 生成 12 位数字字符串。

选择 `math/rand`（伪随机）而非 `crypto/rand`：场景无安全要求，伪随机性能更好、代码更简洁。

## 模块结构
- `main.go`：在现有 HTTP server 中新增路由 `GET /random`，添加 `randomNumberHandler` 处理函数
- 无需新建文件，保持项目简洁

## 关键决策
- **返回格式**：纯文本（`Content-Type: text/plain`），不包装 JSON——接口简单，无需额外结构
- **12 位首位为 0**：用 `fmt.Sprintf("%012d", n)` 确保始终 12 位，`n` 取 `[0, 10^12)` 范围随机整数
- **种子初始化**：`main` 函数启动时调用 `rand.Seed(time.Now().UnixNano())`，避免每次重启返回相同序列
