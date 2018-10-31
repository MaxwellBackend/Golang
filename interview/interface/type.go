package _interface

import "fmt"

func getValue() int{
	return 1
}
func getValue2() interface{} {
	return 1
}
func main() {
	//i := getValue()
	//switch i.(type) {
	//case int:
	//	fmt.Println("int")
	//default:
	//	fmt.Println("default")
	//}
	j := getValue2()
	switch j.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
}
