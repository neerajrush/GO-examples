package main

import "fmt"
import "math"

type Node struct {
	data int
	left *Node
	right *Node
}

func (n *Node) print() {
	fmt.Println(n.data)
}

func insertNode(root* Node, data int) *Node{
	if root == nil {
		root = &Node{data: data, left: nil, right: nil}
		return root
	}
	if data < root.data {
		if root.left == nil {
			root.left = &Node{data: data, left: nil, right: nil,}
		} else {
			insertNode(root.left, data)
		}
	} else if data > root.data {
		if root.right == nil {
			root.right = &Node{data: data, left: nil, right: nil,}
		} else {
			insertNode(root.right, data)
		}
	}
	return root
}

func buildTree(A []int) *Node {
	var root *Node = nil
	for _,v := range A {
		root = insertNode(root, v)
	}
	return root
}

func printPreOrderTree(root* Node, depth int) {
	if root == nil {
		return
	}
	//for i := 0; i < depth; i++ {
		//fmt.Print(" ")
	//}
	fmt.Print(root.data, " ")
	printPreOrderTree(root.left, depth-2)
	printPreOrderTree(root.right, depth+2)
}

func printInOrderTree(root* Node, depth int) {
	if root == nil {
		return
	}
	printInOrderTree(root.left, depth-2)
	//for i := 0; i < depth; i++ {
		//fmt.Print(" ")
	//}
	fmt.Print(root.data, " ")
	printInOrderTree(root.right, depth+2)
}

func printPostOrderTree(root* Node, depth int) {
	if root == nil {
		return
	}
	printPostOrderTree(root.left, depth-2)
	printPostOrderTree(root.right, depth+2)
	//for i := 0; i < depth; i++ {
		//fmt.Print(" ")
	//}
	fmt.Print(root.data, " ")
}

func calcHeight(root *Node) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	if root.right == nil {
		return 1 + calcHeight(root.left)
	}
	if root.left == nil {
		return 1 + calcHeight(root.right)
	}
	l := calcHeight(root.left)
	r := calcHeight(root.right)
	return int(math.Max(float64(l), float64(r))) + 1
}

func calcLevelWidth(root *Node, level int) int {
	if root == nil {
		return 0
	}
	if level == 0 || level == 1 {
		return 1
	}
	lw := calcLevelWidth(root.left, level-1)
	rw := calcLevelWidth(root.right, level-1)
	return lw + rw
}

func calcMaxWidth(root *Node) int {
	if root == nil {
		return 0
	}
	h := calcHeight(root)
	maxWidth := 0
	for l := 1; l < h+1; l++ {
		x := calcLevelWidth(root, l)
		if maxWidth < x {
			maxWidth = x
		}
	}
	return maxWidth
}

func main() {
	A := []int{5, 4, 7, 8, 2, 3, 9, 1, 6}
	root := buildTree(A)
	insertNode(root, 10)
	fmt.Println("----- PreOrder ------")
	printPreOrderTree(root, 10)
	fmt.Println("")
	fmt.Println("----- ------- ------")
	fmt.Println("----- InOrder ------")
	printInOrderTree(root, 10)
	fmt.Println("")
	fmt.Println("----- ------- ------")
	fmt.Println("----- PostOrder ------")
	printPostOrderTree(root, 10)
	fmt.Println("")
	fmt.Println("----- ------- ------")
	fmt.Println("Height: ", calcHeight(root))
	fmt.Println("Max Widht: ", calcMaxWidth(root))
}
