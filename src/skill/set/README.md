# Set使用技巧

## 适用场景
我们经常使用map[X]Y的结构来缓存X已经经过处理了，但Y的不同类型可能对性能的影响不一样，所以下面选取了4个类型进行一次性能对比。

## 性能测试结果
```
[root@localhost Golang]# GOPATH=`pwd` && go test -v -test.bench='.*' -test.benchmem skill/set
goos: linux
goarch: amd64
pkg: skill/set
BenchmarkIssetWithInterface-2   	 3000000	       398 ns/op	      71 B/op	       0 allocs/op
BenchmarkIssetWithStruct-2      	 5000000	       320 ns/op	      25 B/op	       0 allocs/op
BenchmarkIssetWithBool-2        	 5000000	       336 ns/op	      29 B/op	       0 allocs/op
BenchmarkIssetWithInt-2         	 5000000	       418 ns/op	      55 B/op	       0 allocs/op
PASS
ok  	skill/set	8.043s
```

## 测试结果
通过基准测试，发现用struct结构定义的set，执行效率最高，内存分配最少。