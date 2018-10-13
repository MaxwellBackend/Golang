package foreach

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "第一个人的名字", Age: 111},
		{Name: "第二个人的名字", Age: 222},
		{Name: "第三个人的名字", Age: 333},
	}
	// 错误写法
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
		fmt.Println("年龄=>", v.Age)
	}

	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
		fmt.Println("年龄=>", v.Age)
	}
}
func main(){
	pase_student()
}
