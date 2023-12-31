# 二叉树的最大深度

给定一个二叉树，找出其最大深度。二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

### 示例：

给定二叉树 [3,9,20,null,null,15,7]，

​	3

   / \
  9  20
    /  \
   15   7

返回它的最大深度 3 。

### 思路：

递归遍历就可，遍历根节点的左孩子的高度和根节点右孩子的高度，取出两者的最大值再加一即为总高度。

```go
return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
```

# 验证二叉搜索树

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

### 示例 ：

输入：root = [2,1,3]
输出：true

### 思路：

```go
func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}
//二：中序遍历到数组，看数组是否有序
func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}
```

# 对称二叉树

给你一个二叉树的根节点 `root` ， 检查它是否轴对称。

### 示例：

![image-20230612192528623](img/image-20230612192528623.png) 

### 思路：

```go
func isMirror(left *TreeNode, right *TreeNode) bool {
   if left == nil && right == nil {
      return true
   }
   if left == nil || right == nil {
      return false
   }
   return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
```

# 二叉树的层序遍历

给你二叉树的根节点 `root` ，返回其节点值的 **层序遍历** 。 （即逐层地，从左到右访问所有节点）

### 示例：

![image-20230612193627829](img/image-20230612193627829.png) 

输出：[[3],[9,20],[15,7]]

### 思路：

```go
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	var dfsLevel func(node *TreeNode, level int)
	dfsLevel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		dfsLevel(node.Left, level+1)
		dfsLevel(node.Right, level+1)
	}
	dfsLevel(root, 0)
	return res
}
```

# 将有序数组转换为二叉搜索树

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

### 示例：

![image-20230612200628809](img/image-20230612200628809.png) 

### 思路：

```go
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{Val: nums[len(nums)/2], Left: sortedArrayToBST(nums[:len(nums)/2]), Right: sortedArrayToBST(nums[len(nums)/2+1:])}
}
```

