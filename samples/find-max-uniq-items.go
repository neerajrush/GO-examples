package main

import (
	"fmt"
)

type Messenger struct {
	A []int
}

func (m *Messenger) get_data_size() int {
	return len(m.A)
}

func (m *Messenger) get_item(index int) int {
	return m.A[index]
}

type IDMsg struct {
	count int
	ID []bool
	data map[int][]int
}

var IdM IDMsg

func init() {
	IdM = IDMsg{ count: 10, ID: make([]bool, 10), data: make(map[int][]int), }
}

func Solution(m Messenger, whoami int) int {
	if  !IdM.ID[whoami] {
		IdM.ID[whoami] = true
	}
	data_size := m.get_data_size()
	IdM.data[whoami] = make([]int, data_size)
	lx := make([]int, data_size+1)
	for i := 0; i < data_size; i++ {
		data_item := m.get_item(i)
		IdM.data[whoami][i] = data_item
		lx[data_item]++
	}
	result := data_size
	for i := 1; i < data_size+1; i++ {
		if lx[i] == 0 {
			result--
		}
	}
	return result
}

func main() {
	m := Messenger{A: []int{1, 2, 3, 4, 5, 6, 7, 6, 8, 9}, }
	Id := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
	fmt.Println(m.A, Solution(m, Id[1]))
}
