package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	defer close(ch)
	go Walk(t1, ch)
	x1 := make([]int, 10)
	for i := 0; i < 10; i++ {
		x1[i] = <-ch
	}
	go Walk(t2, ch)
	x2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		x2[i] = <-ch
	}
	fmt.Println(x1, len(x1))
	fmt.Println(x2, len(x2))

	for i := 0; i < 10; i++ {
		if x1[i] == x2[i] {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {

	t1 := tree.New(1)
	t2 := tree.New(2)

	result := Same(t1, t2)

	fmt.Println(result)
}
