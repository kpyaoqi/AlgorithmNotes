package StringY

import (
	"Algorithm/utils"
	"strconv"
	"strings"
)

// 反转字符串
func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

// 整数反转
func reverse(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

// 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if result[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

// 有效的字母异位词
func isAnagram(s string, t string) bool {
	chars := map[rune]int{}
	for _, value := range s {
		chars[value]++
	}
	for _, value := range t {
		chars[value]--
	}
	for _, value := range chars {
		if value != 0 {
			return false
		}
	}
	return true
}

// 验证回文串
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !utils.IsChar(s[i]) {
			i++
		}
		for i < j && !utils.IsChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 字符串转换整数
func myAtoi(s string) int {
	sign := 1
	issign := false
	notk := false
	num := 0
	for _, c := range s {
		if c == ' ' && !notk {
			continue
		}
		if c == '+' && !issign {
			issign = true
			continue
		}
		if c == '-' && !issign {
			issign = true
			sign = -1
			continue
		}
		if '0' <= c && c <= '9' {
			num = num*10 + (int(c) - 48)
		} else {
			return int32o(sign * num)
		}
	}
	return int32o(sign * num)
}

func int32o(num int) int {
	if num > (1<<31)-1 {
		return (1 << 31) - 1
	} else if num < -1*(1<<31) {
		return -1 * (1 << 31)
	} else {
		return num
	}
}

// 实现 strStr()
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
	//return strings.Index(haystack, needle)
}

// 外观数列
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return transfer(countAndSay(n - 1))
}

func transfer(s string) string {
	var sb string
	count := 1
	length := len(s)
	i := 1
	temp := s[0]
	for i <= length {
		for i < length && temp == s[i] {
			count++
			i++
		}
		if i < length {
			temp = s[i]
		}

		sb += strconv.Itoa(count)
		sb += string(s[i-1])

		count = 1
		i++
	}
	return sb
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}
