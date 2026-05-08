# Change: 增加时区查询接口

## 为什么（Why）
客户端需要获取服务器的时区信息，用于时间相关的业务逻辑处理或展示。

## 变更内容（What Changes）
- 新增 `/timezone` HTTP GET 接口
- 返回系统当前时区名称（如 `Asia/Shanghai`）
- 返回时区偏移量（如 `UTC+8`）

## 影响范围（Impact）
- 新增接口，不影响现有功能
- 涉及模块：`main.go`（新增 handler 和路由注册）
