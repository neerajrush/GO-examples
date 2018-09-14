package main

import "fmt"

type Shadow interface {}

type Vehicle struct {
	name string
	color string
}

type Stone struct {
	name string
	color string
}

func (v *Vehicle) ChangeColor(c string) bool {
	fmt.Println("ChangeColor(old):", v)
        v.color = c
	fmt.Println("ChangeColor(new):", v)
	return true
}

func (v Vehicle) Drive() bool {
	fmt.Println("Drive:", v)
	return true
}

func Display(s interface{}) {
	fmt.Println("Shadow:", s)
}

func main() {
	V1 := Vehicle{ name: "Tesla", color: "Red", }
	V1.Drive()
	Display(V1)

	S1 := Stone{ name: "Topaz", color: "Gold", }
	//S1.Drive()
	Display(S1)

	V1.ChangeColor("White")
	Display(V1)
}
