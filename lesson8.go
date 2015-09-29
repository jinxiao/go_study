package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "OK"

	a := m[1]
	fmt.Println(a)

}
