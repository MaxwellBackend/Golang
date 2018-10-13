#考察interface相关知识
## 1、type的使用（type.go）
结果：
代码报错

结论：
 type只能用于接口，只可以用interface.(type)来确定接口类型
##2、考察interface内部结构（interface.go）
结果：
````
not nil
nil
````
结论：
interface是内建类型，类似于C/C++的void*，第一个为nil是因为interface的类型不是int