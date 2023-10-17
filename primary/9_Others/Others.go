package Others

import (
	"strconv"
)

// 位1的个数
func HammingWeight(num uint32) int {
	binaryStr := strconv.FormatUint(uint64(num), 2)
	count := 0
	for _, ch := range binaryStr {
		if ch == '1' {
			count++
		}
	}
	return count
}

// 汉明距离
func HammingDistance(x int, y int) (ret int) {
	for temp := x ^ y; temp > 0; temp >>= 1 {
		ret += (1 & temp)
	}
	return
}

// 颠倒二进制位
func reverseBits(n uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result <<= 1    // 左移一位，为下一位腾出空间
		result |= n & 1 // 将 n 的最后一位加到 result 上
		n >>= 1         // 右移一位，准备处理下一位
	}
	return result
}

// 杨辉三角
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

// 有效的括号
func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
