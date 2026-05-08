# Design: 添加 SHA256 哈希计算端点

## 技术方案
采用与现有 MD5 端点一致的实现方式：
- 在 main.go 中新增 `sha256Handler` 处理函数
- 使用 Go 标准库 `crypto/sha256` 计算哈希值
- 复用现有的请求/响应结构设计模式

选择此方案的原因：
1. 与现有代码风格一致，降低维护成本
2. 标准库稳定可靠，无需引入外部依赖
3. 实现简单直接，无需复杂的抽象层

## 模块结构
```
main.go
├── Sha256Request    // 请求结构（新增）
├── Sha256Response   // 响应结构（新增）
└── sha256Handler    // 处理函数（新增）
```

### 新增代码说明

**请求结构**
```go
type Sha256Request struct {
    Input string `json:"input"` // 待加密的字符串
}
```

**响应结构**
```go
type Sha256Response struct {
    Hash string `json:"hash"` // SHA256 哈希值（64位小写十六进制）
}
```

**处理函数**
- 路由：`POST /sha256`
- 方法校验：只接受 POST 请求
- 请求体解析：JSON 格式
- 哈希计算：`sha256.Sum256([]byte(input))`
- 响应格式：JSON

## 关键决策
1. **不使用统一接口**：虽然 MD5 和 SHA256 功能相似，但不创建统一的加密接口，保持简单独立，避免过度设计
2. **请求结构独立**：新建 `Sha256Request` 而非复用 `Md5Request`，便于未来各自扩展（如添加盐值、迭代次数等）
3. **响应格式一致**：保持 `{"hash": "..."}` 格式与 MD5 接口一致，便于客户端统一处理