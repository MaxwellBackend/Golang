## go语言可控制协程

go语言可控制协程主要是解决分布式系统中异步请求Callback模式代码不易阅读,容易出错,此实现利用go语言goroutine+channel

## 使用

1. 创建协程

```golang
co := coroutine.NewCoroutine()
```

2. 协程运行

```golang
err := co.Run(func() error)
```

3. 阻塞协程,恢复主线程

```golang
co.Yield()
```

4. 阻塞主线程(),恢复协程

```golang
co.Resume()
```

5. 恢复主线程,释放协程

```golang
co.Done()
```

## 代码示例

```golang
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
```

## 相关设置
```golang
coroutine.SetResumeWaitSecond(d time.Duration)  设置全局协程Resume后主线程最大等待时间
co.SetResumeWaitSecond(d time.Duration)  设置某个协程Resume后主线程最大等待时间
```
