package main 

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func String(i int) string {
    return fmt.Sprintf("%d", i)
}

func buildLL(val int) *ListNode {
    s := String(val);
    N := len(s)
    if N == 0 {
        return nil
    }
    var tmp *ListNode = nil
    for i := 0; i < N; i++ {
        ll := &ListNode{Val: int(s[i]-'0'), Next: tmp, }
        tmp = ll
    }
    return tmp
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0, Next: nil,}
	t1 := l1
	t2 := l2
	t := result
	carry := 0

	for {
		t.Val += carry
		carry = 0
		if t1 != nil  {
			t.Val += t1.Val
		}
		if t2 != nil {
			t.Val += t2.Val
		}
		if t.Val >= 10 {
			t.Val = t.Val % 10
			carry = 1
		}
		if carry == 1 || (t1 != nil && t1.Next != nil) || (t2 != nil && t2.Next != nil) {
			t.Next = &ListNode{Val: 0, Next: nil,}
			t = t.Next
		} else {
			break
		}
		if t1 != nil {
			t1 = t1.Next
		}
		if t2 != nil {
			t2 = t2.Next
		}
	}

	return result
}

func (ll *ListNode) Print() {
    fmt.Print("LL: [")
    for tail := ll; tail != nil; tail = tail.Next {
        fmt.Print(tail.Val, "->")
    }
    fmt.Println("]")

}

func main() {
	l1 := buildLL(342)
	l1.Print()
	l2 := buildLL(267)
	l2.Print()
	lSum := addTwoNumbers(l1, l2)
	lSum.Print()

        l4 := buildLL(5)
	l4.Print()
        l5 := buildLL(5)
	l5.Print()
	lSum = addTwoNumbers(l4, l5)
	lSum.Print()
}
