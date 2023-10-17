package ArrayY

import (
	"fmt"
	"testing"
)

// 测试删除排序数组中的重复项
func TestRemoveDuplicates(t *testing.T) {
	nums1 := []int{1, 1, 2, 3, 3, 4, 4, 5}
	removeDuplicates(nums1)
}

// 测试买卖股票的最佳时机
func TestMaxProfit(t *testing.T) {
	nums2 := []int{7, 1, 2, 6, 5}
	maxProfit(nums2)
}

// 测试旋转数组
func TestRotate(t *testing.T) {
	nums3 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums3, k)
}

// 测试存在重复元素
func TestContainsDuplicate(t *testing.T) {
	nums4 := []int{5, 2, 3, 6, 1, 5, 2}
	fmt.Print(containsDuplicate(nums4))
}

// 测试只出现一次的数字
func TestSingleNumber(t *testing.T) {
	nums5 := []int{4, 1, 2, 1, 2}
	fmt.Print(singleNumber(nums5))
}

// 测试两个数组的交集
func TestIntersect(t *testing.T) {
	nums6 := []int{1, 2, 3, 2, 4, 5}
	nums6Y := []int{2, 4, 5, 2}
	intersect(nums6, nums6Y)
}

// 测试加一
func TestPlusOne(t *testing.T) {
	nums7 := []int{9, 9, 9, 9}
	fmt.Print(plusOne(nums7))
}

// 测试移动零
func TestMoveZeroes(t *testing.T) {
	nums8 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums8)
}

// 测试两树之和
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	sum := twoSum(nums, target)
	fmt.Println(sum)
	nums1 := []int{3, 2, 4}
	target1 := 6
	sum1 := twoSum(nums1, target1)
	fmt.Println(sum1)
}

// 测试有效数独
func TestIsValidSudoku(t *testing.T) {
	para := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Print(isValidSudoku(para))
}

// 测试旋转图像
func TestRotateImg(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateImg(matrix)
}
