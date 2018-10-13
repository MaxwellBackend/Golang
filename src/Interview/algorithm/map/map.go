package _map

import (
	"sync"
	"fmt"
)

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
