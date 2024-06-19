package leetcode

func ReverseSinglyLinkedList(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		Next := curr.Next
		curr.Next = prev
		prev = curr
		curr = Next
	}
	return prev
}
