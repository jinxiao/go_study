package main

import (
	"fmt"
)

//const a int = 65
/*const b = 'A'
const (
	d =1
	e
	f
)*/
const (
	a = 'A'
	b = iota
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
	//fmt.Println(d,"",e,"",f)
	fmt.Println(a &^ c)
}
