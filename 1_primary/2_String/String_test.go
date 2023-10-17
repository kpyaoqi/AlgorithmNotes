package StringY

import (
	"Algorithm/utils"
	"fmt"
	"testing"
)

// 测试反转字符串
func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	utils.PrintfByteArr(s)
}

// 测试整数反转
func TestReverse(t *testing.T) {
	x1 := -123
	x2 := 123
	x3 := 120
	fmt.Println(reverse(x1))
	fmt.Println(reverse(x2))
	fmt.Println(reverse(x3))
}

// 字符串中的第一个唯一字符
func TestFirstUniqChar(t *testing.T) {
	s1 := "leetcode"
	s2 := "loveleetcode"
	fmt.Print(firstUniqChar(s1), firstUniqChar(s2))
}

// 测试有效的字母异位词
func TestIsAnagram(t *testing.T) {
	s1 := "abcabc"
	s2 := "cbacba"
	s3 := "abab"
	fmt.Printf("%t\n", isAnagram(s1, s2))
	fmt.Printf("%t\n", isAnagram(s2, s3))
}

// 测试验证回文串
func TestIsPalindrome(t *testing.T) {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	fmt.Printf("%t\n", isPalindrome(s1))
	fmt.Printf("%t\n", isPalindrome(s2))
}

// 测试字符串转换整数
func TestMyAtio(t *testing.T) {
	s1 := "   -42"
	s2 := "  +0 123"
	s3 := "00000-42a1234"
	fmt.Printf("%d\n", myAtoi(s1))
	fmt.Printf("%d\n", myAtoi(s2))
	fmt.Printf("%d\n", myAtoi(s3))
}

// 测试strStr()
func TestStrStr(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Print(strStr(haystack, needle))
}

// 测试外观数列
func TestCountAndSay(t *testing.T) {
	fmt.Print(countAndSay(100))
}

// 测试最长公共前缀
func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
