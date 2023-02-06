package stringUtil

import "strings"

// IsNotBlank 不为空
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// IsBlank 为空
func IsBlank(str string) bool {
	if len(strings.TrimSpace(str)) == 0 {
		return true
	}
	return false
}
