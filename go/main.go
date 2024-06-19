package main

import (
	leetcode "app/leetcode/linked_list"
	"fmt"
)

func main() {
	node1 := &leetcode.Node{Val: 1}
	node2 := &leetcode.Node{Val: 2}
	node3 := &leetcode.Node{Val: 3}
	node4 := &leetcode.Node{Val: 4}
	node5 := &leetcode.Node{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	fmt.Println("Original:")
	leetcode.PrintLinkedList(node1)
	reversedHead := leetcode.ReverseSinglyLinkedList(node1)
	fmt.Println("Reversed:")
	leetcode.PrintLinkedList(reversedHead)
}
