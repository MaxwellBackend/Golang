#考察select的随机性
##提问输出结果（代码select.go）
结果：
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
结论：

select随机执行一个case语句，我执行了多次，有的抛出了异常，有的没有，证明select的随机性