package leetcode

import "fmt"

func (l *List) Add(value int) {
	newNode := &Node{Val: value}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func (l *List) Remove(value int) {
	if l.Head == nil {
		return
	}

	if l.Head.Val == value {
		l.Head = l.Head.Next
		return
	}

	curr := l.Head
	for curr.Next != nil && curr.Next.Val != value {
		curr = curr.Next
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func (l *List) print() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func MakeList(input []int) *List {
	list := &List{}

	for _, v := range input {
		list.Add(v)
	}

	return list
}

func AddTwoNumbers(l1 *Node, l2 *Node) *Node {

	resList := &Node{}
	tmp := resList
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &Node{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &Node{}
		}
		tmp = tmp.Next
	}

	return resList

}
