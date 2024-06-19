package leetcode

import "fmt"

type List struct {
	Head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func PrintLinkedList(l *Node) {
	var result []int
	cur := l

	for cur != nil {

		result = append(result, cur.Val)

		cur = cur.Next
	}

	fmt.Println(result)
}
