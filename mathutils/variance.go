package mathutils

import (
  "errors"
  "math"
)

// Variance 计算总体方差，空列表返回 error
func Variance(items []float64) (float64, error) {
  if len(items) == 0 {
    return 0, errors.New("items must not be empty")
  }
  if len(items) == 1 {
    return 0, nil
  }

  var sum float64
  for _, v := range items {
    sum += v
  }
  mean := sum / float64(len(items))

  var variance float64
  for _, v := range items {
    diff := v - mean
    variance += diff * diff
  }
  return math.Round(variance/float64(len(items))*1e10) / 1e10, nil
}
