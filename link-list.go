package main

import (
	"fmt"
)

type Node struct {
	data string
	next* Node
}

func (node Node) getNext() (*Node) {
	return node.next
}

func (node *Node) setNext(newNode* Node) {
	node.next = newNode
}

func (node Node) getData() (string) {
	return node.data
}

type LinkList struct {
	head* Node
}

func (ll* LinkList) insert(newNode Node) {
	if ll.head == nil {
		ll.head = &newNode
		ll.head.setNext(nil)
	}
	itor := ll.head
	prev := itor
	for itor != nil {
	       prev = itor
               itor = itor.getNext()
	}
	prev.setNext(&newNode)
	newNode.setNext(nil)
}

func (ll LinkList) printList() {
	itor := ll.head
        fmt.Println("-----------------------")
	for itor != nil {
		fmt.Println(itor.getData())
		itor = itor.getNext()
	}
        fmt.Println("-----------------------")
}

func (ll LinkList) size() int {
	itor := ll.head
	count := 0
	for itor != nil {
		itor = itor.getNext()
		count += 1
	}
	return count
}

func (ll* LinkList) reverse() {
	if ll.head == nil || ll.head.getNext() == nil {
		return
	}

	x := ll.head
	y := x.getNext()
	x.setNext(nil)
	p := x
	for y != nil {
		p = x
		x = y
		y = x.getNext()
		x.setNext(p)
	}
	ll.head = x
	return
}

func (ll LinkList) search(data string) (bool, int) {
	itor := ll.head
	count := 0
	for itor != nil {
		if itor.getData() == data {
			count += 1
		}
		itor = itor.getNext()
	}
	return count > 0, count
}

func (ll* LinkList) reverseRecursive(x* Node, y* Node, p* Node) {
	if y == nil {
		ll.head = x
		return
	}

        p = x
	x = y
	y = x.getNext()
	x.setNext(p)

	ll.reverseRecursive(x, y, p)
}

func (ll* LinkList) reverseIt() {
	if ll.head == nil || ll.head.getNext() == nil {
		return
	}

	x := ll.head
	y := x.getNext()
	x.setNext(nil)
	p := x

	ll.reverseRecursive(x, y, p)
}

func main() {
	a1 := Node{"A", nil}
	a2 := Node{"B", nil}
	a3 := Node{"C", nil}
	a4 := Node{"D", nil}
	a5 := Node{"E", nil}
	a6 := Node{"A", nil}
	a7 := Node{"F", nil}

	ll := LinkList{nil}

	ll.insert(a1)
	ll.insert(a2)
	ll.insert(a3)
	ll.insert(a4)
	ll.insert(a5)
	ll.insert(a6)
	ll.insert(a7)

	ll.print_list()
	fmt.Println(ll.size())

	ll.reverse()
	ll.print_list()
	fmt.Println(ll.size())

	found, count := ll.search("A")
	fmt.Printf("Found: %t Count:%d\n", found, count)

	found, count = ll.search("P")
	fmt.Printf("Found: %t Count:%d\n", found, count)

	found, count = ll.search("F")
	fmt.Printf("Found: %t Count:%d\n", found, count)

	ll.reverseIt()
	ll.printList()
	fmt.Println(ll.size())
}
