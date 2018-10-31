package bibao

import "fmt"

func bibao(x int) (func(), func()) {

	return func() {
		fmt.Println(x)
		x++
	}, func() {
		fmt.Println(x)
	}
}
func main() {
	a, b := bibao(2)
	a()
	b()
}
