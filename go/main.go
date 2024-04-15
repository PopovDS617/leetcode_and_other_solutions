package main

import (
	leetcode "app/leetcode/linked_list"
	"fmt"
)

func main() {
	l1List := []int{1, 2, 3}
	l2List := []int{1, 2, 3}

	l1 := leetcode.MakeList(l1List)
	l2 := leetcode.MakeList(l2List)

	res := leetcode.AddTwoNumbers(l1.Head, l2.Head)

	tmp := res
	for {
		fmt.Println(tmp.Val)

		if tmp.Next != nil {
			tmp = tmp.Next
		} else {
			fmt.Println("breaking")
			break
		}

	}

}
