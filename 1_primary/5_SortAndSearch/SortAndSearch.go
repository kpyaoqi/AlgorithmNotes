package SortAndSearch

// 合并两个有序数组
func Merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		p1     int = 0
		p2     int = 0
		newArr     = make([]int, 0, m+n)
	)
	for {
		if p1 == m {
			newArr = append(newArr, nums2[p2:]...)
			break
		}
		if p2 == n {
			newArr = append(newArr, nums1[p1:]...)
			break
		}
		if nums1[p1] > nums2[p2] {
			newArr = append(newArr, nums2[p2])
			p2++
		} else {
			newArr = append(newArr, nums1[p1])
			p1++
		}
	}
	copy(nums1, newArr)
}

// 第一个错误的版本
func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(mid int) bool {
	return true
}
