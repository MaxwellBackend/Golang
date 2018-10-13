package slice

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}

	s1 = append(s1, s2)
	//s1 = append(s1,s2...) // 正确的方式
	fmt.Println(s1)
}
