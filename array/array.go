package main

import "fmt"

// ListNode 表示单向链表的一个节点，只保存数据和指向下一个节点的指针。
// 最后一个节点的 Next 为 nil；空链表的头节点也为 nil。
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 使用快慢指针（Floyd 算法）判断单向链表是否有环。
// 慢指针每次走一步，快指针每次走两步：
//   - 无环时，快指针会走到链表末尾，返回 false。
//   - 有环时，两个指针进入环后，快指针每轮相对慢指针多走一步，最终会相遇。
//
// 注意：比较的是节点指针是否相同，而不是节点的值是否相同。
// 时间复杂度 O(n)，空间复杂度 O(1)；空链表返回 false。
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	// 先检查 fast 和 fast.Next，确保向前走两步时不会解引用空指针。
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 必须先移动再比较，因为初始化时两个指针都指向 head。
		if slow == fast {
			return true
		}
	}
	return false
}

// reverseList 原地反转一个无环的单向链表，并返回新的头节点。
// 例如：1 -> 2 -> 3 -> nil，反转后为 3 -> 2 -> 1 -> nil。
// 前提：链表无环；若不确定，应先使用 hasCycle 检查。
// 该函数会修改原节点的 Next，调用方需要用返回值更新头节点。
// 时间复杂度 O(n)，空间复杂度 O(1)；空链表返回 nil，单节点链表保持不变。
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // 已反转部分的头节点，初始为空。
	current := head    // 当前待反转的节点。
	for current != nil {
		next := current.Next // 1. 保存原后继，避免修改指针后丢失剩余链表。
		current.Next = prev  // 2. 将当前节点指向前一个节点，完成指针反转。
		prev = current       // 3. 已反转部分的头节点前移。
		current = next       // 4. 继续处理原链表中的下一个节点。
	}
	return prev // current 为 nil 时，prev 就是原尾节点，即新的头节点。
}

func main() {
	// 创建 int 切片：长度为 3，容量为 5，元素默认值为 0
	s := make([]int, 3, 5)

	fmt.Printf("切片变量的地址：%p\n", &s)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	// s[3] = 40 // 这里会报错，索引越界，因为切片的长度是 3，不能直接访问索引为 3 的元素
	s = append(s, 40)

	fmt.Println(s)      // [10 20 30 40]
	fmt.Println(len(s)) // 4，当前元素数量
	fmt.Println(cap(s)) // 5，底层数组的容量
	fmt.Println("------------")
	fmt.Printf("切片变量的地址：%p\n", &s)
	s = append(s, 40)
	s = append(s, 40)
	s = append(s, 40)
	s = append(s, 40)
	s = append(s, 40)
	fmt.Printf("切片变量的地址：%p\n", &s)

	fmt.Println("------------ 单向链表示例 ------------")
	// 创建无环链表：1 -> 2 -> 3 -> nil。
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	fmt.Println("反转前是否有环：", hasCycle(head)) // false
	head = reverseList(head)                // 接收新头节点，链表变为 3 -> 2 -> 1 -> nil。
	fmt.Print("反转后的链表：")
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Val)
	}
	fmt.Println("nil") // 3 -> 2 -> 1 -> nil

	// 将尾节点 1 指向中间节点 2，构造环：3 -> 2 -> 1 -> 2 -> ...。
	tail := head.Next.Next
	tail.Next = head.Next
	fmt.Println("构造环后是否有环：", hasCycle(head)) // true
	// 有环时不能按上面的方式遍历到 nil；再次反转前也应先断开环。
	tail.Next = nil
}
