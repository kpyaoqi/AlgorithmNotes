package utils

// 判断 c 是否是字符或者数字
func IsChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
