package main

import "fmt"

func main() {
	//	var a [2] int
	//	var b [2] int
	//a := [10]int
	p := [...]int{1, 3, 50, 3, 2}
	l := len(p)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if p[i] > p[j] {
				temp := p[i]
				p[i] = p[j]
				p[j] = temp
			}
		}
	}
	fmt.Println(p)
}
