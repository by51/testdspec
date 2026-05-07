# Tasks: add-variance-calculation-api

## 实现任务

### mathutils 包
- [x] 实现 `mathutils/variance.go`：导出 `Variance(items []float64) (float64, error)`，空列表返回 error，单元素返回 0，其他按总体方差公式计算
- [x] 单测 `mathutils/variance_test.go`：覆盖正常多元素（结果 4.0）、单元素（结果 0.0）、空列表（返回 error）三个场景

### HTTP 层
- [x] 在 `main.go` 新增 `varianceHandler`：解析 JSON body `{"items": [...]}` → 调用 `mathutils.Variance` → 返回 `{"variance": ...}`；覆盖空列表 400、非法 body 400、非 POST 方法 405 的错误处理
- [x] 在 `main.go` 的 `mux` 注册 `POST /variance` 路由

### 集成
- [x] 集成测试（可在 `main_test.go` 或新文件）：验证 `POST /variance` 端到端返回正确方差值，并验证空列表返回 400
