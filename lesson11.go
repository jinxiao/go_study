package main
import "fmt"
type TZ int

func main() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}

func (tz *TZ) Increase(num int) {
	*tz += TZ(num)

}