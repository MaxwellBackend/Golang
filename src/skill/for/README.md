# For循环技巧

## 适用场景
当遍历一个数组时，采用索引遍历的方式比Range遍历要更快更省内存，原因是：Range时会拷贝value的值

## 性能测试结果
```
[root@localhost Golang]# go test -v -test.bench='.*' skill/for
goos: linux
goarch: amd64
pkg: skill/for
BenchmarkForSliceWithIndex-2   	   50000	     35474 ns/op
BenchmarkForSliceWithRange-2   	   20000	     69703 ns/op
```