package stringutils

import (
	"reflect"
	"testing"
)

func TestSplitBySpace(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{"普通多词字符串", "hello world foo", []string{"hello", "world", "foo"}},
		{"连续空格", "hello  world", []string{"hello", "world"}},
		{"首尾有空格", "  hello world  ", []string{"hello", "world"}},
		{"空字符串", "", []string{}},
		{"仅含空格", "   ", []string{}},
		{"单词无空格", "hello", []string{"hello"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitBySpace(tt.input)
			// strings.Fields 在空输入时返回 nil，统一转为空切片比较
			if got == nil {
				got = []string{}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitBySpace(%q) = %v，期望 %v", tt.input, got, tt.want)
			}
		})
	}
}
