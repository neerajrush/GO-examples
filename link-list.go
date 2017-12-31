package main

import (
	"fmt"
)

type Node struct {
	data string
	next* Node
}

func (node Node) get_next() (*Node) {
	return node.next
}

func (node *Node) set_next(newNode* Node) {
	node.next = newNode
}

func (node Node) get_data() (string) {
	return node.data
}

type LinkList struct {
	head* Node
}

func (link_list* LinkList) insert(newNode Node) {
	if link_list.head == nil {
		link_list.head = &newNode
		link_list.head.set_next(nil)
	}
	itor := link_list.head
	prev := itor
	for itor != nil {
	       prev = itor
               itor = itor.get_next()
	}
	prev.set_next(&newNode)
	newNode.set_next(nil)
}

func (link_list LinkList) print_list() {
	itor := link_list.head
        fmt.Println("-----------------------")
	for itor != nil {
		fmt.Println(itor.get_data())
		itor = itor.get_next()
	}
        fmt.Println("-----------------------")
}

func (link_list LinkList) size() int {
	itor := link_list.head
	count := 0
	for itor != nil {
		itor = itor.get_next()
		count += 1
	}
	return count
}

func (link_list* LinkList) reverse() {
	if link_list.head == nil || link_list.head.get_next() == nil {
		return
	}

	x := link_list.head
	y := x.get_next()
	x.set_next(nil)
	p := x
	for y != nil {
		p = x
		x = y
		y = x.get_next()
		x.set_next(p)
	}
	link_list.head = x
	return
}

func (link_list LinkList) search(data string) (bool, int) {
	itor := link_list.head
	count := 0
	for itor != nil {
		if itor.get_data() == data {
			count += 1
		}
		itor = itor.get_next()
	}
	return count > 0, count
}

func (link_list* LinkList) reverseRecursive(x* Node, y* Node, p* Node) {
	if y == nil {
		link_list.head = x
		return
	}

        p = x
	x = y
	y = x.get_next()
	x.set_next(p)

	link_list.reverseRecursive(x, y, p)
}

func (link_list* LinkList) reverseIt() {
	if link_list.head == nil || link_list.head.get_next() == nil {
		return
	}

	x := link_list.head
	y := x.get_next()
	x.set_next(nil)
	p := x

	link_list.reverseRecursive(x, y, p)
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
	ll.print_list()
	fmt.Println(ll.size())
}
