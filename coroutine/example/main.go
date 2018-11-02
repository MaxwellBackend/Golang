package main

import (
	"coroutine"
	"fmt"
)

func main() {
	for {
		co1 := coroutine.NewCoroutine()
		co1.Run(func() error {
			fmt.Println("handle client req 1")
			fmt.Println("send rpc 1")
			co1.Yield() // 协程等待rpc 1返回
			fmt.Println("handle rpc 1 result")
			fmt.Println("handle client req 1 complete")
			co1.Done()
			return nil
		})

		co2 := coroutine.NewCoroutine()
		co2.Run(func() error {
			fmt.Println("handle client req 2")
			return nil
		})

		co3 := coroutine.NewCoroutine()
		co3.Run(func() error {
			fmt.Println("recived rpc 1 result")
			co1.Resume()
			fmt.Println("handle rpc 1 result complete")
			return nil
		})
		fmt.Println("收到STOP")
		break
	}
}
