package main

import "fmt"

const  (
	A_SIZE = 26
)

type Node struct {
	data string
	nArr []*Node
        isEOW bool
}

func (n *Node) insert(data string) {
	pCrawl := n
	for _,v := range data {
		idx := byte(v) - 'A'
		if pCrawl.nArr[idx] == nil {
			pCrawl.nArr[idx] = &Node{data: string(v), nArr: make([]*Node, A_SIZE), isEOW: false}
		}
		pCrawl = pCrawl.nArr[idx]
	}
	pCrawl.isEOW = true
}

func (n Node) print_trie(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Println(n.data)
	if n.isEOW {
		return
	}
	for i := 0; i < A_SIZE; i++ {
		if n.nArr[i] != nil {
			n.nArr[i].print_trie(depth+2)
		}
	}
}

func (n Node) search(data string) bool {
	if len(data) == 0 {
		return false
	}
	pCrawl := &n
	for _,v := range data {
		idx := byte(v) - 'A'
		if pCrawl.nArr[idx] == nil {
		        fmt.Println(string(v), idx, "nil")
			return false
		}
		pCrawl = pCrawl.nArr[idx]
	}
	if pCrawl != nil && pCrawl.isEOW {
		return true
	}
	return false
}

func remove(A []*Node) []*Node {
	n := len(A)
	if n == 0 {
		return A
	}
	D := A[:n-1]
	S := A[1:n]
	copy(D, S)
	return D
}

func (n Node) traverseBFS() []string {
	pCrawl := &n
	wl := make([]string, 0)
	if pCrawl.isEOW {
		return wl
	}
	qu := make([]*Node, 0)
	qu = append(qu, pCrawl)
	for len(qu) > 0 {
		pCrawl = qu[0]
		qu = remove(qu)
		wl = append(wl, pCrawl.data)
		if pCrawl.isEOW {
			wl = append(wl, "*")
		}
		for i := 0; i < A_SIZE; i++ {
			if pCrawl.nArr[i] != nil {
				qu = append(qu, pCrawl.nArr[i])
			}
		}
	}
	return wl
}

func printPath(node Node, start_word, end_word string) {
        idx := 0
        prefix := ""
        var words_list []string
        top_node := node

        for idx < len(start_word) {
                if start_word[idx] != end_word[idx] {
                        break
		}
                node = *(node.nArr[byte(start_word[idx]) - 'A'])
                prefix += string(start_word[idx])
                idx += 1
	}

        words_list = top_node.traverseBFS()
        fmt.Println(words_list)
}

func main() {
	root := &Node{data: "", nArr: make([]*Node, A_SIZE), isEOW: false,}
	root.insert("LISP")
        root.insert("LIMP")
        root.insert("LIST")
        root.insert("LUMP")
        root.insert("LIKE")
	root.print_trie(0)
	fmt.Println("-----------------")
        fmt.Println(root.search("LISP"))
        fmt.Println(root.search("LIMP"))
        fmt.Println(root.search("LIST"))
        fmt.Println(root.search("LUMP"))
        fmt.Println(root.search("LIKE"))
        fmt.Println(root.search("LIKES"))
        fmt.Println("-----------------")
	printPath(*root, "LISP", "LIKE")
}
