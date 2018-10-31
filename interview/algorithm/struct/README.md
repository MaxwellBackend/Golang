# 算法、数据结构相关面试题
## 请问下面代码结果（代码struct.go）
### 考察struct的初始化
运行结果：出错

正确初始化struct的三种方式：
````
  a、var stu Student

  b、var stu *Student=new(Student)

  c、var stu *Student=&Student{}
````