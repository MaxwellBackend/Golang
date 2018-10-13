#golang语法相关面试题
## defer
###1、提问输出结果（代码defer.go）
###考察defer执行顺序，defer和panic关系
运行结果： 
````
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
````
结论：
defer严格执行后进先出，panic会等待defer执行完成后抛出异常，未声明defer的panic是不可能捕获到异常的

###2、提问执行结果（代码panic.go）
###考察panic和defer关系
运行结果 ：
````
进入f()
进入defer函数
recover 截获的panic err: panic里面的内容
完成了recover
````
结论： 

可以看出在执行main函数之后，先进入f()函数，在函数f()中碰到panic,立马去执行defer的内容， 在defer里面碰到了recover函数，截取panic内容，不在继续执行panic及之后的内容

###3、提问执行结果（代码defer2.go）
###考察defer和函数返回值
运行结果：
````
4
1
3
````
结论：

整个defer是在return之前执行的，所以函数1返回4，func()里面类似于1 += 3，这里的t作用域是整个函数 函数2返回1，func()里面类似于1 += 3，但是并不会修改return 的值，这里的t作用域只在func()里 函数3返回3，func（）获取到返回的t=2将其修改为3