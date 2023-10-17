package LinkedList

import (
	"testing"
)

// 测试删除链表中的节点
func TestDeleteNode(t *testing.T) {
	// 定义一个数组
	nums := []int{1, 2, 3, 4}
	node := ListNode{Val: 2}
	// 根据数组建立链表
	head := BuildLinkedList(nums)
	deleteNode(&node)
	// 打印链表
	PrintList(head)
}

func TestReverseList(t *testing.T) {
	// 定义一个数组
	nums := []int{1, 2, 3, 4}
	// 根据数组建立链表
	head := BuildLinkedList(nums)
	list := reverseList(head)
	PrintList(list)

}
