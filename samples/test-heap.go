package main

import (
	"fmt"
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntHeap) Fix() {
	heap.Fix(h, len(*h)-1)
}

func main() {
	h := &IntHeap{3, 1, 5}
        heap.Init(h)
	fmt.Println(h)
	h.Push(2)
	fmt.Println(h)
	h.Fix()
	fmt.Println(h)
	x := h.Pop()
	fmt.Println(x, h)
}
