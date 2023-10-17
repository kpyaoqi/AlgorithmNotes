package ArrayY

import (
	"Algorithm/utils"
	"fmt"
	"sort"
)

// 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	arr := nums[:index]
	utils.PrintfIntArr(arr)
	return index
}

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	fmt.Printf("profit:%d", profit)
	return profit
}

// 旋转数组
func rotate(nums []int, k int) {
	numslen := len(nums)
	var arr []int = make([]int, numslen)
	for i := 0; i < numslen; i++ {
		arr[(i+k)%numslen] = nums[i]
	}
	for i, _ := range arr {
		nums[i] = arr[i]
	}
	utils.PrintfIntArr(arr)
}

// 存在重复元素
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

// 两个数组的交集
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j, index := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[index] = nums1[i]
			index++
			i++
			j++
		}
	}
	utils.PrintfIntArr(nums1[:index])
	return nums1[:index]
}

// 加一
func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

// 移动零
func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index++
	}
	//i := 0
	//for j := 0; j < len(nums); j++ {
	//	if nums[j] == 0 { //如果当前数字是0就不操作
	//		i++
	//	} else if i != 0 {
	//		//否则，把当前数字放到最前面那个0的位置，然后再把
	//		//当前位置设为0
	//		nums[j-i] = nums[j]
	//		nums[j] = 0
	//	}
	//}
	utils.PrintfIntArr(nums)
}

// 两树之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

// 有效数独
func isValidSudoku(board [][]byte) bool {
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				num := board[r][c] - '0' - byte(1)
				if rowbuf[r][num] || colbuf[r][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}
				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true
			}
		}
	}
	return true
}

// 旋转图像
func rotateImg(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}
