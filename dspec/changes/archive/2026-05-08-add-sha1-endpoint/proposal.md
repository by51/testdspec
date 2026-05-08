# Change: 添加 SHA1 哈希计算端点

## 为什么（Why）
系统已提供 MD5 和 SHA256 哈希计算端点，但缺少 SHA1 哈希计算能力。SHA1 虽然不再推荐用于安全敏感场景，但在兼容旧系统、文件校验、Git 提交哈希等场景仍有广泛应用。提供 SHA1 端点可以完善系统的加密能力覆盖。

## 变更内容（What Changes）
- 新增 `POST /sha1` 端点
- 接收 JSON 请求体 `{"input": "待加密字符串"}`
- 返回 40 位小写十六进制 SHA1 哈希值

## 影响范围（Impact）
- **新增功能**：不修改现有接口，无破坏性变更
- **依赖模块**：使用 Go 标准库 `crypto/sha1`，无外部依赖
- **相关文件**：`main.go` 新增处理函数和路由注册
- **API 规范**：`dspec/specs/api/spec.md` 新增接口文档
- **加密规范**：`dspec/specs/encryption/spec.md` 新增 SHA1 需求