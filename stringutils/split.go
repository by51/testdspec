package stringutils

import "strings"

// SplitBySpace 按空格拆分字符串，连续空格和首尾空格不产生空元素
func SplitBySpace(s string) []string {
	return strings.Fields(s)
}
