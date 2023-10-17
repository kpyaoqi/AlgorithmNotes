package SortAndSearch

import (
	"testing"
)

func TestMerge(t *testing.T) {
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//m := 3
	//nums2 := []int{2, 5, 6}
	//n := 3
	nums1 := []int{4, 0, 0, 0, 0, 0}
	m := 1
	nums2 := []int{1, 2, 3, 5, 6}
	n := 5
	//nums1 := []int{1}
	//m := 1
	//nums2 := []int{}
	//n := 0
	//nums1 := []int{0}
	//m := 0
	//nums2 := []int{1}
	//n := 1
	//nums1 := []int{1, 0}
	//m := 1
	//nums2 := []int{2}
	//n := 1

	Merge(nums1, m, nums2, n)
}
