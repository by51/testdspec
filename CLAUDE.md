# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 常用命令

```bash
# 运行程序
go run main.go

# 构建
go build -o testdespec .

# 运行测试
go test ./...

# 运行单个测试
go test -run TestFunctionName ./...

# 格式化代码
go fmt ./...

# 静态检查
go vet ./...
```

## 项目结构

这是一个标准 Go 模块项目（module: `testdespec`，Go 1.25）：

- `main.go` — 程序入口，`package main`
- `go.mod` — 模块定义，目前无外部依赖
