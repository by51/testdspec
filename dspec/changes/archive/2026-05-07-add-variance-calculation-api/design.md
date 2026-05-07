# Design: 增加方差计算接口

## 技术方案
新增独立的 `mathutils` 包，封装方差计算逻辑，与 HTTP 层解耦。Handler 层仅负责解析请求和序列化响应，计算逻辑全部在 `mathutils.Variance()` 函数中实现。

方差计算采用**总体方差**（Population Variance）公式：
```
Var = Σ(xᵢ - mean)² / n
```

选择总体方差而非样本方差，因为接口语义是"对给定数据列表求方差"，数据集本身即为全体。

## 模块结构
```
mathutils/
  variance.go       # Variance([]float64) (float64, error) 函数
  variance_test.go  # 单元测试
main.go             # 新增 /variance 路由注册 + handler
```

## 关键决策
- **使用 `[]float64` 而非 `[]int`**：接口接受浮点数，覆盖更广的场景，JSON 数字统一反序列化为 float64
- **空列表返回 400**：方差在数学上对空集无定义，返回错误比返回 0 更语义准确
- **不引入外部依赖**：标准库 `math` 已足够，保持零外部依赖
