package Math

import (
	"sort"
	"strconv"
)

// Fizz Buzz
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		num := i + 1
		if num%15 == 0 {
			res[i] = "FizzBuzz"
		} else if num%3 == 0 {
			res[i] = "Fizz"
		} else if num%5 == 0 {
			res[i] = "Buzz"
		} else {
			res[i] = strconv.Itoa(num)
		}
	}
	return res
}

// 计数质数
var primes []int

const primeList int = 5e6

func init() {
	limit := primeList
	flags := make([]bool, limit+1)
	primes = append(primes, 2)
	for i := 3; i <= limit; i += 2 {
		if flags[i] == false {
			primes = append(primes, i)
		}
		for _, prime := range primes {
			if i*prime > limit {
				break
			}
			flags[i*prime] = true
		}
	}
}

func countPrimes(n int) int {
	return sort.SearchInts(primes, n)
}

// 3的幂
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n > 3 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return n == 3
}

// 罗马数字转整数
func romanToInt(s string) int {
	byte2Int := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	prev := 0 // 用于存储前一个字符的值
	for i := len(s) - 1; i >= 0; i-- {
		cur := byte2Int[s[i]]
		if cur < prev {
			result -= cur // 当前字符小于前一个字符时，减去当前字符的值
		} else {
			result += cur // 否则加上当前字符的值
		}
		prev = cur // 更新前一个字符的值
	}
	return result
}
