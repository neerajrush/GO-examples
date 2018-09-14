package main

import (
	"container/list"
	"fmt"
)

func getMappedVal(ip int) string {
	switch ip {
	case 0: return "0"
	case 1: return "1"
	case 2: return "ABC"
	case 3: return "DEF"
	default: return ""
	}
	return ""
}

func generateCombo(resultList *list.List, input, resultStr string, pos int) {
	if pos >=  len(input) {
		resultList.PushBack(resultStr)
		return
	}
	ip := []byte(input)
	str := getMappedVal(int(ip[pos] - '0'))
	for _,s := range str {
		generateCombo(resultList, input, resultStr + string(s), pos+1)
	}
}

func main() {
	resultList := list.New()
	//fmt.Println("Type: %T", resultList)
	input := "23"
	generateCombo(resultList, input, " ", 0)
	for e := resultList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
