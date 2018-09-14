package main

import (
	"fmt"
	"sync"
	"errors"
)

const (
	MAX_STACK = 100
	MAX_QUEUE = 100
)

/**************************************************/
type Stack struct {
	index int
        data  []interface{}
        mux   sync.Mutex
}

func (s *Stack) Push(x interface{}) error {
	s.mux.Lock()
        defer s.mux.Unlock()
	if s.index == MAX_STACK {
		return errors.New("overflow")
	}
	s.data[s.index] = x
	s.index++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	s.mux.Lock()
        defer s.mux.Unlock()
	if s.index == 0 {
		return nil, errors.New("empty")
	}
	s.index--
	x := s.data[s.index]
	return x, nil
}

func testStack(St Stack) {
	A := []string{ "A1", "A2", "B1", "B2", "C1", "C2" }
	for _,v := range A {
		St.Push(v)
	}
	fmt.Print("Stack:")
	for {
            v,err := St.Pop()
            if err != nil {
                break
	    }
	    fmt.Print(v, ",")
	}
	fmt.Println()
}
/**************************************************/

type QNode struct {
	data interface{}
	next *QNode
}

type Queue struct {
	root *QNode
        size  int
	mux  sync.Mutex
}

func New(x interface{}) *QNode {
	qNode := &QNode{ data: x, next: nil, }
	return qNode
}

func (q *Queue) Pushback(x interface{}) error {
	q.mux.Lock()
        defer q.mux.Unlock()
	if q.size == MAX_QUEUE {
		return errors.New("overflow")
	}
	if q.root == nil {
		q.root = New(x)
	        q.size++
		return nil
	}
	var t *QNode
	t = nil
        var lastNode *QNode
	lastNode = nil
	for t = q.root; t != nil; t = t.next {
		lastNode = t
	}
	lastNode.next = New(x)
	q.size++
	return nil
}

func (q *Queue) Popfront() (interface{}, error) {
	q.mux.Lock()
        defer q.mux.Unlock()
	if q.size == 0 {
		return nil, errors.New("empty")
	}
        firstNode := q.root
	x := firstNode.data
	q.root = firstNode.next
	q.size--
	return x, nil
}

func (q Queue) Size() int {
	return q.size
}

func testQueue(Q Queue) {
	A := []string{ "A1", "A2", "B1", "B2", "C1", "C2" }
	for _,v := range A {
		Q.Pushback(v)
	}
	fmt.Print("Queue(size):", Q.Size(), " ")
	for {
            v,err := Q.Popfront()
            if err != nil {
                break
	    }
	    fmt.Print(v, ",")
	}
        fmt.Println()
}

/************
Graph:
       A--------B
        \      /
         \    /
           C
************/

type Node struct {
	data string
	vertices []*Node
}

type Graph struct {
	root *Node
}

func New (nodes []string, adjacencies [][]string) *Graph {
	G := &Graph{root: nil,} 
	for i,v := range nodes {
		node := New(v, adjacencies[i])
		if G.root == nil {
			G.root = node
			continue
		}
	}
	return G
}

func (g Graph) AddNode(to, a *Node) error {
	return nil
}

func (g Graph) Print() {
	t := g.root
	for _,x := range t.vertices {
		fmt.Println(x.data, ": [")
		for _,v := range x.vertices {
		    fmt.Print(v.data, ",")
		}
		fmt.Println("]")
	}
}

func (g Graph) Traverse() {
}

func main() {
	St := Stack{index: 0, data: make([]interface{}, MAX_STACK), mux: sync.Mutex{}, }
	testStack(St)

	Q := Queue{root: nil, size: 0, mux: sync.Mutex{}, }
	testQueue(Q)
        nodes := []string{"A1", "A2", "B1", "B2", "C1", "C2"}
	adjacencies := [][]string{ {"B1", "C1"}, {"B2", "C2"}, {"B2", "C1"}, {"C2", "A1"}, {"A2", "C2"}, {"B1", "A1"} }

        G := New{nodes, adjacencies}
	G.searchDFS("C1")
}
