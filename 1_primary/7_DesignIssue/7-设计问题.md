# 打乱数组

给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是等可能的。

实现 Solution class:

    Solution(int[] nums) 使用整数数组 nums 初始化对象
    int[] reset() 重设数组到它的初始状态并返回
    int[] shuffle() 返回数组随机打乱后的结果

### 示例 ：

输入
["Solution", "shuffle", "reset", "shuffle"]
[[[1, 2, 3]], [], [], []]
输出
[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]

解释
Solution solution = new Solution([1, 2, 3]);
solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3, 1, 2]
solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]

### 思路：
首先创建一个 Solution 结构体，该结构体包含原始数组 original 和打乱后的数组 shuffled。构造函数 Constructor 初始化这两个数组，并设置随机数生成器的种子。

Reset 方法返回原始数组，而 Shuffle 方法使用 Fisher-Yates 洗牌算法来打乱数组。请确保导入 "math/rand" 包并使用 rand.Seed(time.Now().UnixNano()) 来初始化随机数生成器。

# 最小栈

设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

    MinStack() 初始化堆栈对象。
    void push(int val) 将元素val推入堆栈。
    void pop() 删除堆栈顶部的元素。
    int top() 获取堆栈顶部的元素。
    int getMin() 获取堆栈中的最小元素。

### 示例 :

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

### 思路：
一个主栈 stack 用来存储元素，另一个辅助栈 minStack 用来保存每个状态下的最小元素。在 Push 操作中，如果元素小于等于当前辅助栈的栈顶元素，将元素推入辅助栈。在 Pop 操作中，如果弹出的元素等于辅助栈的栈顶元素，同时从辅助栈中也弹出。在 Top 和 GetMin 操作中，分别返回主栈的栈顶元素和辅助栈的栈顶元素，即最小元素。