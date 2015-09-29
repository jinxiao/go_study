package main

import (
	"fmt"
)

type person struct {
	Name string
	age  int
}

func main() {
	a := person{
		Name: "joe",
		age:  12,
	}
	fmt.Println(a)
}

func A(per person) {
	per.age = 13
	fmt.Println("A", per)

}
