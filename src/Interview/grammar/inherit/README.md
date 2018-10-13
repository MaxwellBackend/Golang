#考察go的组合继承
##提问输出结果（代码inherit.go）
结果：
````
showA
showB
````
结论：

go语言的组合，被包含people里的方法上升成了包含者Teacher的方法，但是无法调用到类型方法， 在people类型方法里无法调用到位置的Teacher类型方法