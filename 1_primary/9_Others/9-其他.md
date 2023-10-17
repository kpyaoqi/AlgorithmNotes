# 位1的个数

编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。

### 示例 ：

输入：n = 00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

### 思路：

通过strconv.FormatUint(uint64(num), 2)转，循环当为1时++

# 汉明距离

两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

### 示例：

输入：x = 1, y = 4
输出：2
解释：
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
上面的箭头指出了对应二进制位不同的位置。

### 思路：

```go
for temp := x ^ y; temp > 0; temp >>= 1 {
    ret += (1 & temp)
}
```

# 颠倒二进制位

颠倒给定的 32 位无符号整数的二进制位。

### 示例：

输入：n = 00000010100101000001111010011100
输出：964176192 (00111001011110000010100101000000)
解释：输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
     因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。

### 思路：

```go
var result uint32
for i := 0; i < 32; i++ {
    result <<= 1    // 左移一位，为下一位腾出空间
    result |= n & 1 // 将 n 的最后一位加到 result 上
    n >>= 1         // 右移一位，准备处理下一位
}
return result
```

# 杨辉三角

给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

### 示例 :

输入: numRows = 5，输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

### 思路：

```go
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
```

# 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

### 示例 ：

输入：s = "()[]{}"
输出：true

### 思路：

```go
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
```

# 缺失数字

给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

### 示例 ：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。

### 思路：

```go
n := len(nums)
expectedSum := n * (n + 1) / 2
actualSum := 0
for _, num := range nums {
    actualSum += num
}
return expectedSum - actualSum
```

