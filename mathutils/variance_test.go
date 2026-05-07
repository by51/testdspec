package mathutils

import (
  "testing"
)

func TestVariance(t *testing.T) {
  // 正常多元素：经典示例，方差为 4.0
  result, err := Variance([]float64{2, 4, 4, 4, 5, 5, 7, 9})
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  if result != 4.0 {
    t.Errorf("expected 4.0, got %f", result)
  }

  // 单元素：方差为 0
  result, err = Variance([]float64{42})
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  if result != 0 {
    t.Errorf("expected 0, got %f", result)
  }

  // 空列表：返回 error
  _, err = Variance([]float64{})
  if err == nil {
    t.Error("expected error for empty list, got nil")
  }
}
