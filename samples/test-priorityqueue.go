package main

import (
	"fmt"
	"container/heap"
)

type Item struct {
	data string
	index int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int             { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool   { return  pq[i].priority > pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)        { pq[i], pq[j] = pq[j], pq[i]
						pq[i].index = i
						pq[j].index = j
                                              }

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	item.index = -1
	return item
}

// update modify the priority of an item in the queue.
func (pq *PriorityQueue) update(item *Item, data string, priority int) {
	item.data = data
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) print() {
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Printf("%.2d:%s \n", item.priority, item.data)
	}
}

func main() {
	pq := &PriorityQueue{ {data: "Test1", index: 0, priority: 1, }, {data: "Test2", index: 2, priority: 2, }, {data: "Test3", index: 1, priority: 2, } }
	heap.Init(pq)
	pq.print()
}

