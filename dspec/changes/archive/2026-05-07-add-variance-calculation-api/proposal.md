# Change: 增加方差计算接口

## 为什么（Why）
当前服务仅提供字符串处理等基础功能，缺乏数值统计计算能力。需要新增一个接口，支持对一组数据列表计算其方差，满足统计分析场景需求。

## 变更内容（What Changes）
- 新增 HTTP 接口 `POST /variance`，接受数字列表，返回方差值
- 新增 `mathutils` 包，封装方差计算逻辑
- 在 `main.go` 中注册新路由

## 影响范围（Impact）
- 新增文件：`mathutils/variance.go`、`mathutils/variance_test.go`
- 修改文件：`main.go`（注册路由）
- 无破坏性变更，不影响现有接口
