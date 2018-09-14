package main

import ( 
	"fmt"
	"errors"
)

type Pizza struct {
	dough string
	souce string
	cheese string
        toppings []string
}

func OrderPizza() (*Pizza, error) {
    c := make(chan int)
    defer close(c)
    P := &Pizza{}
    go P.AddDough(c)
    c <- 0
    for {
	select {
	case n := <-c:
	   fmt.Println("Received:", n)
           if n == 1 {
		go P.AddSource(c)
                c <- 2
	   }
           if n == 3 {
                go P.AddCheese(c)
                c <- 4
	   }
           if n == 5 {
                go P.AddToppings(c)
                c <- 6
	   }
	   if n == 7 {
		return P, nil
	   }
	default:
	}
     }
     return nil, errors.New("Unable to take order.")
}

func Deliver(P *Pizza) {
        fmt.Println("Pizza being delivered:", *P)
}

func (p* Pizza) AddDough(c chan int) {
	select {
	case n := <-c:
	   fmt.Println("AddDough Received:", n)
           if n == 0 {
	       p.dough = "Wheat"
	       c <- 1
	   }
	default:
	}
}

func (p* Pizza) AddSource(c chan int) {
	select {
	case n := <-c:
	   fmt.Println("AddSouce Received:", n)
           if n == 2 {
	       p.souce = "Red"
	       c <- 3
	   }
	default:
	}
}

func (p* Pizza) AddCheese(c chan int) {
	select {
	case n := <-c:
	   fmt.Println("AddCheese Received:", n)
           if n == 4 {
	       p.cheese = "Mozzarella"
	       c <- 5
	   }
	default:
	}
}

func (p* Pizza) AddToppings(c chan int) {
	select {
	case n := <-c:
	   fmt.Println("AddToppings Received:", n)
           if n == 6 {
                p.toppings = make([]string, 3)
                for i := 0; i < 3; i++ {
	            p.toppings[i] = "Topping Type: " +  fmt.Sprintf("%d", i+1)
                }
	         c <- 7
	   }
	default:
	}
}

func main() {
	P, err := OrderPizza()
	if err != nil {
		fmt.Println(err)
		return
	}
	Deliver(P)
}
