package main

import (
	"fmt"
	"math"
	"container/list"
)

type Node struct {
	data int
	left *Node
	right *Node
}

func insertNode(root* Node, data int) *Node {
	if root == nil {
		root = &Node{data: data, left: nil, right: nil,}
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

func printTree(root *Node, order, depth int) {
	if root == nil {
		return
	}
	switch order {
	case 0:    // PreOrder
		for i := 0; i < depth; i++ {
			fmt.Print("-")
		}
		fmt.Println(root.data)
		printTree(root.left, order, depth-2)
		printTree(root.right, order, depth+2)
	case 1:    // InOrder
		printTree(root.left, order, depth-2)
		for i := 0; i < depth; i++ {
			fmt.Print("-")
		}
		fmt.Println(root.data)
		printTree(root.right, order, depth+2)
	case 2:    // PostOrder
		printTree(root.left, order, depth-2)
		printTree(root.right, order, depth+2)
		for i := 0; i < depth; i++ {
			fmt.Print("-")
		}
		fmt.Println(root.data)
	}
}

var prev *Node

func isBSTInOrderTraversal(root *Node) bool {
	if root == nil {
		return true
	}
	if !isBSTInOrderTraversal(root.left) {
		return false
	}
	if prev != nil && prev.data > root.data {
		fmt.Println("prev: ", prev.data, " curr: ", root.data)
		return false
	}
	prev = root
	if isBSTInOrderTraversal(root.right) {
		return true
	}
	return false
}

func isLeftSubtreeLesser(root *Node, data int) bool {
	if root == nil {
		return true
	}
	if root.data < data  &&
	    isLeftSubtreeLesser(root.left, data) &&
             isLeftSubtreeLesser(root.right, data) {
		return true
	}
	return false
}

func isRightSubtreeGreater(root *Node, data int) bool {
	if root == nil {
		return true
	}
	if root.data > data &&
	    isRightSubtreeGreater(root.left, data)  &&
             isRightSubtreeGreater(root.right, data) {
		return true
	}
	return false
}

func isBST(root *Node) bool {
	if root == nil {
		return true
	}
	if  isLeftSubtreeLesser(root.left, root.data) &&
	      isRightSubtreeGreater(root.right, root.data) &&
		isBST(root.left) &&
		 isBST(root.right) {
		return true
	}
	return false
}

func addNode(root* Node, to, what int) {
	if root == nil {
		return
	}
	if root.data == to {
		if root.data > what {
			if root.left == nil {
				root.left = &Node{data: what, left: nil, right: nil,}
				return
			} else {
				return
			}
		}
		if root.data < what {
			if root.right == nil {
				root.right = &Node{data: what, left: nil, right: nil,}
				return
			} else {
				return
			}
		}
	}
	addNode(root.left, to, what)
	addNode(root.right, to, what)
}

func calcMaxHeight(root *Node) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	if root.left == nil {
		return calcMaxHeight(root.right) + 1
	}
	if root.right == nil {
		return calcMaxHeight(root.left) + 1
	}
	l := calcMaxHeight(root.left)
	r := calcMaxHeight(root.right)
	return int(math.Max(float64(l), float64(r))) + 1
}

func calcLevelWidth(root *Node, level int) int {
	if root == nil {
		return 0
	}
	if level == 0 || level == 1 {
		return 1
	}
	return calcLevelWidth(root.left, level-1) + calcLevelWidth(root.right, level-1)
}

func calcWidth(root *Node)int {
	if root == nil {
		return 0
	}
	h := calcMaxHeight(root)
	maxWidth := 0
	for l := 1; l < h+1; l++ {
		wl := calcLevelWidth(root, l)
		//fmt.Println(wl)
		if maxWidth < wl {
			maxWidth = wl
		}
	}
	return maxWidth
}

var qu *list.List

func calcMaxWidth(root *Node) int {
	if root == nil {
		return 0
	}
	qu.PushBack(root)
	maxWidth := 0
	for qu.Len() > 0 {
		if qu.Len() > maxWidth {
			maxWidth = qu.Len()
		}
		//fmt.Println(qu.Len(), maxWidth)
		for e := qu.Front(); e != nil; e = e.Next() {
			if e.Value.(*Node).left != nil {
				qu.PushBack(e.Value.(*Node).left)
			}
			if e.Value.(*Node).right != nil {
				qu.PushBack(e.Value.(*Node).right)
			}
			qu.Remove(e)
		}
	}
	return maxWidth
}

func init() {
	prev = nil
	qu = list.New()
}

func main() {
	A := []int{8, 5, 3, 4, 1, 2, 12, 10, 11, 9}
	depth := 12
	root := buildTree(A)
	printTree(root, 0, depth)
	fmt.Println("++++++++++++++++++")
	printTree(root, 1, depth)
	fmt.Println("++++++++++++++++++")
	printTree(root, 2, depth)
	fmt.Println("++++++++++++++++++")
	result := isBSTInOrderTraversal(root)
	fmt.Println("isBST: ", result)
	fmt.Println("++++++++++++++++++")
	addNode(root, 5, 20) // more than 8 makes it non-BST
	printTree(root, 1, depth)
	prev = nil
	result = isBSTInOrderTraversal(root)
	result1 := isBST(root)
	h := calcMaxHeight(root)
	w := calcMaxWidth(root)
	wl := calcWidth(root)
	fmt.Println("isBST:", result, "isBST:", result1, "Height:", h, "Width:", w, "width;", wl)
	fmt.Println("++++++++++++++++++")
}
