package _defer

import "fmt"

func main() {
	// 声明defer，否则无法panic捕获异常
	defer func() {
		fmt.Println("进入defer函数")
		if err := recover(); err != nil {
			fmt.Println("recover 截获的panic err:", err)
		}
		fmt.Println("完成了recover")
	}()
	f()
}
func f() {
	fmt.Println("进入f()")
	panic("panic里面的内容")
	fmt.Println("执行完panic之后的东西")
}
