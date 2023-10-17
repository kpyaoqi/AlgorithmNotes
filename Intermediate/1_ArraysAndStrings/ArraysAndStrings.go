package __ArraysAndStrings

import (
	"Algorithm/utils"
	"sort"
)

// 三数之和 O(n^2)
func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	if n < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 避免重复的元素
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

// 矩阵置零 O(m * n)
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	row := make([]bool, m)
	col := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 字母异位词分组O(n * k)其中 n 是字符串数组的长度，k 是字符串的平均长度。
func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}
	res := [][]string{}

	sign := func(s string) string {
		strB := [26]byte{}
		for _, v := range s {
			strB[v-'a']++
		}
		return string(strB[:])
	}
	for _, v := range strs {
		signV := sign(v)
		hashMap[signV] = append(hashMap[signV], v)
	}

	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}

// 无重复字符的最长子串O(n)
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	charSet := make(map[byte]int)
	maxLength := 0
	left := 0

	for right := 0; right < n; right++ {
		char := s[right]
		if _, exists := charSet[char]; exists {
			left = utils.Max(charSet[char]+1, left)
		}
		charSet[char] = right
		maxLength = utils.Max(maxLength, right-left+1)
	}

	return maxLength
}

// 最长回文子串
func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := utils.Max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 递增的三元子序列 O(n)
func increasingTriplet(nums []int) bool {
	min := int(^uint(0) >> 1)       // 初始化为最大整数
	secondMin := int(^uint(0) >> 1) // 初始化为最大整数
	for _, num := range nums {
		if num <= min {
			min = num
		} else if num <= secondMin {
			secondMin = num
		} else {
			return true
		}
	}
	return false
}
