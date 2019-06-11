package _

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	//l128 := &ListNode{1, nil}
	//l127 := &ListNode{0, l128}
	//l126 := &ListNode{0, l127}
	//l125 := &ListNode{0, l126}
	//l124 := &ListNode{0, l125}
	//l123 := &ListNode{0, l124}
	//l122 := &ListNode{0, l123}
	//l121 := &ListNode{0, l122}
	//l120 := &ListNode{0, l121}
	//l119 := &ListNode{0, l120}
	//l118 := &ListNode{0, l119}
	//l117 := &ListNode{0, l118}
	//l116 := &ListNode{0, l117}
	//l115 := &ListNode{0, l116}
	//l114 := &ListNode{0, l115}
	//l113 := &ListNode{0, l114}
	//l112 := &ListNode{0, l113}
	//l111 := &ListNode{0, l112}
	//l110 := &ListNode{0, l111}
	//l19 := &ListNode{0, l110}
	//l18 := &ListNode{0, l19}
	//l17 := &ListNode{0, l18}
	//l16 := &ListNode{0, l17}
	//l15 := &ListNode{0, l16}
	//l14 := &ListNode{0, l15}
	l13 := &ListNode{3, nil}
	l12 := &ListNode{4, l13}
	l11 := &ListNode{2, l12}

	l23 := &ListNode{4, nil}
	l22 := &ListNode{6, l23}
	l21 := &ListNode{5, l22}

	addTwoNumbers(l11, l21)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	slice1 := make([]int, 0, 1)
	slice2 := make([]int, 0, 1)
	tmp := l1
	for tmp != nil {
		slice1 = append(slice1, tmp.Val)
		tmp = tmp.Next
	}
	tmp = l2
	for tmp != nil {
		slice2 = append(slice2, tmp.Val)
		tmp = tmp.Next
	}

	//slice1 := []int{1, 2, 5, 6}
	//slice2 := []int{1, 2, 9, 1}

	fmt.Println(slice1)
	fmt.Println(slice2)

	reverse(slice1)
	reverse(slice2)

	fmt.Println(slice1)
	fmt.Println(slice2)

	sum := add(slice1, slice2)
	//str := strconv.FormatUint(sum, 10)

	fmt.Println(sum)

	head := &ListNode{}
	pos := head

	for count := 0; count < len(sum); count++ {
		//char := str[count-1 : count]
		//num, _ := strconv.Atoi(char)
		//fmt.Println(num)
		node := &ListNode{sum[count], nil}
		pos.Next = node
		pos = node
	}

	for head.Next != nil {
		fmt.Println(head.Next.Val)
		head.Next = head.Next.Next
	}
	return head.Next
}

func reverse(data []int) []int {

	for left, right := 0, len(data)-1; left < right; left, right = left+1, right-1 {
		data[left], data[right] = data[right], data[left]
	}
	return data
}

func add(a []int, b []int) []int {
	lenA := len(a)
	lenB := len(b)
	var lenShort = lenB
	var lenLong = lenA
	var short = b
	var long = a
	if lenA < lenB {
		lenShort = lenA
		lenLong = lenB
		long = b
		short = a
	}

	var lenDiff = lenLong - lenShort
	var carry = 0
	var mod = 0
	sum := make([]int, 0, 1)
	for count := lenShort; count > 0; count-- {
		tmpShort := short[count-1]
		tmpLong := long[lenDiff+count-1]
		tmpSum := tmpShort + tmpLong + carry
		carry = tmpSum / 10
		mod = tmpSum % 10
		sum = append(sum, mod)
	}

	for count := lenDiff; count > 0; count-- {
		tmp := long[count-1] + carry
		carry = tmp / 10
		mod = tmp % 10
		sum = append(sum, mod)
	}
	if carry != 0 {
		sum = append(sum, carry)
	}
	return sum
}
