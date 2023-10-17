package utils

import "fmt"

func PrintfIntArr(nums []int) {
	for i, num := range nums {
		fmt.Printf("第%d个数是%d\n", i, num)
	}
}

func PrintfByteArr(nums []byte) {
	for i, num := range nums {
		fmt.Printf("第%d个数是%q\n", i, num)
	}
}
