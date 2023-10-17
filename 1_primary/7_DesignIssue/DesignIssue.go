package DesignIssue

// 打乱数组
type Solution struct {
	original []int
	shuffled []int
}

func Constructor(nums []int) Solution {
	// 复制原始数组并初始化随机数生成器
	original := make([]int, len(nums))
	shuffled := make([]int, len(nums))
	copy(original, nums)
	copy(shuffled, nums)
	rand.Seed(time.Now().UnixNano())
	return Solution{original, shuffled}
}

func (this *Solution) Reset() []int {
	return this.original
}

//使用 Fisher-Yates 洗牌算法来打乱数组
func (this *Solution) Shuffle() []int {
	n := len(this.shuffled)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		this.shuffled[i], this.shuffled[j] = this.shuffled[j], this.shuffled[i]
	}
	return this.shuffled
}

// 最小栈
type MinStack struct {
    stack    []int // 主栈
    minStack []int // 辅助栈，用于保存最小元素
}

func MinStack() MinStack {
    return MinStack{}
}

func (ms *MinStack) Push(val int) {
    ms.stack = append(ms.stack, val)
    // 如果辅助栈为空或者 val 小于等于当前辅助栈的栈顶元素，则将 val 推入辅助栈
    if len(ms.minStack) == 0 || val <= ms.minStack[len(ms.minStack)-1] {
        ms.minStack = append(ms.minStack, val)
    }
}

func (ms *MinStack) Pop() {
    if len(ms.stack) == 0 {
        return
    }
    if ms.stack[len(ms.stack)-1] == ms.minStack[len(ms.minStack)-1] {
        ms.minStack = ms.minStack[:len(ms.minStack)-1]
    }
    ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
    if len(ms.stack) == 0 {
        return 0
    }
    return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
    if len(ms.minStack) == 0 {
        return 0
    }
    return ms.minStack[len(ms.minStack)-1]
}