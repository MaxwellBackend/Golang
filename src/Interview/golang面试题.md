#defer
```
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
```
执行结果：
```
打印后
打印中
打印前
panic: 触发异常

goroutine 1 [running]:
main.defer_call()
	/test/defer.go:17 +0x91
main.main()
	/test/defer.go:8 +0x20
exit status 2
```
defer严格执行后进先出，panic会等待defer执行完成后抛出异常，未声明defer的panic是不可能捕获到异常的
#panic
一旦执行panic，逻辑就会去找defer函数，直到defer之执行完成或者被recover截获
```
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
```
执行结果：
````
进入f()
进入defer函数
recover 截获的panic err: panic里面的内容
完成了recover
````
可以看出在执行main函数之后，先进入f()函数，在函数f()中碰到panic,立马去执行defer的内容，
在defer里面碰到了recover函数，截取panic内容，不在继续执行panic及之后的内容
#foreach的相关问题
````
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
````
执行结果
````
第一个人的名字 => 第三个人的名字
年龄=> 333
第二个人的名字 => 第三个人的名字
年龄=> 333
第三个人的名字 => 第三个人的名字
年龄=> 333
第一个人的名字 => 第一个人的名字
年龄=> 111
第二个人的名字 => 第二个人的名字
年龄=> 222
第三个人的名字 => 第三个人的名字
年龄=> 333
````
可以看到用第一种写法，每一次的&stu都是指向最后一个struct结构体，所以是错误的写法，
当我们需要拿到每一个数据时就必须遍历每一个结构体，需要用第二种方式去执行

#go执行的随机性和闭包
````
import (
	"runtime"
	"sync"
	"fmt"
)
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
````

 runtime.GOMAXPROCS(1) 设置单核跑进程，里面的数字为执行的cpu数,
 sync.WaitGroup{} 用来等待一组gorouinte结束  Add来增加协程的个数 Done来结束协程
 
 
 执行结果：
 ````
 B:  9
 A:  10
 A:  10
 A:  10
 A:  10
 A:  10
 A:  10
 A:  10
 A:  10
 A:  10
 A:  10
 B:  0
 B:  1
 B:  2
 B:  3
 B:  4
 B:  5
 B:  6
 B:  7
 B:  8
 ````
 这里我们可以看到A的内容永远都是10，B的内容时0-9随机输出，这样的原因是由于，第一个
 循环里面的i时函数外面的值，for循环运行完后，i的值为10，所以A为10，第二个循环的i是
 函数参数里的i，与for循环的i完全是不同的，后面的（i）发生值拷贝 go func 指向拷贝地址
 
#go的组合继承
 ````
 type People struct{}
 
 func (p *People) ShowA() {
 	fmt.Println("showA")
 	p.ShowB()
 }
 func (p *People) ShowB() {
 	fmt.Println("showB")
 }
 
 type Teacher struct {
 	People
 }
 
 func (t *Teacher) ShowB() {
 	fmt.Println("teacher showB")
 }
 
 func main() {
 	t := Teacher{}
 	fmt.Println(t.People)
 	t.ShowA()
 }
 ````

 
 输出结果
 ````
 showA
 showB
 ````
  
 go语言的组合，被包含people里的方法上升成了包含者Teacher的方法，但是无法调用到类型方法，
 在people类型方法里无法调用到位置的Teacher类型方法

#select的随机性
````
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "maxwell"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
````

输出结果
````
[root@vagrant-centos65 test]# go run select.go
1
[root@vagrant-centos65 test]# go run select.go
1
[root@vagrant-centos65 test]# go run select.go
1
[root@vagrant-centos65 test]# go run select.go
panic: maxwell

goroutine 1 [running]:
main.main()
	/test/select.go:16 +0x233
exit status 2
[root@vagrant-centos65 test]# go run select.go
1
[root@vagrant-centos65 test]# go run select.go
panic: maxwell

goroutine 1 [running]:
main.main()
	/test/select.go:16 +0x233
exit status 2
[root@vagrant-centos65 test]# go run select.go
panic: maxwell

goroutine 1 [running]:
main.main()
	/test/select.go:16 +0x233
exit status 2
[root@vagrant-centos65 test]# go run select.go
1
````
select随机执行一个case语句，我执行了多次，有的抛出了异常，有的没有，证明select的随机性

#map线程安全
````
type AAA struct {
	mm map[string]int
	sync.Mutex
}

func (a *AAA) Add(str string, i int) {
	a.Lock()
	defer a.Unlock()
	a.mm[str] = i
}

func (a *AAA) Get(str string) int {
	if str, ok := a.mm[str]; ok {
		return str
	}
	return -1
}

func main() {
	a := &AAA{
		mm: make(map[string]int),
	}
	//fmt.Println(a)
	a.Add("aaa", 1111)
	i := a.Get("bbb")
	fmt.Println(i)
}
````
上面的Get（）函数是不安全的，有时候会报错，应该也加上锁，或者直接使用(sync.map)，
由于golang的map是引用类型，当多个goroutine并发调用时会产生竞争，共享资源遭到破坏 !

运行结果：
部分机器不会报错

#闭包引用相同变量
````
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
````
上面的结果：
````
2
3
````
后面的x为已经执行过++的i

#interface内部结构
````
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
````
输出结果为：
````
not nil
nil
````
interface是内建类型，类似于C/C++的void*，第一个为nil是因为interface的类型不是int
#defer和函数返回值
````
func main() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
````
返回结果
````
4
1
3
````
整个defer是在return之前执行的，所以函数1返回4，func()里面类似于1 += 3，这里的t作用域是整个函数
函数2返回1，func()里面类似于1 += 3，但是并不会修改return 的值，这里的t作用域只在func()里
函数3返回3，func（）获取到返回的t=2将其修改为3

#type的使用
````
func getValue() int{
	return 1
}
func main() {
	i := getValue()
	switch i.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
}
````
结果报错，
type只能用于接口，只可以用interface.(type)来确定接口类型，改为：
````
func getValue() interface{} {
	return 1
}
func main() {
	i := getValue()
	switch i.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
}
````
#下面是否会报错
````
package main
 
import (
	"fmt"
)
 
type People interface {
	Speak(string) string
}
 
type Stduent struct{}
 
func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
 
func main() {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
````
很显然这个肯定会报错，struct初始化方式错误，正确的是下面三种
````
  a、var stu Student

  b、var stu *Student=new(Student)

  c、var stu *Student=&Student{}
````
#切片append
````
package main
 
import "fmt"
 
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
````
语法错误，切牌拼接应该加上... append(s1,s2...)

#下面是否会报错
````
package main
const cl  = 100
 
var bl    = 123
 
func main()  {
    println(&bl,bl)
    println(&cl,cl)
}
````
会报错，显然c1是常量，常量会在编译预处理前被展开，当作指令数据，所以不能对c1做&操作、
#goto用法
````
package main
 
func main()  {
 
    for i:=0;i<10 ;i++  {
    loop:
        println(i)
    }
    goto loop
}
````
报出错误 goto loop jumps into block ，goto不能跳进for内部

#附上一份很平常的面试题
https://blog.csdn.net/itcastcpp/article/details/80462619


