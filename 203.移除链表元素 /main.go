/**
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。



 示例 1：


输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]


 示例 2：


输入：head = [], val = 1
输出：[]


 示例 3：


输入：head = [7,7,7,7], val = 7
输出：[]




 提示：


 列表中的节点数目在范围 [0, 10⁴] 内
 1 <= Node.val <= 50
 0 <= val <= 50


 Related Topics 递归 链表 👍 1173 👎 0

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归解法
func case1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = case1(head.Next, val)

	if head.Val != val {
		return head
	} else {
		return head.Next
	}
}

// while 解法
func case2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}

	dummyHead.Next = head
	cur := dummyHead

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)
