package Others

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	HammingWeight(100)
}
func TestHammingDistance(t *testing.T) {
	HammingDistance(5, 9)
}
func TestReverseBits(t *testing.T) {
	num := uint32(10)
	reversed := reverseBits(num)
	fmt.Println("Reversed bits:", reversed)

	decimal := int(reversed)
	fmt.Println("Decimal:", decimal)
}
