package _interface

import "fmt"

func Func(x interface{}) {
	if x == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
func main() {
	var i *int = nil
	Func(i)
	var j interface{} = nil
	Func(j)
}
